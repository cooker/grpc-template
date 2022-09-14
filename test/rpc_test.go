package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	c "grpc-template/core"
	bp "grpc-template/proto/generate"
	"testing"
	"time"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func TestHeartBeat(t *testing.T) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacp),
	}
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		panic(err)
	}

	timer := time.NewTicker(1 * time.Second)

	go func() {
		for {

			_, ok := <-timer.C
			if !ok {
				println("退出")
				break
			}
			client := bp.NewHeartBeatServiceClient(conn)
			outgoingContext := metadata.AppendToOutgoingContext(context.TODO(), "authorization", "123456")
			resp, err := client.Send(outgoingContext, &bp.HeartBeatRequest{
				Timestamp: c.NowTime(),
				MsgId:     c.CreateMsgId("1"),
				FromBy:    "1",
			})

			if err != nil {
				fmt.Println("出现异常", err)
				time.Sleep(3 * time.Second)
			}
			fmt.Printf("发送成功 %v\n", resp)

		}
	}()

	select {}
	defer conn.Close()
	defer timer.Stop()
}

func BenchmarkHeartBeat(b *testing.B) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		client := bp.NewHeartBeatServiceClient(conn)
		_, err := client.Send(context.Background(), &bp.HeartBeatRequest{
			Timestamp: c.NowTime(),
			MsgId:     c.CreateMsgId("1"),
			FromBy:    "1",
		})
		if err != nil {
			panic(err)
		}
	}
	conn.Close()
}
