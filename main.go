package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"grpc-template/action"
	c "grpc-template/core"
	f "grpc-template/filter"
	bp "grpc-template/proto/generate"
	"net"
	"time"
)

var (
	port = flag.Int("port", 50051, "The server port")
	kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}
	kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	}
)

//grpc 服务端
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		c.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp),
	}

	opts = append(opts, registerFilter()...)

	s := grpc.NewServer(opts...)
	registerAction(s)

	c.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		c.Fatalf("failed to serve: %v", err)
	}
}

//注册action
func registerAction(s *grpc.Server) {
	bp.RegisterHeartBeatServiceServer(s, &action.HeartBeatAction{})
	bp.RegisterMessageServiceServer(s, &action.MessageAction{})
}

//拦截器
func registerFilter() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(f.ValidAuthFilter),
	}
}
