package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"model/kratos/controllers/admin"
	"model/kratos/middleware/customize"
	information "model/kratos/struct/Information"
	"net/http"
)

func DefaultRoutes(r *gin.Engine) {
	r.Group("/")
	{
		// 加载、返回html
		r.LoadHTMLGlob("../../html/*")
		r.GET("/upload", admin.Upload)
		// 上传并保存文件
		r.POST("/upload", admin.UploadData)

		// 上传相同文件名的文件
		r.GET("/uploadsamename", admin.UploadSameName)
		// 上传并保存文件
		r.POST("/uploadsamename", admin.UploadSameNameData)

		r.GET("/show", admin.UserController{}.UserIndex)

		// 回调函数 // http://localhost:8080/jsonp?cal1back=xxxx 执行前端的函数xxx
		// xxxx("User_id":  4,"Username": "JohnMike", "Age":30, "Address": "123 Street", "Sex":"Male", "Work":"Engineer", "Email":"john@example.com")})；
		r.GET("/jsonp", admin.Jsonp)

		// 返回xml
		r.GET("xml", customize.InitMiddleWareOne, customize.InitMiddleWareTwo, func(context *gin.Context) { //中间件执行顺序   1-我是中间件---->3-我是中间件---->我要返回xml数据---->4-我是中间件---->   2-我是中间件
			context.XML(http.StatusOK, information.GetXml{
				Title:   "返回的数据",
				Content: "true",
			})
			fmt.Println("我要返回xml数据")
			username, _ := context.Get("username") //来自全局中间件的值
			v, _ := username.(string)              // 类型断言，将username由interdace{}转化为string
			fmt.Println(v)
		})

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
		r.GET("/mysqldate", admin.UserController{}.GetMysqlDate)
		r.GET("/getcookies", admin.UserController{}.GetCookies)

	}
}
