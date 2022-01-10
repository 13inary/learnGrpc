package model

import (
	"context"
	"learnGrpc/proto/stream"
	"log"
	"time"

	"google.golang.org/grpc"
)

// ClientStreamModel 客户端数据流模式 client send more request and server send one respose
func ClientStreamModel() {
	log.Println("get con by ip:port")
	con, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	defer con.Close()

	log.Println("new client by con")
	cli := stream.NewStreamClient(con)

	log.Println("call remote method")
	rsp, err := cli.Put(context.Background())
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	for i := 1; i <= 10; i++ {
		err := rsp.Send(&stream.Req{
			Name: "i am client",
		})
		if err != nil {
			log.Printf("err = %+v\n", err)
			break
		}

		time.Sleep(time.Second)
	}

}
