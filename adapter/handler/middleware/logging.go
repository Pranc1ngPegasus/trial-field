package middleware

import (
	"net/http"
	"time"

	"github.com/Pranc1ngPegasus/trial-field/domain/logger"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{
		w,
		http.StatusOK,
		0,
	}
}

func (m *loggingResponseWriter) Status() int {
	return m.status
}

func (m *loggingResponseWriter) BytesWritten() int {
	return m.size
}

func (m *loggingResponseWriter) WriteHeader(status int) {
	m.status = status
	m.ResponseWriter.WriteHeader(status)
}

func (m *loggingResponseWriter) Write(b []byte) (int, error) {
	m.size = len(b)

	return m.ResponseWriter.Write(b)
}

func Logging(logger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := NewLoggingResponseWriter(w)
			t1 := time.Now()

			defer func() {
				logger.Info("Served",
					logger.Field("proto", r.Proto),
					logger.Field("status", r.Method),
					logger.Field("path", r.URL.Path),
					logger.Field("duration", time.Since(t1).String()),
					logger.Field("status", ww.Status()),
					logger.Field("size", ww.BytesWritten()),
				)
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
