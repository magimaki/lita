package transport

import (
	"lita/endpoint"

	"github.com/gin-gonic/gin"
)

func NewBaseHTTPHandler(ep endpoint.BaseEndpoint, group *gin.RouterGroup) {
	group.GET("/healthCheck", ep.HealthCheckEndpoint)
}
