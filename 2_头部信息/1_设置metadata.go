package metadata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// package metadata

// SetMetadata key最终都转化为小写
func SetMetadata() {
	// 方式1
	md := metadata.New(map[string]string{
		"key1": "val1",
		"key2": "val2",
	})

	// 方式2
	md2 := metadata.Pairs(
		"KEY1", "value1",
		"KEY1", "value2",
		"KEY1", "value3", // 会组成[]string{value1, value2, value3}
		"key2", "VALUE2",
	)

	// 加入 context 中
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 在调用 远程方法的时候会把 ctx作为参数
	//rsp, err := client.SomeRPC(ctx, someReq)
}
