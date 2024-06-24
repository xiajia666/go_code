package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "model/account/grpc/client/proto"
)

func main() {
	// 连接到server端，创建grpc连接，此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建客户端，建立连接
	client := pb.NewGetTeacherInfoClient(conn)
	resp, _ := client.GetTeacherInfo(context.Background(), &pb.RequestInfo{Id: 1, Name: "程子韬", Worknum: "xiajia666"})
	fmt.Println(resp.Teacherinfo)
}
