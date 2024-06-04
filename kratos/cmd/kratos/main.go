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
	r := gin.Default()                //默认引擎自带两个中间件
	r.Use(customize.GolbalMiddleWare) //全局中间件，可配置多个
	routes.DefaultRoutes(r)
	routes.ApiRoutes(r)
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"TimeToUnix": models.TimeToUnix,
	})

	// 运行 HTTP 服务器
	r.Run(":8080")

}
