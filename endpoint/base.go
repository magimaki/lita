package endpoint

import (
	"lita/service"

	"github.com/gin-gonic/gin"
)

type BaseEndpoint struct {
	HealthCheckEndpoint gin.HandlerFunc
}

func NewBaseEndpoint(svc service.BaseService) BaseEndpoint {
	return BaseEndpoint{
		HealthCheckEndpoint: makeHealthCheck(svc),
	}
}

type HealthCheckRequest struct{}

type HealthCheckResponse struct {
	Status bool `json:"status"`
}

func makeHealthCheck(svc service.BaseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := svc.HealthCheck()
		c.JSON(200, HealthCheckResponse{
			Status: status,
		})
	}
}
