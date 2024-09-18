package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//获取默认的路由引擎
	router := gin.Default()
	//get 获取请求参数 url后面的参数
	router.GET("/userInfo", func(ctx *gin.Context) {
		id := ctx.Query("id")
		username := ctx.DefaultQuery("username", "zhangsan") //默认值：zhangsan
		ctx.JSON(http.StatusOK, gin.H{
			"id":       id,
			"username": username,
		})
	})
	//post 获取表单的数据
	router.POST("/from", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.DefaultPostForm("password", "111")
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})

	})
	//使用结构体来接收请求参数
	router.POST("/addCutomer", func(ctx *gin.Context) {
		customer := Customer{}
		/*
			这种方式可以介绍json数据 form数据和xml数据 只需在结构体中有对应的标签就行
			json
			form
			xml
		*/
		err := ctx.ShouldBind(&customer) //这里需要传入指针
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(customer)
		ctx.JSON(http.StatusOK, customer)

	})
	//path 参数 路径传参 http://localhost:8080/getInfo/111/zhansan
	router.GET("/getInfo/:uid/:uname", func(ctx *gin.Context) {
		uid := ctx.Param("uid")
		uname := ctx.Param(("uname"))
		ctx.JSON(http.StatusOK, gin.H{
			"uid":   uid,
			"uname": uname,
		})
	})

	//启动 绑定8080 端口
	router.Run(":8080")
}

type Customer struct {
	/*
		form 是固定的key 来指定接收form的数据
		 json是来接收body的数据
		 xml是请求数据是xml的时候也可以 支付的时候可能会用到
	*/
	CustomerId   string `json:"customerId" form:"customerId" xml:"customerId"`       // 这里指定form是用来接收form的数据
	CustomerName string `json:"customerName" form:"customerName" xml:"customerName"` // 这里指定form是用来接收form的数据
}
