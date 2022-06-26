package metadata

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

// package metadata

// 远程函数
type server struct {
}

func (s *server) SomeRPC(ctx context.Context, in *pb.SomeReq) (*pb.SomeRsp, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		// 处理错误
	}
	for k, v := range md {
		fmt.Println(k)
		fmt.Println(v)
	}
	if token, ok := md["token"]; !ok {
	}

	return nil, nil
}
