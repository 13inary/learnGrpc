package main

import (
	"net"
	"server/pb"
	"time"

	"google.golang.org/grpc"
)

// package main

// Server 注意参数和返回值 和简单模式不一样
type Server struct {
}

func (s *Server) GetStream(req *pb.Req, rsp pb.Hello_GetStreamServer) error {
	// 不断返回信息给客户端
	for i := 0; i < 10; i++ {
		err := rsp.Send(&pb.Rsp{Info: "server stream"})
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

// main
func main() {
	g := grpc.NewServer()

	pb.RegisterHelloServer(g, &Server{})

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		return
	}
	err = g.Serve(lis)
	if err != nil {
		return
	}
}
