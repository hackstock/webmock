package api

import (
	"fmt"
	"net/http"

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
	for i := 1; i <= 5; i++ {
		s.router.GET(fmt.Sprintf("/%d", i), func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"key": fmt.Sprintf("%d", i),
			})
		})
	}
	return s.router.Run(fmt.Sprintf(":%d", port))
}
