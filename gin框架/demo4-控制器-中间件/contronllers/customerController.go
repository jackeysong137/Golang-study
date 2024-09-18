package contronllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
}

func (con CustomerController) Add(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加客户接口")
}

func (con CustomerController) Edit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "修改客户接口")
}
