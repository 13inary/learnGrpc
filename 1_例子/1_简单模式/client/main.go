package main

import (
	"context"
	"fmt"
	"simple/pb"

	"google.golang.org/grpc"
)

// package main

// main
func main() {
	// 新建链接
	con, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer con.Close()

	// 新建特定服务的客户端
	// proto文件service后面的服务名 是该函数的一部分
	c := pb.NewHelloClient(con)

	// 调用服务
	rsp, err := c.Say(context.Background(), &pb.Req{Name: "client"})
	if err != nil {
		return
	}
	fmt.Println(rsp.Info)
}
