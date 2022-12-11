package server

import (
	"context"
	"net/http"
	"time"

	"github.com/Pranc1ngPegasus/trial-field/domain/configuration"
	"github.com/Pranc1ngPegasus/trial-field/domain/logger"
)

func NewServer(
	ctx context.Context,
	logger logger.Logger,
	config configuration.Configuration,
	rootHandler http.Handler,
) *http.Server {
	cfg := config.Config()

	logger.Info(ctx, "listen on", logger.Field("port", cfg.ServerPort))

	return &http.Server{
		Addr:              ":" + cfg.ServerPort,
		Handler:           rootHandler,
		ReadHeaderTimeout: 10 * time.Second,
	}
}
