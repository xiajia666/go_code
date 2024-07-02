package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"model/kratos/cmd/kratos/internal/change"
	"model/kratos/cmd/kratos/internal/project"
	"model/kratos/cmd/kratos/internal/proto"
	"model/kratos/cmd/kratos/internal/run"
	"model/kratos/cmd/kratos/internal/upgrade"
	"model/kratos/middleware/customize"
	"model/kratos/models"
	"model/kratos/routes"
	"text/template"
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

func loadRoutes(r *gin.Engine) {
	routes.DefaultRoutes(r)
	routes.ApiRoutes(r)
	routes.TeacherInfoRoutes(r)
	routes.ArticleInfoRoutes(r)
	routes.EmailRoutes(r)
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"TimeToUnix": models.TimeToUnix,
	})
}

func loadMiddleware(r *gin.Engine) {
	r.Use(customize.GolbalMiddleWare) //全局中间件，可配置多个
}

func main() {
	r := gin.Default() //默认引擎自带两个中间件
	loadRoutes(r)      //加载路由
	loadMiddleware(r)  //加载中间件
	r.Run(":8080")     // 运行 HTTP 服务器
}
