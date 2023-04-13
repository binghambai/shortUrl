package routers

import (
	service "binghambai.com/shortUrl/app/service"
	"github.com/gin-gonic/gin"
)

func InitAllRouters(router *gin.Engine) {
	v1 := router.Group("/api/json")
	{
		v1.POST("/login", service.Test)
	}
}
