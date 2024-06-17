package teacherController

import (
	"github.com/gin-gonic/gin"
	"model/kratos/internal/Struct"
	"model/kratos/internal/mysql"
	"net/http"
	"strconv"
)

func Save(content *gin.Context) {
	name := content.Query("username")
	sex := content.DefaultQuery("sex", "男")
	id, _ := strconv.Atoi(content.Query("id"))
	info := Struct.ItTeacherInfo{
		Id:     id,
		Name:   name,
		Gender: sex,
	}
	mysql.DBTeacherInfo.Model(&Struct.ItTeacherInfo{}).Create(&info)
	content.String(http.StatusOK, "", gin.H{"status": "200", "message": "保存成功"})
}

func Delete(context *gin.Context) {
	id, _ := strconv.Atoi(context.Query("id"))
	info := Struct.ItTeacherInfo{
		Id: id,
	}
	mysql.DBTeacherInfo.Model(&Struct.ItTeacherInfo{}).Where("id = ?", id).Delete(info)
	context.String(http.StatusOK, "", gin.H{"status": "200", "message": "删除成功"})

}

func Update(content *gin.Context) {
	name := content.Query("username")
	sex := content.DefaultQuery("sex", "男")
	id, _ := strconv.Atoi(content.Query("id"))
	info := Struct.ItTeacherInfo{
		Id:     id,
		Name:   name,
		Gender: sex,
	}
	mysql.DBTeacherInfo.Model(&info).Where("id = ?", id).Updates(info)        // 多个更新
	mysql.DBTeacherInfo.Model(&info).Where("id = ?", id).Update("Name", name) // 单个更新
	content.String(http.StatusOK, "", gin.H{"status": "200", "message": "更新成功"})
}

func Search(content *gin.Context) {
	name := content.Query("username")
	id := content.Query("id")
	intId, _ := strconv.Atoi(content.Query("id"))
	searchMode := content.DefaultQuery("type", "id")
	info := Struct.ItTeacherInfo{
		Id:   intId,
		Name: name,
	}
	switch searchMode {
	case id:
		mysql.DBTeacherInfo.Model(&Struct.ItTeacherInfo{}).Where("id = ?", id).First(&info)
	case name:
		mysql.DBTeacherInfo.Model(&Struct.ItTeacherInfo{}).Where("name = ?", name).First(&info)
	default:
		mysql.DBTeacherInfo.Model(&Struct.ItTeacherInfo{}).Where("id = ? or name = ?", id, name).First(&info)
	}
}
func GetInfo(context *gin.Context) {
	id := context.Query("id")
	var teacherInfo Struct.ItTeacherInfo
	mysql.DBTeacherInfo.First(&teacherInfo, id)
	context.JSON(http.StatusOK, gin.H{"status": "200", "message": teacherInfo, "id": id})
}
