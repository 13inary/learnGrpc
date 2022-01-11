package auth

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

// package auth

type clientAuth struct {
}

func (c clientAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"token": "xxxx",
	}, nil

}

func (c clientAuth) RequireTransportSecurity() bool {
	return false
}

// client
func client() {
	// equals metadata + inteceptor
	opt := grpc.WithPerRPCCredentials(clientAuth{})

	con, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), opt)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	fmt.Printf("con = %+v\n", con)

}
