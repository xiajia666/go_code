package main

import (
	"context"
	"google.golang.org/grpc"
	pb "model/account/grpc/server/proto"
	"net"
	"strconv"
)

type server struct {
	pb.UnimplementedGetTeacherInfoServer
}

func (s *server) GetTeacherInfo(ctx context.Context, req *pb.RequestInfo) (*pb.ResponseInfo, error) {
	return &pb.ResponseInfo{Teacherinfo: []string{strconv.Itoa(int(req.Id)), req.Name, req.Worknum}}, nil
}

func main() {
	// 开启端口
	lis, _ := net.Listen("tcp", ":8080")

	// 创建grpc服务
	grpcServer := grpc.NewServer()

	// 在grpc服务中注册服务
	pb.RegisterGetTeacherInfoServer(grpcServer, &server{})

	// 启动服务
	err := grpcServer.Serve(lis)
	if err != nil {
		return
	}
}
