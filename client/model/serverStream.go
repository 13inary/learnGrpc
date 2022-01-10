package model

import (
	"context"
	"learnGrpc/proto/stream"
	"log"

	"google.golang.org/grpc"
)

// ServerStreamModel 服务端数据流模 client send one request and server send more response
func ServerStreamModel() {
	log.Println("get con by ip:port")
	con, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	defer con.Close()

	log.Println("new client by con")
	cli := stream.NewStreamClient(con)

	log.Println("call remote method")
	rsp, err := cli.Get(context.Background(), &stream.Req{
		Name: "i am client",
	})
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	for {
		data, err := rsp.Recv()
		if err != nil {
			log.Printf("err = %+v\n", err)
			break
		}

		log.Printf("data = %+v\n", data)
	}
}
