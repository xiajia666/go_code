package rpc

import (
	"fmt"
	"log"
	"model/xueyuanjunCode/utils"
	"net/rpc"
)

func main() {
	var serverAddress = "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress+":8080")
	if err != nil {
		log.Fatal("建立与服务端连接失败:", err)
	}
	args := &utils.Args{A: 10, B: 10}
	var reply int
	err = client.Call("MathService.Multiply", args, &reply)
	if err != nil {
		log.Fatal("调用远程方法 MathService.Multiply 失败:", err)
	}
	fmt.Printf("%d*%d=%d\n", args.A, args.B, reply)
}
