package server

import (
	"github.com/prairir/gin-no-floc/noFloc"
)

func (s *Server) LoadMiddleware() {
	s.Router.Use(noFloc.Middleware())
}
