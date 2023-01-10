package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router *gin.Engine
	logger *zap.Logger
}

func NewServer(r *gin.Engine, l *zap.Logger) *Server {
	return &Server{
		router: r,
		logger: l,
	}
}

func (s *Server) Run(port int) error {
	return s.router.Run(fmt.Sprintf(":%d", port))
}
