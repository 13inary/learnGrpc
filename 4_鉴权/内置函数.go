package auth

import (
	"context"

	"google.golang.org/grpc"
)

// package auth

type addToken struct {
}

func (c addToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"token": "xxxx",
	}, nil

}

func (c addToken) RequireTransportSecurity() bool {
	return false
}

// client
func client() {
	// 等同于 metadata + inteceptor
	opt := grpc.WithPerRPCCredentials(addToken{})

	con, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), opt)
}
