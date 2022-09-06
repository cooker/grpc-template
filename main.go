package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	c "grpc-template/core"
	bp "grpc-template/proto/generate"
	"net"
	"time"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type Server c.Server

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *bp.MessagePayload) (*bp.MessagePayload, error) {
	c.Infof("Received: %v", in)
	return &bp.MessagePayload{Header: &bp.MessageHeader{Code: "200"}}, nil
}

func (s *Server) SayHelloStream(in *bp.MessagePayload, out bp.Greeter_SayHelloStreamServer) error {
	c.Infof("Received Stream: %v", in)
	for i := 0; i < 10; i++ {
		if err := out.Send(&bp.MessagePayload{
			Header: &bp.MessageHeader{Code: "200", Message: "ccc"},
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

//grpc 服务端
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		c.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	bp.RegisterGreeterServer(s, &Server{})

	c.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		c.Fatalf("failed to serve: %v", err)
	}
}
