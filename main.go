package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
	"grpc-template/action"
	c "grpc-template/core"
	f "grpc-template/filter"
	bp "grpc-template/proto/generate"
	"io/ioutil"
	"math"
	"net"
	"time"
)

var (
	port   = flag.Int("port", 50051, "server port")
	config = flag.String("config", "config.yml", "server config")
	debug  = flag.Bool("debug", true, "server debug")
	kaep   = keepalive.EnforcementPolicy{
		MinTime:             30 * time.Second, // 如果客戶端每 5 秒 ping 一次以上，則終止連接
		PermitWithoutStream: true,             // 即使沒有活動流也允許 ping
	}
	kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second,             // 如果客戶端空閒 15 秒，則發送 GOAWAY
		MaxConnectionAge:      time.Duration(math.MaxInt64), // 在強制關閉連接之前等待 5 秒等待 RPC 完成
		MaxConnectionAgeGrace: 5 * time.Second,              // 在強制關閉連接之前等待 5 秒等待 RPC 完成
		Time:                  10 * time.Second,             // 如果客戶端空閒 5 秒，則 Ping 客戶端以確保連接仍然處於活動狀態
		Timeout:               3 * time.Second,              // 在假設連接已死之前等待 1 秒等待 ping 確認
	}
)

//grpc 服务端
func main() {
	flag.Parse()
	confData, err := ioutil.ReadFile(*config)
	if err != nil {
		c.Errorf("配置文件读取失败", err)
		return
	}
	err = yaml.Unmarshal(confData, &c.ConfigYml)
	if err != nil {
		c.Errorf("配置文件读取失败", err)
		return
	}
	c.Infof("配置文件：%s", c.Json2Str(c.ConfigYml))
	c.ConfigYml.Done()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.ConfigYml.Port))
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
