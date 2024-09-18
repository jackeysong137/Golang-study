package routers

import (
	"demo4/contronllers"
	"demo4/middlewares"

	"github.com/gin-gonic/gin"
)

func FirmRouterInit(router *gin.Engine) {
	//获取路由分组 并设置路由中间件
	firmRouter := router.Group("/firm", middlewares.FirmMiddleware_1, middlewares.FirmMiddleware_2)

	{
		//controller
		firmController := contronllers.FirmController{}
		firmRouter.POST("/add", firmController.Add)
		firmRouter.POST("/edit", firmController.Edit)
	}

}
