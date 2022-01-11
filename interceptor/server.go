package interceptor

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

// package interceptor

// server
func server() {
	myInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (rsp interface{}, err error) {
		log.Println("recevive a new req")

		// here can handle metadata
		// return nil, status.Error(codes.Unauthenticated, "not have token message")

		rsp, err = handler(ctx, req)

		log.Println("req is END")
		return rsp, err
	}
	opt := grpc.UnaryInterceptor(myInterceptor)

	g := grpc.NewServer(opt)

	fmt.Printf("g = %+v\n", g)
}
