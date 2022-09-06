package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prairir/gin-no-floc/noFloc"
	"go.uber.org/zap"
)

func (s *Server) LoadMiddleware() {
	s.Router.Use(noFloc.Middleware())
	s.Router.Use(StructuredLogger(s.Log))
	s.Router.Use(gin.Recovery())
}

// server.StructuredLogger: is a logging middleware using a zap sugared logger instance
// prints in the format `time, response status, latency, address, request type, route`
func StructuredLogger(l *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		path := c.Request.URL.Path
		if c.Request.URL.RawQuery != "" {
			path = path + "?" + c.Request.URL.RawQuery
		}

		if c.Writer.Status() >= 500 {
			l.Errorw("request error",
				zap.String("client_addr", c.ClientIP()),
				zap.Duration("latency", time.Now().Sub(start)),
				zap.Int("status", c.Writer.Status()),
				zap.String("mathod", c.Request.Method),
				zap.String("path", path),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			)
			return
		}

		l.Debugw("request",
			zap.String("client_addr", c.ClientIP()),
			zap.Duration("latency", time.Now().Sub(start)),
			zap.Int("status", c.Writer.Status()),
			zap.String("mathod", c.Request.Method),
			zap.String("path", path),
		)
	}
}
