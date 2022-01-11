package statusCode

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// package statusCode

// set
func set() error {
	if true {
		return status.Error(codes.InvalidArgument, "error")
	}

	a := 10
	return status.Errorf(codes.InvalidArgument, "error%v", a)
}

// get
func get(err error) {
	statu, ok := status.FromError(err)
	if !ok {
		//
	}
	fmt.Printf("statu.Message() = %+v\n", statu.Message())
	fmt.Printf("statu.Code() = %+v\n", statu.Code())
}
