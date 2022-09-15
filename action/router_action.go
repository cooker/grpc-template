package action

import (
	"encoding/json"
	cmap "github.com/orcaman/concurrent-map/v2"
	c "grpc-template/core"
	bp "grpc-template/proto/generate"
)

//路由

var (
	MsgRouter = new(MessageRouter)
	msgQueue  = make(chan *bp.MessagePayload, 10000)
	connMap   = cmap.New[*ConnManager]()
)

type ConnManager struct {
	conn chan struct{}
	out  *bp.MessageService_PullServer
}

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

func init() {
	go startRoute()
}

func startRoute() {
	for {
		select {
		case msg := <-msgQueue:
			send(msg)
		}
	}
}

func send(msg *bp.MessagePayload) {
	var isSend bool
	for k, v := range connMap.Items() {
		if msg.Header != nil {
			if msg.Header.FromBy != k {
				err := (*(*v).out).Send(msg)
				if err != nil {
					close((*v).conn)
					connMap.Remove(k)
				} else {
					isSend = true
				}
			}
		} else {
			println("消息不合法")
		}
	}
	marshal, err := json.Marshal(*msg)
	if err == nil {
		c.Infof("消息发送状态：%t，%s\n", isSend, string(marshal))
	} else {
		panic(err)
	}
}
