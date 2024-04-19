package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"model/kratos/cmd/kratos/internal/change"
	"model/kratos/cmd/kratos/internal/project"
	"model/kratos/cmd/kratos/internal/proto"
	"model/kratos/cmd/kratos/internal/run"
	"model/kratos/cmd/kratos/internal/upgrade"
	"model/kratos/struct/Information"
	"net/http"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:     "kratos",
	Short:   "Kratos: An elegant toolkit for Go microservices.",
	Long:    `Kratos: An elegant toolkit for Go microservices.`,
	Version: "2.7.2",
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(change.CmdChange)
	rootCmd.AddCommand(run.CmdRun)
}

//func customMiddleware(handler middleware.Handler) middleware.Handler {
//	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
//		if tr, ok := transport.FromServerContext(ctx); ok {
//			fmt.Println("operation:", tr.Operation())
//		}
//		reply, err = handler(ctx, req)
//		return
//	}
//}

func hello() string {
	return "sb"
}

//func hello(c http.Context) string {
//	return "sb"
//}

func main() {

	//
	//app := kratos.New(
	//	kratos.Name("cors"),
	//	kratos.Server(
	//		httpSrv,
	//	),
	//)
	//if err := app.Run(); err != nil {
	//	log.Fatal(err)
	//}
	//if err := app.Run(); err != nil {
	//	log.Fatal(err)
	//}

	db, err := gorm.Open(mysql.Open("root:985211@tcp(118.25.191.214:3306)/Information"), &gorm.Config{})
	if db == nil && err != nil {
		panic("failed to connect")
	}

	var allInfo = information.PersonInfo{
		User_id:  1,
		Username: "John",
		Age:      30,
		Address:  "123 Street",
		Sex:      "Male",
		Work:     "Engineer",
		Email:    "john@example.com",
	}
	db.Create(&allInfo)

	r := gin.Default()
	r.LoadHTMLFiles("../../html/upload.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})
	//r.GET("/show", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "show.html", gin.H{})
	//})

	// 添加路由处理程序
	//r.GET("/users", func(c *gin.Context) {
	//	//var users []User
	//	//// 查询所有用户
	//	//db.Find(&users)
	//	c.JSON(200, "ososo")
	//})
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

		allInfo = information.PersonInfo{
			Username: Username,
			Age:      Age,
			Address:  Address,
			Sex:      Sex,
			Work:     Work,
			Email:    Email,
		}
		db.Create(allInfo)
		err = context.SaveUploadedFile(file, "../../images/person/"+file.Filename)
		if err != nil {
			return
		}
	})

	// 运行 HTTP 服务器
	r.Run(":8080")

}
