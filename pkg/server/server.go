package server

import (
	"go.uber.org/zap"

	"github.com/uwindsorcss/site/pkg/config"
)

// server.Server: server struct, it holds the majority of the Server
// `Log`s buffer should be flushed by user
type Server struct {
	Log    *zap.SugaredLogger
	Config *config.Config
}

func ProvideServer(l *zap.SugaredLogger, c *config.Config) Server {
	return Server{
		Log:    l,
		Config: c,
	}
}
