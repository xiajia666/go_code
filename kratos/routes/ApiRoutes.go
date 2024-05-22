package routes

import (
	"github.com/gin-gonic/gin"
	"model/kratos/controllers/admin"
)

func ApiRoutes(context *gin.Engine) {
	context.GET("/apiget", admin.Upload)
}
