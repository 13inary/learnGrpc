package main

import (
	"context"
	"net"
	"simple/pb"

	"google.golang.org/grpc"
)

// package main

// Server 远程函数的 实现数据类型
// 所有远程函数都要添加 context.Context 和 error
// req 和 rsp 数据类型必须是指针
type Server struct {
}

func (s *Server) Say(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	return &pb.Rsp{Info: req.Name}, nil
}

// main
func main() {
	// 新建服务端
	g := grpc.NewServer()

	// 注册服务（注册远程函数）
	// proto文件service后面的服务名 是该函数的一部分
	pb.RegisterHelloServer(g, &Server{})

	// 运行服务
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		return
	}
	err = g.Serve(lis)
	if err != nil {
		return
	}
}
