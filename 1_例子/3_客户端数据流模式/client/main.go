package main

import (
	"client/pb"
	"context"
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

	req, err := cli.PutStream(context.Background())
	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		req.Send(&pb.Req{Name: "client"})
		time.Sleep(time.Second)
	}
}
