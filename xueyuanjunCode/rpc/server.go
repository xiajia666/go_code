package rpc

import (
	"errors"
	"log"
	"model/xueyuanjunCode/utils"
	"net"
	"net/http"
	"net/rpc"
)

type MathService struct {
}

func (m *MathService) Multiply(args *utils.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (m *MathService) Divide(args *utils.Args, reply *int) error {
	if args.B == 0 {
		return errors.New("除数不能为0")
	}
	*reply = args.A / args.B
	return nil
}

func main() {
	// 启动 RPC 服务端
	math := new(MathService)
	rpc.Register(math)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("启动服务监听失败:", err)
	}
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("启动 HTTP 服务失败:", err)
	}
}
