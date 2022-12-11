//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"net/http"

	"github.com/Pranc1ngPegasus/trial-field/adapter/handler"
	"github.com/Pranc1ngPegasus/trial-field/adapter/resolver"
	"github.com/Pranc1ngPegasus/trial-field/adapter/server"
	domainlogger "github.com/Pranc1ngPegasus/trial-field/domain/logger"
	domaintracer "github.com/Pranc1ngPegasus/trial-field/domain/tracer"
	"github.com/Pranc1ngPegasus/trial-field/infra/configuration"
	"github.com/Pranc1ngPegasus/trial-field/infra/logger"
	"github.com/Pranc1ngPegasus/trial-field/infra/tracer"
	"github.com/google/wire"
)

type app struct {
	logger domainlogger.Logger
	tracer domaintracer.Tracer
	server *http.Server
}

func initialize() (*app, error) {
	wire.Build(
		context.Background,

		logger.NewLoggerSet,

		configuration.NewConfigurationSet,

		tracer.NewTracerSet,

		resolver.NewSchema,

		handler.NewHandlerSet,

		server.NewServer,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
