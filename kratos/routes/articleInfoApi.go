package routes

import (
	"github.com/gin-gonic/gin"
	"model/kratos/controllers/articleController"
)

func ArticleInfoRoutes(r *gin.Engine) {
	r.GET("/articleInfo/getInfo", articleController.GetInfo)
	r.POST("/articleInfo/save", articleController.Save)
	r.POST("/articleInfo/update", articleController.Update)
	r.POST("/articleInfo/delete", articleController.Delete)
	r.POST("/articleInfo/search", articleController.Search)
	r.POST("/articleInfo/searchBy", articleController.SearchBy)
}
