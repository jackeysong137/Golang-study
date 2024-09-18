package main

import (
	"demo3/routers"

	"github.com/gin-gonic/gin"
)

/*
所有的路由写在一个文件很不好维护 我们把路由进行分组
*/
func main() {
	//获取默认的路由引擎
	router := gin.Default()

	//路由组册
	//customer 相关
	routers.CustomerRouteInit(router)

	// firm相关
	routers.FirmRouterInit(router)
	router.Run(":8080")
}
