package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// routes.Load: load routes into engine
func (s *Server) LoadRoutes() {
	// example routes
	s.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	s.Router.GET("/fuck", func(c *gin.Context) {
		c.Status(500)
	})
}
