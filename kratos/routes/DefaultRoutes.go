package routes

import (
	"github.com/gin-gonic/gin"
	"model/kratos/controllers/admin"
	"net/http"
)

func DefaultRoutes(r *gin.Engine) {
	r.Group("/")
	{
		// 加载、返回html
		r.LoadHTMLGlob("../../html/*")
		r.GET("/upload", admin.Upload)
		r.GET("/show", admin.UserController{}.UserIndex)

		// 回调函数 // http://localhost:8080/jsonp?cal1back=xxxx 执行前端的函数xxx
		// xxxx("User_id":  4,"Username": "JohnMike", "Age":30, "Address": "123 Street", "Sex":"Male", "Work":"Engineer", "Email":"john@example.com")})；
		r.GET("/jsonp", admin.Jsonp)

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
		r.POST("/upload", admin.UploadData)

		//获取get传值---username、sex
		r.GET("/getVaule", admin.GetValue)

		//获取post传值---
		r.POST("/postVaule", admin.PostValue)

		// 获取GET、POST数据绑定到结构体
		r.GET("/getBindStruct", admin.GetBindStruct)

		// 获取GET、POST中的xml数据
		r.POST("/getXml", admin.GetXml)

		// 动态路由配置
		r.GET("/getDynamicRoute/:cid", admin.GetDynamicRoute)
		r.GET("/inherit", admin.UserController{}.Inherit)
	}
}
