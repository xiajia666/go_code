package admin

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	information "model/kratos/struct/Information"
	"net/http"
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
	file, err := context.FormFile("file")
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
	err = context.SaveUploadedFile(file, "../../images/person/"+file.Filename)
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, allInfo)

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
