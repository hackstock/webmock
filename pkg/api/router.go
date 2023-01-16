package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackstock/webmock/pkg/parsing"
	"go.uber.org/zap"
)

func InitRoutes(endpoints map[string]parsing.Endpoint, logger *zap.Logger) *gin.Engine {
	router := gin.Default()

	router.Any("/:path", func(ctx *gin.Context) {
		requestPath := ctx.Param("path")

		if endpoint, found := endpoints[fmt.Sprintf("/%s", requestPath)]; found {
			if ctx.Request.Method == endpoint.HTTPMethod {
				ctx.JSON(endpoint.StatusCode, endpoint.Response)
			} else {
				// TODO : Add handling for other HTTP methods
			}
		} else {
			ctx.Status(http.StatusNotFound)
		}

	})

	return router
}
