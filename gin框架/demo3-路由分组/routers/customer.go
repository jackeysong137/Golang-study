package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
外部使用必须要为公有 就是首字母大写
*/
func CustomerRouteInit(router *gin.Engine) {
	customerRoute := router.Group("/customer") //创建一个路由分组 和java的controller很像
	/*
		{} 标识是一个代码块  目的是为了更好区分，这个也是go的编码风格
		 表示这个是一组
	*/
	{
		//这里的属于一个分组 路由都是/customer开头
		customerRoute.POST("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "添加客户")
		})
		customerRoute.GET("customerId", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "查询客户")
		})

	}
}
