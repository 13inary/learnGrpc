package model

import (
	"context"
	"learnGrpc/proto/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) Hello(ctx context.Context, req *hello.HelloReq) (rsp *hello.HelloRsp, err error) {
	return &hello.HelloRsp{
		Info: "" + req.Name,
	}, nil
}

// SimplieModel 简单模式 client send one reqest and server send one response
func SimplieModel() {
	log.Println("listen port")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Printf("err = %+v\n", err)
	}

	log.Println("new server and register method")
	g := grpc.NewServer()
	hello.RegisterHelloServer(g, &Server{})

	log.Println("run server")
	err = g.Serve(lis)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
}
