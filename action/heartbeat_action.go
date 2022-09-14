package action

import (
	"context"
	c "grpc-template/core"
	bp "grpc-template/proto/generate"
	"strconv"
)

type HeartBeatAction struct {
	bp.UnimplementedHeartBeatServiceServer
}

func (HeartBeatAction) Send(ctx context.Context, in *bp.HeartBeatRequest) (*bp.MessageResponse, error) {
	c.Infof("心跳 %v", in)
	return &bp.MessageResponse{
		Timestamp: c.NowTime(),
		MsgId:     c.CreateMsgId(strconv.Itoa(c.GATEWAY)),
	}, nil
}
