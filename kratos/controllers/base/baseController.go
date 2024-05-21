package base

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (con BaseController) success(ctx *gin.Context) {
	ctx.String(200, "success")
}
