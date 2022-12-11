package middleware

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Tracer(h http.Handler) http.Handler {
	return otelhttp.NewHandler(
		h,
		"server",
		otelhttp.WithMessageEvents(
			otelhttp.ReadEvents,
			otelhttp.WriteEvents,
		),
	)
}
