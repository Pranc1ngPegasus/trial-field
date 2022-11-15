package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Pranc1ngPegasus/middlechain"
	"github.com/google/wire"
)

var _ http.Handler = (*Handler)(nil)

var NewHandlerSet = wire.NewSet(
	wire.Bind(new(http.Handler), new(*Handler)),
	NewHandler,
)

type Handler struct {
	schema graphql.ExecutableSchema

	router http.Handler
}

func NewHandler(
	schema graphql.ExecutableSchema,
) *Handler {
	mux := http.NewServeMux()

	h := &Handler{
		schema: schema,
		router: mux,
	}

	mux.HandleFunc("/ping", h.ping)
	mux.HandleFunc("/graphql", h.graphql)
	mux.HandleFunc("/play", h.playground)

	return h
}

func (h *Handler) ping(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("pong")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) graphql(w http.ResponseWriter, r *http.Request) {
	srv := handler.NewDefaultServer(h.schema)
	handler := middlechain.Chain(srv)
	handler.ServeHTTP(w, r)
}

func (h *Handler) playground(w http.ResponseWriter, r *http.Request) {
	play := playground.Handler("Trial Field", "/graphql")
	play.ServeHTTP(w, r)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
