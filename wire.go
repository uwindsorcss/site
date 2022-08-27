package main

import (
	"github.com/google/wire"

	"github.com/uwindsorcss/site/pkg/config"
	"github.com/uwindsorcss/site/pkg/log"
	"github.com/uwindsorcss/site/pkg/server"
)

func InjectDependencies() (server.Server, error) {
	wire.Build(config.ProvideConfig, log.ProvideLog, server.ProvideServer)
	return server.Server{}, nil
}
