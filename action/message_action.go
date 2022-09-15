package action

import (
	"context"
	"fmt"
	c "grpc-template/core"
	bp "grpc-template/proto/generate"
	"strconv"
)

type MessageAction struct {
	bp.UnimplementedMessageServiceServer
}

func (s *MessageAction) Send(ctx context.Context, playload *bp.MessagePayload) (*bp.MessageResponse, error) {
	c.Infof("接收消息：%v", *playload)
	MsgRouter.Push(playload)
	return &bp.MessageResponse{
		Timestamp: c.NowTime(),
		MsgId:     c.CreateMsgId(strconv.Itoa(c.GATEWAY)),
		State:     nil,
	}, nil
}

func (s *MessageAction) Pull(request *bp.ClientPullRequest, out bp.MessageService_PullServer) error {
	conn := make(chan struct{})
	manager := ConnManager{
		conn: conn,
		out:  &out,
	}
	connMap.Set(request.MsgId, &manager)

	select {
	case <-conn:
		fmt.Printf("链接关闭 %s", request.MsgId)
	}
	return nil
}
