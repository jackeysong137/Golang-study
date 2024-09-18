package main

import (
	"demo4/middlewares"
	"demo4/routers"

	"github.com/gin-gonic/gin"
)

/*
	  控制器就是controller
	  中间件(middleware)类似java netty的处理器
		router.GET("path","中间件1","中间件2"...控制器)
		在绑定路由的时候 后面可以传入一系列的方法，按顺序执行，最后一个叫控制器，其他的都是中间件
		gin 有全局中间件 分组中间件 和路由中间件
		全局中间件 所有的请求都会先执行的方法 使用router.use(middleware1,middleware2...)
		分组中间件 路由分组内都会先执行
		路由中间件 当前路由会执行
		gin.Default() 就执行设置了全局中间件 ，Logger 和Recovery 记录日志和 处理panic
*/
func main() {

	router := gin.Default()
	//设置全局中间件
	router.Use(middlewares.TimeMiddlerware_1)

	//注册路由
	routers.CustomerRouterInit(router)
	routers.FirmRouterInit(router)

	router.Run(":8080")
}
