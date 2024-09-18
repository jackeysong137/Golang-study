package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

/*

 记录执行时间
*/

func TimeMiddlerware_1(ctx *gin.Context) {
	fmt.Println("全局中间件：记录执行时间")
	start := time.Now().Unix()
	//执行下面的中间件
	ctx.Next()
	//等上面 一步后面的所有中间件执行完成后执行这里
	end := time.Now().Unix()
	fmt.Println("记录执行时间", end-start)

}
