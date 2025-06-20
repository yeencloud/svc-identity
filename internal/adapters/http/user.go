package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type notimplementedError struct{}

func (e notimplementedError) Error() string {
	return "not implemented"
}

func (e notimplementedError) RestCode() int {
	return http.StatusNotImplemented
}

func (s *HTTPServer) createUserHandler(ctx *gin.Context) (any, error) {
	return nil, notimplementedError{}
}
