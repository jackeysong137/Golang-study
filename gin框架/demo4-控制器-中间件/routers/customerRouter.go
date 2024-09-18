package routers

import (
	"demo4/contronllers"
	"demo4/middlewares"

	"github.com/gin-gonic/gin"
)

func CustomerRouterInit(router *gin.Engine) {
	//获取路由分组 并设置分组中间件
	customerRouter := router.Group("/customer", middlewares.CustomerMiddileware_1, middlewares.CustomerMiddileware_2)

	{
		//控制器分组 其实就是抽取controller
		customerController := contronllers.CustomerController{}
		customerRouter.POST("/add", customerController.Add)
		customerRouter.POST("/edit", customerController.Edit)
	}

}
