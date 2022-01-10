package model

import (
	"context"
	"learnGrpc/proto/stream"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
)

// StreamModel 双向数据流模式 client and server can send more data
func StreamModel() {
	log.Println("get con by ip:port")
	con, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	defer con.Close()

	log.Println("new client by con")
	cli := stream.NewStreamClient(con)

	log.Println("call remote method")
	rsp, err := cli.Chat(context.Background())
	if err != nil {
		log.Printf("err = %+v\n", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			err := rsp.Send(&stream.Req{
				Name: "i am client",
			})
			if err != nil {
				log.Printf("err = %+v\n", err)
				break
			}

			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			data, err := rsp.Recv()
			if err != nil {
				log.Printf("err = %+v\n", err)
				break
			}

			log.Printf("data = %+v\n", data)
		}
	}()

	wg.Wait()
}
