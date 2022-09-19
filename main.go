package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"grpc-template/action"
	c "grpc-template/core"
	f "grpc-template/filter"
	bp "grpc-template/proto/generate"
	"math"
	"net"
	"time"
)

var (
	port  = flag.Int("port", 50051, "The server port")
	debug = flag.Bool("debug", true, "server debug")
	kaep  = keepalive.EnforcementPolicy{
		MinTime:             30 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,             // Allow pings even when there are no active streams
	}
	kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second,             // 如果客戶端空閒 15 秒，則發送 GOAWAY
		MaxConnectionAge:      time.Duration(math.MaxInt64), // 在強制關閉連接之前等待 5 秒等待 RPC 完成
		MaxConnectionAgeGrace: 5 * time.Second,              // 在強制關閉連接之前等待 5 秒等待 RPC 完成
		Time:                  10 * time.Second,             // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               3 * time.Second,              // Wait 1 second for the ping ack before assuming the connection is dead
	}
)

//grpc 服务端
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		c.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(serverOption()...)
	registerAction(s)
	if *debug {
		reflection.Register(s)
	}
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

func serverOption() []grpc.ServerOption {
	return []grpc.ServerOption{
		//心跳
		grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp),
		//拦截器
		grpc.UnaryInterceptor(f.ValidAuthFilter),
		grpc.StreamInterceptor(f.ValidAuthStreamFilter),
	}
}
