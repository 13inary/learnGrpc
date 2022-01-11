package interceptor

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

// package interceptor

// client
func client() {
	myInterceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		// here can add metadata

		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
	opt := grpc.WithUnaryInterceptor(myInterceptor)

	// var opts []grpc.DialOption
	// opts = append(opts, grpc.WithInsecure())
	// opts = append(opts, WithUnaryInterceptor(myInterceptor))
	con, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), opt)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	fmt.Printf("con = %+v\n", con)
}
