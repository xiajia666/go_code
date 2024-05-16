package routes

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	information "model/kratos/struct/Information"
	"net/http"
	"strconv"
)

func DefaultRoutes(r *gin.Engine) {
	r.Group("/")
	{
		// 加载、返回html
		r.LoadHTMLGlob("../../html/*")
		r.GET("/upload", func(c *gin.Context) {
			//c.String(200, "ok", "is")
			c.HTML(http.StatusOK, "upload.html", gin.H{"title": "我是测试",
				"address": "www.5lmh.com"})
		})
		r.GET("/show", func(c *gin.Context) {
			c.HTML(http.StatusOK, "show.html", gin.H{"title": "我是测试", //gin.H就是map[string]interface{}，interface{}表示任意类型
				"address": "www.5lmh.com"})
		})

		// 回调函数 // http://localhost:8080/jsonp?cal1back=xxxx 执行前端的函数xxx
		// xxxx("User_id":  4,"Username": "JohnMike", "Age":30, "Address": "123 Street", "Sex":"Male", "Work":"Engineer", "Email":"john@example.com")})；
		r.GET("jsonp", func(context *gin.Context) {
			a := &information.PersonInfo{
				User_id:  4,
				Username: "JohnMike",
				Age:      30,
				Address:  "123 Street",
				Sex:      "Male",
				Work:     "Engineer",
				Email:    "john@example.com",
			}
			context.JSONP(http.StatusOK, a)
		})

		// 返回xml
		r.GET("xml", func(context *gin.Context) {
			context.XML(http.StatusOK, map[string]interface{}{
				"status":  true,
				"message": "返回的数据",
			})
		},
		)

		//r.GET("/show", func(c *gin.Context) {
		//	c.HTML(http.StatusOK, "show.html", gin.H{})
		//})
		//
		////添加路由处理程序
		//r.GET("/users", func(c *gin.Context) {
		//	//var users []User
		//	//// 查询所有用户
		//	//db.Find(&users)
		//	c.JSON(200, "ososo")
		//})

		// 上传并保存文件
		r.POST("/upload", func(context *gin.Context) {
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
		})

		//获取get传值---username、sex
		r.GET("/getVaule", func(context *gin.Context) {
			username := context.Query("username")
			sex := context.DefaultQuery("sex", "男")
			context.JSONP(http.StatusOK, gin.H{
				"username": username,
				"sex":      sex,
			})
		})

		//获取post传值---
		r.POST("/postVaule", func(context *gin.Context) {
			username := context.PostForm("username")
			sex := context.DefaultPostForm("sex", "女")
			context.JSONP(http.StatusOK, gin.H{
				"username": username,
				"sex":      sex,
			})
		})

		// 获取GET、POST数据绑定到结构体
		r.GET("/getBingStruct", func(context *gin.Context) {
			user := information.GetBindStruct{}
			if err := context.ShouldBind(&user); err != nil {
				context.JSON(http.StatusOK, gin.H{
					"err": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, user)
			}
		})

		// 获取GET、POST中的xml数据
		r.POST("/getXml", func(context *gin.Context) {
			sliceData, _ := context.GetRawData()
			xmlStruct := &information.GetXml{}
			if err := xml.Unmarshal(sliceData, xmlStruct); err == nil {
				context.JSON(http.StatusOK, xmlStruct)
			} else {
				context.JSON(http.StatusBadRequest, err.Error())
			}

		})

		// 动态路由配置
		r.GET("/getDynamicRoute/:cid", func(context *gin.Context) {
			cid := context.Param("cid")
			context.JSON(http.StatusOK, cid)
		})

	}
}
