package contronllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FirmController struct {
}

func (con FirmController) Add(ctx *gin.Context) {
	ctx.String(http.StatusOK, "添加企业接口")
}

func (con FirmController) Edit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "修改企业接口")
}
