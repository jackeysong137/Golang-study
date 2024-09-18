package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FirmRouterInit(router *gin.Engine) {
	//创建一个路由分组
	firmRouter := router.Group("/firm")
	{
		firmRouter.POST("/addFrim", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "添加企业")
		})
		firmRouter.GET("/getFirm", func(ctx *gin.Context) {

			ctx.String(http.StatusOK, "查询企业")
		})
	}
}
