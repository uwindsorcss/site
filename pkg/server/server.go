package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/uwindsorcss/site/pkg/config"
)

// server.Server: server struct, it holds the majority of the Server
// `Log`s buffer should be flushed by user
type Server struct {
	Log    *zap.SugaredLogger
	Config *config.Config
	Router *gin.Engine
}

func ProvideServer(l *zap.SugaredLogger, c *config.Config) Server {
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	s := Server{
		Log:    l,
		Config: c,
		Router: r,
	}

	s.LoadMiddleware()
	s.LoadRoutes()

	return s
}
