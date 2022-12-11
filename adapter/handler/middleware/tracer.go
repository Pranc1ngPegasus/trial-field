package middleware

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Tracer(h http.Handler) http.Handler {
	return otelhttp.NewHandler(
		h,
		"trial-field",
		otelhttp.WithMessageEvents(
			otelhttp.ReadEvents,
			otelhttp.WriteEvents,
		),
	)
}
