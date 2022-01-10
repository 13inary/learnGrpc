package model

import (
	"context"
	"fmt"
	"learnGrpc/proto/hello"
	"log"

	"google.golang.org/grpc"
)

// SimplieModel 简单模式 client send one reqest and server send one response
func SimplieModel() {
	log.Println("get con by ip:port")
	con, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	defer con.Close()

	log.Println("new client by con")
	cli := hello.NewHelloClient(con)

	log.Println("call remote method")
	rsp, err := cli.Hello(context.Background(), &hello.HelloReq{
		Name: "i am hello req",
	})
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	fmt.Printf("rsp = %+v\n", rsp)
}
