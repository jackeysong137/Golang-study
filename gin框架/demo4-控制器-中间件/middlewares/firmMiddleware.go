package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func FirmMiddleware_1(ctx *gin.Context) {
	fmt.Println("FirmMiddleware_1.....start")
	ctx.Next()
	fmt.Println("FirmMiddleware_1.....end")
}

func FirmMiddleware_2(ctx *gin.Context) {
	fmt.Println("FirmMiddleware_2.....start")
	ctx.Next()
	fmt.Println("FirmMiddleware_2.....end")
}
