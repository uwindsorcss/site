package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// routes.Load: load routes into engine
func (s *Server) LoadRoutes() {
	s.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
