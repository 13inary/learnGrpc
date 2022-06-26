package main

import (
	"context"
	"fmt"
	"server/pb"

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

	rsp, err := cli.GetStream(context.Background(), &pb.Req{Name: "client"})
	if err != nil {
		return
	}
	// 不断接受服务端的返回
	for {
		one, err := rsp.Recv()
		if err != nil {
			return
		}
		fmt.Println(one)
	}
}
