package action

import (
	"context"
	c "grpc-template/core"
	bp "grpc-template/proto/generate"
)

type HeartBeatAction struct {
	bp.UnimplementedHeartBeatServiceServer
}

func (HeartBeatAction) Send(ctx context.Context, in *bp.HeartBeatRequest) (*bp.HeartBeatResponse, error) {
	c.Infof("心跳 %v\n", in)
	return &bp.HeartBeatResponse{
		Timestamp: c.NowTime(),
	}, nil
}
