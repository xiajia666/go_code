package routes

import (
	"github.com/gin-gonic/gin"
	"model/kratos/controllers/teacherController"
)

func TeacherInfoRoutes(r *gin.Engine) {
	r.Group("/teacherController")
	{
		r.POST("/update", teacherController.Update)
		r.POST("/search", teacherController.Search)
		r.POST("/save", teacherController.Save)
		r.POST("/delete", teacherController.Delete)
	}
}
