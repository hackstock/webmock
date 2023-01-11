package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackstock/webmock/pkg/parsing"
	"go.uber.org/zap"
)

type Server struct {
	router    *gin.Engine
	endpoints map[string]parsing.Endpoint
	logger    *zap.Logger
}

func NewServer(r *gin.Engine, e map[string]parsing.Endpoint, l *zap.Logger) *Server {
	return &Server{
		router:    r,
		endpoints: e,
		logger:    l,
	}
}

func (s *Server) Run(port int) error {
	s.router.Any("/:path", func(ctx *gin.Context) {
		requestPath := ctx.Param("path")

		if endpoint, found := s.endpoints[fmt.Sprintf("/%s", requestPath)]; found {
			if ctx.Request.Method == endpoint.HTTPMethod {
				ctx.JSON(endpoint.StatusCode, endpoint.Response)
			} else {
				ctx.Status(http.StatusNotFound)
			}
		}

	})

	return s.router.Run(fmt.Sprintf(":%d", port))
}
