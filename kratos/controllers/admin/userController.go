package admin

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"model/kratos/models"
	information "model/kratos/struct/Information"
	"net/http"
	"os"
	"path"
	"strconv"
)

type UserController struct {
	BaseController
}

func (con UserController) UserIndex(content *gin.Context) {
	{
		content.HTML(http.StatusOK, "show.html", gin.H{"title": "我是测试", //gin.H就是map[string]interface{}，interface{}表示任意类型
			"address": "www.5lmh.com"})
	}
}

func Upload(content *gin.Context) {
	//c.String(200, "ok", "is")
	content.HTML(http.StatusOK, "upload.html", gin.H{"title": "我是测试",
		"address": "www.5lmh.com"})
}

func UploadSameName(content *gin.Context) {
	//c.String(200, "ok", "is")
	content.HTML(http.StatusOK, "uploadSameName.html", gin.H{"title": "我是测试",
		"address": "www.5lmh.com"})
}

func Jsonp(content *gin.Context) {
	data := &information.PersonInfo{
		User_id:  4,
		Username: "JohnMike",
		Age:      30,
		Address:  "123 Street",
		Sex:      "Male",
		Work:     "Engineer",
		Email:    "john@example.com",
	}
	content.JSONP(http.StatusOK, data)
}

func UploadData(context *gin.Context) {
	file, err := context.FormFile("file1")
	if err != nil {
		context.String(http.StatusInternalServerError, "读取file失败: "+err.Error())
		return
	}
	//defer file.Close()
	Username := context.PostForm("Username")
	Age, _ := strconv.Atoi(context.PostForm("Age"))
	Address := context.PostForm("Address")
	Sex := context.PostForm("Sex")
	Work := context.PostForm("Work")
	Email := context.PostForm("Email")

	allInfo := information.PersonInfo{
		Username: Username,
		Age:      Age,
		Address:  Address,
		Sex:      Sex,
		Work:     Work,
		Email:    Email,
	}
	dst := path.Join("../../images/person/", file.Filename) //路径基于main.go位置
	err = context.SaveUploadedFile(file, dst)
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, allInfo)

}

func UploadSameNameData(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		context.String(http.StatusInternalServerError, "读取file失败: "+err.Error())
		return
	}
	file := form.File["file[]"]
	for _, file := range file {
		dst := path.Join("../../images/person/", file.Filename) //路径基于main.go位置
		err = context.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}
	}
	//defer file.Close()
	context.JSON(http.StatusOK, "success")

}

func GetValue(context *gin.Context) {
	username := context.Query("username")
	sex := context.DefaultQuery("sex", "男")
	context.JSONP(http.StatusOK, gin.H{
		"username": username,
		"sex":      sex,
	})
}

func PostValue(context *gin.Context) {

	username := context.PostForm("username")
	sex := context.DefaultPostForm("sex", "女")
	context.JSONP(http.StatusOK, gin.H{
		"username": username,
		"sex":      sex,
	})
}

func GetBindStruct(context *gin.Context) {
	user := information.GetBindStruct{}
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, user)
	}
}

func GetXml(context *gin.Context) {
	sliceData, _ := context.GetRawData()
	xmlStruct := &information.GetXml{}
	if err := xml.Unmarshal(sliceData, xmlStruct); err == nil {
		context.JSON(http.StatusOK, xmlStruct)
	} else {
		context.JSON(http.StatusBadRequest, err.Error())
	}

}

func GetDynamicRoute(context *gin.Context) {
	cid := context.Param("cid")
	context.JSON(http.StatusOK, cid)
}

func (con UserController) Inherit(ctx *gin.Context) {
	con.BaseController.success(ctx)
}

func (con UserController) UoloadByDate(context *gin.Context) {
	file, err := context.FormFile("file")
	if err == nil {
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		if _, err := allowExtMap[extName]; !err {
			context.String(200, "上传的文件不合法")
			return
		}
		day := models.GetDate()

		dir := "../images" + day
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			context.String(200, "MkdirAll失败")
			return
		}

		FileName := models.GetDay() + extName
		dst := path.Join(dir, FileName)
		err = context.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"dst":     dst,
		})
	}
}

func (con UserController) GetMysqlDate(context *gin.Context) {
	ItTeacherInfoList := []models.ItTeacherInfo{}
	models.DB.First(&ItTeacherInfoList)
	models.ConvTimeStamps(&ItTeacherInfoList) //传入指针，修改时间戳不带区域时间
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ItTeacherInfoList,
	})
}
