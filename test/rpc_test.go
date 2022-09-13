package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	c "grpc-template/core"
	bp "grpc-template/proto/generate"
	"testing"
	"time"
)

func TestHeartBeat(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
			resp, err := client.Send(context.TODO(), &bp.HeartBeatRequest{
				Timestamp: c.NowTime(),
				MsgId:     c.CreateMsgId("1"),
				FromBy:    "1",
			})

			if err != nil {
				panic(err)
			}
			fmt.Printf("发送成功 %v\n", resp)

		}
	}()

	time.Sleep(30 * time.Second)
	timer.Stop()
	conn.Close()
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
