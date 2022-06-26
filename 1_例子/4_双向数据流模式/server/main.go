package main

import (
	"all/pb"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

// package main

// Server 注意参数和返回值 和简单模式不一样
type Server struct {
}

func (s *Server) AllStream(allStr pb.Hello_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			data, err := allStr.Recv()
			if err != nil {
				return
			}
			fmt.Println(data)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			err := allStr.Send(&pb.Rsp{Info: "server"})
			if err != nil {
				return
			}
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
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
