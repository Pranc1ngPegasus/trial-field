package tracer

import (
	"context"
	"fmt"

	exporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	gcppropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"github.com/Pranc1ngPegasus/trial-field/domain/configuration"
	"github.com/Pranc1ngPegasus/trial-field/domain/logger"
	domain "github.com/Pranc1ngPegasus/trial-field/domain/tracer"
	"github.com/google/wire"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.9.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"
)

var _ domain.Tracer = (*Tracer)(nil)

var NewTracerSet = wire.NewSet(
	wire.Bind(new(domain.Tracer), new(*Tracer)),
	NewTracer,
)

type Tracer struct {
	exporter *exporter.Exporter
	provider *sdktrace.TracerProvider
}

func NewTracer(
	logger logger.Logger,
	config configuration.Configuration,
) (*Tracer, error) {
	exporter, err := exporter.New(exporter.WithProjectID(config.Config().GCPProjectID))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize exporter: %w", err)
	}

	res, err := resource.New(context.Background(),
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("trial-field"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize resource: %w", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			gcppropagator.CloudTraceOneWayPropagator{},
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return &Tracer{
		exporter: exporter,
		provider: tp,
	}, nil
}

func (t *Tracer) Tracer() trace.Tracer {
	return t.provider.Tracer("")
}

func (t *Tracer) Stop(ctx context.Context) error {
	var merr error

	multierr.AppendInto(&merr, t.exporter.Shutdown(ctx))
	multierr.AppendInto(&merr, t.provider.Shutdown(ctx))

	return nil
}
