package customize

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 局部中间件
func InitMiddleWareOne(context *gin.Context) {
	//start := time.Now().Unix()
	//time.Sleep(time.Second * 5)
	fmt.Println("1-我是中间件")
	context.Next() // 调用当前程序的剩余程序，最后再由后向前执行
	//context.Abort() // 终止当其他程序
	fmt.Println("2-我是中间件")
	//end := time.Now().Unix()
	//fmt.Println(end - start)
}

func InitMiddleWareTwo(context *gin.Context) {
	//start := time.Now().Unix()
	//time.Sleep(time.Second * 5)
	fmt.Println("3-我是中间件")
	context.Next() // 调用当前程序的剩余程序，最后再由后向前执行
	//context.Abort() // 终止其他程序
	fmt.Println("4-我是中间件")
	//end := time.Now().Unix()
	//fmt.Println(end - start)
}

// 全局中间件
func GolbalMiddleWare(context *gin.Context) {
	fmt.Println("我是全局中间件")
	context.Set("username", "我叫张三")
	contextCopy := context.Copy()
	go func() { //协程，当在中间件中执行协程时候必须把context复制一份再使用
		fmt.Println(contextCopy.Request.URL.Path)
	}()
}
