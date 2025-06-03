//go:build wireinject
// +build wireinject

package main

import (
	"boilerplate/config"
	"boilerplate/config/infra"
	"boilerplate/config/logger"
	"boilerplate/config/router"
	"boilerplate/external"
	"boilerplate/internal/handler"
	"boilerplate/internal/repository/example"
	"boilerplate/internal/repository/tx"
	"boilerplate/internal/usecase"
	"boilerplate/transport"
	"github.com/google/wire"
)

var Configs = wire.NewSet(
	config.ProvideConfig,
)

var LoggerSet = wire.NewSet(
	logger.ProvideLogger,
)

var Redis = wire.NewSet(
	infra.RedisNewClient,
)

var InfraSet = wire.NewSet(
	infra.ProvideInfra,
)

var logDb = wire.NewSet(
	logger.NewLoggerDb,
)

var External = wire.NewSet(
	external.ProvideExternalSvc,
)

var RepoSet = wire.NewSet(
	example.ProvideBankRepo,
	tx.ProvideTxManager,
)

var InternalDomain = wire.NewSet(
	RepoSet,
	usecase.ProvideUsc,
)

var Handler = wire.NewSet(
	handler.ProvideHandler,
)

var Server = wire.NewSet(
	Handler,
	router.ProvideRoute,
	transport.ProvideHttp,
)

func ServerApp() *transport.HTTP {
	wire.Build(
		Configs,
		LoggerSet,
		Redis,
		InfraSet,
		External,
		logDb,
		InternalDomain,
		Server)

	return &transport.HTTP{}
}
