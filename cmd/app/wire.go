//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/Pranc1ngPegasus/trial-field/adapter/handler"
	"github.com/Pranc1ngPegasus/trial-field/adapter/resolver"
	"github.com/Pranc1ngPegasus/trial-field/adapter/server"
	domainlogger "github.com/Pranc1ngPegasus/trial-field/domain/logger"
	"github.com/Pranc1ngPegasus/trial-field/infra/configuration"
	"github.com/Pranc1ngPegasus/trial-field/infra/logger"
	"github.com/google/wire"
)

type app struct {
	logger domainlogger.Logger
	server *http.Server
}

func initialize() (*app, error) {
	wire.Build(
		logger.NewLoggerSet,

		configuration.NewConfigurationSet,

		resolver.NewSchema,

		handler.NewHandlerSet,

		server.NewServer,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
