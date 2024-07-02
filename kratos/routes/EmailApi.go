package routes

import (
	"github.com/gin-gonic/gin"
	"model/kratos/controllers/emailController"
)

func EmailRoutes(r *gin.Engine) {
	r.Group("/email")
	{
		r.GET("/send", emailController.EmailSend)
	}
}
