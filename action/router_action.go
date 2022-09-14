package action

import (
	bp "grpc-template/proto/generate"
)

//路由

var (
	MsgRouter = new(MessageRouter)
	msgQueue  = make(chan *bp.MessagePayload, 10000)
)

type IMessageRouter interface {
	Push(payload *bp.MessagePayload) error
}

type MessageRouter struct {
	IMessageRouter
}

func (r *MessageRouter) Push(payload *bp.MessagePayload) error {
	msgQueue <- payload
	return nil
}
