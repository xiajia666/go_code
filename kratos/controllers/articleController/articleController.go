package articleController

import (
	"github.com/gin-gonic/gin"
	"model/kratos/internal/Struct"
	"model/kratos/internal/mysql"
	"net/http"
	"strconv"
)

func GetInfo(context *gin.Context) {
	currentPage, _ := strconv.Atoi(context.DefaultQuery("currentPage", "0"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "1"))
	offset := (currentPage - 1) * pageSize
	var article []Struct.Article
	mysql.DBArticle.Preload("ArticleCategory").Offset(offset).Limit(pageSize).Find(&article)
	context.JSON(http.StatusOK, gin.H{"status": "200", "message": article})
}

func Save(content *gin.Context) {
	articleName := content.Query("articleName")
	categoryId, _ := strconv.Atoi(content.Query("categoryId"))
	//id, _ := strconv.Atoi(content.Query("id"))
	info := Struct.Article{
		//Id:     id,
		ArticleName: articleName,
		CateId:      categoryId,
	}
	mysql.DBArticle.Model(&Struct.Article{}).Create(&info)
	content.JSON(http.StatusOK, gin.H{"status": "200", "message": "保存成功"})
}

func Delete(context *gin.Context) {
	articleName := context.Query("articleName")
	info := Struct.Article{
		ArticleName: articleName,
	}
	mysql.DBArticle.Model(&Struct.Article{}).Where("articleName = ?", articleName).Delete(info)
	context.String(http.StatusOK, "", gin.H{"status": "200", "message": "删除成功"})
}

func Update(content *gin.Context) {
	id := content.Query("id")
	articleName := content.Query("articleName")
	cateId, _ := strconv.Atoi(content.Query("CateId"))
	info := Struct.Article{
		ArticleName: articleName,
		CateId:      cateId,
	}
	mysql.DBArticle.Model(&info).Where("id = ?", id).Updates(info) // 多个更新
	content.JSON(http.StatusOK, gin.H{"status": "200", "message": "更新成功"})
}

func Search(content *gin.Context) {
	name := content.Query("articleName")
	var info []Struct.Article
	mysql.DBArticle.Model(&Struct.Article{}).Where("article_name = ?", name).Find(&info)
	content.JSON(http.StatusOK, gin.H{"status": "200", "message": info})
}

func SearchBy(context *gin.Context) {
	value := context.Query("value")
	var info []Struct.ArticleCategory
	mysql.DBArticle.Preload("Articles").Where("id = ?", value).Find(&info)
	context.JSON(http.StatusOK, gin.H{"status": "200", "message": info})
}
