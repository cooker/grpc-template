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
	Time:                5 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             3 * time.Second, // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,            // send pings even without active streams
}

func connection() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacp),
	}
	conn, err := grpc.Dial("localhost:50051", opts...)
	return conn, err
}

func TestHeartBeat(t *testing.T) {

	conn, err := connection()
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

	for {

	}
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

func TestSendMessage(t *testing.T) {
	conn, err := connection()
	if err != nil {
		panic(err)
	}

	client := bp.NewMessageServiceClient(conn)
	outgoingContext := metadata.AppendToOutgoingContext(context.TODO(), "authorization", "123456")
	client.Send(outgoingContext, &bp.MessagePayload{
		Header: &bp.MessageHeader{
			Timestamp: c.NowTime(),
			MsgId:     c.CreateMsgId("12"),
			FromBy:    "12",
			SendTo:    "33",
			RouteType: bp.RouteType_RULE,
			State: &bp.State{
				Code:    "200",
				Message: "success",
			},
		},
		Property: nil,
		Body: &bp.MessageBody{
			Content: "hello",
		},
	})
}

func TestPullMessage(t *testing.T) {
	conn, err := connection()
	if err != nil {
		panic(err)
	}

	client := bp.NewMessageServiceClient(conn)
	outgoingContext := metadata.AppendToOutgoingContext(context.TODO(), "authorization", "123456")
	request := bp.ClientPullRequest{
		FromBy: "222",
	}
	pull, err := client.Pull(outgoingContext, &request)
	if err == nil {
		for {
			recv, err := pull.Recv()
			if err == nil && recv != nil {
				t.Logf("收到消息 %v\n", recv)
			} else {
				t.Log(err)
				break
			}
		}
	} else {
		panic(err)
	}

}
