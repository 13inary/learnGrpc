package model

import (
	"fmt"
	"learnGrpc/proto/stream"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Get(req *stream.Req, res stream.Stream_GetServer) (err error) {
	for i := 1; i <= 10; i++ {
		err = res.Send(&stream.Rsp{
			Info: fmt.Sprintf("%v i am server", time.Now()),
		})
		if err != nil {
			return err
		}

		time.Sleep(time.Second)
	}
	return nil
}

func (s *server) Put(reqs stream.Stream_PutServer) error {
	for {
		data, err := reqs.Recv()
		if err != nil {
			log.Printf("err = %+v\n", err)
			return err
		}

		log.Printf("data = %+v\n", data)
	}
}

func (s *server) Chat(alls stream.Stream_ChatServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			data, err := alls.Recv()
			if err != nil {
				log.Printf("err = %+v\n", err)
				break
			}

			log.Printf("data = %+v\n", data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			err := alls.Send(&stream.Rsp{
				Info: "i am server",
			})
			if err != nil {
				log.Printf("err = %+v\n", err)
				break
			}

			time.Sleep(time.Second)
		}
	}()

	wg.Wait()

	return nil
}

// StreamModel 双向数据流模式 client and server can send more data
func StreamModel() {
	log.Println("listen port")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Printf("err = %+v\n", err)
	}

	log.Println("new server and register method")
	g := grpc.NewServer()
	stream.RegisterStreamServer(g, &server{})

	log.Println("run server")
	err = g.Serve(lis)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
}
