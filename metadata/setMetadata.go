package metadata

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/metadata"
)

// package metadata

// SetMetadata
func SetMetadata() {
	log.Println("create medatata : way 1")
	md := metadata.New(map[string]string{
		"key1": "val1",
		"key2": "val2",
	})
	fmt.Printf("md = %+v\n", md)

	log.Println("create metadata : way 2 all big case will be transform to low case")
	md2 := metadata.Pairs(
		"KEY1", "value1",
		"key2", "VALUE2",
	)
	fmt.Printf("md2 = %+v\n", md2)

	log.Println("new context which include metadata")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	fmt.Printf("ctx = %+v\n", ctx)

	log.Println("send metadata")
	//rsp, err := client.SomeRPC(ctx, someReq)
}
