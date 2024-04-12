package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "https://127.0.0.1:8443/hello?name=学院君", nil)
	if err != nil {
		fmt.Printf("请求初始化失败：%v", err)
		return
	}

	// 设置跳过不安全的 HTTPS
	tls11Transport := &http.Transport{
		MaxIdleConnsPerHost: 10,
		TLSClientConfig: &tls.Config{
			MaxVersion:         tls.VersionTLS12,
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: tls11Transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("客户端发起请求失败：%v", err)
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
