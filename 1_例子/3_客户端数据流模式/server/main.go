package main

import (
	"client/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// package main

// Server 注意参数和返回值 和简单模式不一样
type Server struct {
}

func (s *Server) PutStream(cliStr pb.Hello_PutStreamServer) error {
	for {
		data, err := cliStr.Recv()
		if err != nil {
			fmt.Println("客户端流调用结束")
			return nil
		}
		fmt.Println(data)
	}
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
	fmt.Println("a")
}
