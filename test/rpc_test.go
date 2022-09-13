package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	c "grpc-template/core"
	bp "grpc-template/proto/generate"
	"testing"
)

func TestHeartBeat(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := bp.NewHeartBeatServiceClient(conn)
	resp, err := client.Send(context.Background(), &bp.HeartBeatRequest{
		Timestamp: c.NowTime(),
		MsgId:     c.CreateMsgId("1"),
		FromBy:    "1",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("发送成功 %v\n", resp)

	conn.Close()
}
