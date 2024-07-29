package routes

import (
	"github.com/gin-gonic/gin"
	"model/kratos/controllers/emailController"
)

func EmailRoutes(r *gin.Engine) {
	email := r.Group("/email")
	{
		email.GET("/send", emailController.EmailSend)
	}
}
