package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

// package interceptor

func myInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (rsp interface{}, err error) {
	fmt.Println("收到一个请求")

	// 这里可以对metadata进行处理
	// return nil, status.Error(codes.Unauthenticated, "not have token message")

	// 这里是执行 远程函数
	rsp, err = handler(ctx, req)

	// 这里可以对响应进行修改

	fmt.Println("请求结束")
	return rsp, err
}

// server
func server() {
	opt := grpc.UnaryInterceptor(myInterceptor)

	// 新建服务端，可以加多个拦截器
	g := grpc.NewServer(opt)
}
