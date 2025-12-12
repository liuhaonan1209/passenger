package router

import (
	"class/passenger/api-gateway/internal/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/demo", service.Demo)
	return r
}
