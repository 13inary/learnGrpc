package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

// package interceptor

func myInterceptor2(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	// 这里可以添加metadata

	// 执行请求 远程函数
	err := invoker(ctx, method, req, reply, cc, opts...)

	// 处理返回值

	return err
}

// client
func client() {
	opt := grpc.WithUnaryInterceptor(myInterceptor2)

	// 做法1
	con, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), opt)

	// 做法2
	// var opts []grpc.DialOption
	// opts = append(opts, grpc.WithInsecure())
	// opts = append(opts, WithUnaryInterceptor(myInterceptor))
	// con, err := grpc.Dial("127.0.0.1:8080", opts...)
}
