package logger

import (
	"context"
	"fmt"

	domain "github.com/Pranc1ngPegasus/trial-field/domain/logger"
	"github.com/google/wire"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var _ domain.Logger = (*Logger)(nil)

type Logger struct {
	logger *zap.Logger
}

var NewLoggerSet = wire.NewSet(
	wire.Bind(new(domain.Logger), new(*Logger)),
	NewLogger,
)

func NewLogger() (*Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	return &Logger{
		logger: log,
	}, nil
}

func (l *Logger) Field(key string, message interface{}) domain.Field {
	return domain.Field{
		Key:       key,
		Interface: message,
	}
}

func (l *Logger) field(field domain.Field) zap.Field {
	switch i := field.Interface.(type) {
	case error:
		return zap.Error(i)
	case string:
		return zap.String(field.Key, i)
	case int:
		return zap.Int(field.Key, i)
	case bool:
		return zap.Bool(field.Key, i)
	default:
		return zap.Any(field.Key, i)
	}
}

func (l *Logger) traceFor(ctx context.Context) []domain.Field {
	span := trace.SpanContextFromContext(ctx)

	return []domain.Field{
		l.Field("logging.googleapis.com/trace", span.TraceID().String()),
		l.Field("logging.googleapis.com/spanId", span.SpanID().String()),
		l.Field("logging.googleapis.com/trace_sampled", span.IsSampled()),
	}
}

func (l *Logger) Info(ctx context.Context, message string, fields ...domain.Field) {
	fields = append(fields, l.traceFor(ctx)...)

	zapfields := lo.Map(fields, func(field domain.Field, _ int) zap.Field {
		return l.field(field)
	})

	l.logger.With()

	l.logger.Info(message, zapfields...)
}

func (l *Logger) Error(ctx context.Context, message string, fields ...domain.Field) {
	fields = append(fields, l.traceFor(ctx)...)

	zapfields := lo.Map(fields, func(field domain.Field, _ int) zap.Field {
		return l.field(field)
	})

	l.logger.Error(message, zapfields...)
}
