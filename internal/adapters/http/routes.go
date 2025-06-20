package http

import (
	"github.com/gin-gonic/gin"
)

func (s *HTTPServer) registerRoutes(engine *gin.Engine) {
	r := engine.Group("/identity", s.server.RequireCorrelationID, s.server.RequireRequestID)

	r.POST("/user", s.wrapped(s.createUserHandler))
}
