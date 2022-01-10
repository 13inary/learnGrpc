package main

import (
	"fmt"
	"learnGrpc/proto/hello"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	req := hello.HelloReq{
		Name: "hello proto",
	}

	rsp, err := proto.Marshal(&req)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}

	fmt.Printf("rsp = %+v\n", rsp)
}
