package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Pranc1ngPegasus/trial-field/domain/configuration"
	"github.com/Pranc1ngPegasus/trial-field/domain/logger"
)

func NewServer(
	logger logger.Logger,
	config configuration.Configuration,
	rootHandler http.Handler,
) *http.Server {
	cfg := config.Config()

	logger.Info(fmt.Sprintf("listen on %s", cfg.ServerPort))

	return &http.Server{
		Addr:              ":" + cfg.ServerPort,
		Handler:           rootHandler,
		ReadHeaderTimeout: 10 * time.Second,
	}
}
