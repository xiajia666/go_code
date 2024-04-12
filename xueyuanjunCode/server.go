package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		params := request.URL.Query()
		fmt.Fprintf(writer, "你好, %s", params.Get("name"))
	})
	err := http.ListenAndServeTLS(":8443", "xueyuanjunCode/ca.crt", "xueyuanjunCode/ca.key", nil)
	if err != nil {
		log.Fatalf("启动 HTTPS 服务器失败: %v", err)
	}
}
