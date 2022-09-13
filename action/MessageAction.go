package action

import bp "grpc-template/proto/generate"

type MessageAction struct {
	bp.UnimplementedMessageServiceServer
}
