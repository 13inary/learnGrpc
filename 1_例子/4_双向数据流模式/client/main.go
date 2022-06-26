package main

import (
	"all/pb"
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
)

// package main

// main
func main() {
	con, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer con.Close()

	cli := pb.NewHelloClient(con)

	allStr, err := cli.AllStream(context.Background())
	if err != nil {
		return
	}

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
			err := allStr.Send(&pb.Req{Name: "client"})
			if err != nil {
				return
			}
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
