package service

import (
	"class/passenger/api-gateway/global"
	"class/passenger/api-gateway/internal/dto"
	carrpc "class/passenger/api-gateway/proto/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Demo 模板Demo
func Demo(c *gin.Context) {
	var form dto.Demo
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	demo, err := global.CarRpcClient.Demo(c, &carrpc.DemoReq{
		Name: form.Name,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Demo 错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Demo 正常", "data": gin.H{
		"demo": demo.Demo,
	}})
	return
}
