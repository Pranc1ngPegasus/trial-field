//go:build wireinject
// +build wireinject

package main

import (
	domainconfiguration "github.com/Pranc1ngPegasus/trial-field/domain/configuration"
	domainlogger "github.com/Pranc1ngPegasus/trial-field/domain/logger"
	"github.com/Pranc1ngPegasus/trial-field/infra/configuration"
	"github.com/Pranc1ngPegasus/trial-field/infra/logger"
	"github.com/google/wire"
)

type app struct {
	logger domainlogger.Logger
	config domainconfiguration.Configuration
}

func initialize() (*app, error) {
	wire.Build(
		logger.NewLoggerSet,

		configuration.NewConfigurationSet,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
