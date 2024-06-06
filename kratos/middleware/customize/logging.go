package customize

import (
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentDate := time.Now().Format("2006-01-02")
		err := os.MkdirAll("../../logs/"+currentDate, 0755)
		if err != nil {
			return
		}
		//
		//if _, err := os.Stat("../../logs/" + currentDate + "/logs.txt"); err != nil {
		//	file, err := os.Create("../../logs/" + currentDate + "/logs.txt")
		//	if err != nil {
		//		return
		//	}
		//	// 关闭文件
		//	defer file.Close()
		//}
		openFile, err := os.OpenFile("../../logs/"+currentDate+"/logs.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return
		}
		openFile.WriteString("dddkk\n")
		defer openFile.Close()

	}
}
