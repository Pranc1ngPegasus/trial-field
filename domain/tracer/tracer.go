package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type Tracer interface {
	Tracer() trace.Tracer
	Stop(context.Context) error
}
