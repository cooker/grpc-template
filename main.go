package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc-template/action"
	c "grpc-template/core"
	bp "grpc-template/proto/generate"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

//grpc 服务端
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		c.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	bp.RegisterHeartBeatServiceServer(s, &action.HeartBeatAction{})
	//bp.RegisterGreeterServer(s, &Server{})

	c.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		c.Fatalf("failed to serve: %v", err)
	}
}
