package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CustomerMiddileware_1(ctx *gin.Context) {

	fmt.Println("CustomerMiddileware_1 start ..")
	ctx.Next()
	fmt.Println("CustomerMiddileware_1 end ..")
}

func CustomerMiddileware_2(ctx *gin.Context) {

	fmt.Println("CustomerMiddileware_2 start ..")
	ctx.Next()
	fmt.Println("CustomerMiddileware_2 end ..")
}
