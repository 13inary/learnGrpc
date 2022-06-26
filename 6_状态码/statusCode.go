package statusCode

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// package statusCode
// github.com/grpc/grpc/blob/master/doc/statuscodes.md

// set 服务端返回错误
func set() error {
	return status.New(codes.InvalidArgument, "error")

	return status.Error(codes.InvalidArgument, "error")

	a := 10
	return status.Errorf(codes.InvalidArgument, "error%v", a)
}

// get 客户端接受错误
func get(err error) {
	// 调用远程函数

	// 检查err
	statu, ok := status.FromError(err)
	if !ok {
		//
	}
	fmt.Printf("statu.Message() = %+v\n", statu.Message())
	fmt.Printf("statu.Code() = %+v\n", statu.Code())
}
