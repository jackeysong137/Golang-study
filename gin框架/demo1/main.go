package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//获取默认的路由引擎
	engine := gin.Default()
	//设置html的模版路径（这个是为了配合模版使用）
	engine.LoadHTMLGlob("./templates/**")

	//测试get请求
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello gin ")
	})
	engine.GET("/info", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello gin ")
	})
	//返回json数据
	engine.GET("/userInfo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"username": "jack",
			"age":      18,
		})
	})
	//返回结构体
	engine.POST("/artcle", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, Aritcle{
			Title:  "red and black",
			Author: "xx",
		})
	})

	/*
		jsonp 请求http://localhost:8080/testjosnp?callback=xx  在请求的后么添加callback函数
		ctx.JSONP（）返回的数据作为参数执行callback的方法
	*/
	engine.GET("testjosnp", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, Aritcle{
			Title:  "red and black",
			Author: "xx",
		})
	})

	//xml
	engine.GET("xml", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, gin.H{
			"name": "carol",
		})
	})
	//模版渲染  在html中 使用{{.key}} 实现占位符的替换
	engine.GET("/html1", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"name": "karel",
		})
	})

	//启动 默认8080端口 需要修改->engine.Run(":8081")
	engine.Run()
}

type Aritcle struct {
	Title  string
	Author string
}
