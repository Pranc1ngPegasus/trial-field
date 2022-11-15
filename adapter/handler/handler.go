package handler

import (
	"net/http"

	"github.com/google/wire"
)

var _ http.Handler = (*Handler)(nil)

var NewHandlerSet = wire.NewSet(
	wire.Bind(new(http.Handler), new(*Handler)),
	NewHandler,
)

type Handler struct {
	router http.Handler
}

func NewHandler() *Handler {
	mux := http.NewServeMux()

	h := &Handler{
		router: mux,
	}

	mux.HandleFunc("/ping", h.ping)

	return h
}

func (h *Handler) ping(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
