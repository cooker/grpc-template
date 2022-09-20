package action

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	cmap "github.com/orcaman/concurrent-map/v2"
	"google.golang.org/protobuf/proto"
	c "grpc-template/core"
	bp "grpc-template/proto/generate"
)

//路由

var (
	MsgRouter = new(MessageRouter)
	msgQueue  = make(chan *bp.MessagePayload, 10000)
	connMap   = cmap.New[*ConnManager]()
	nc        *nats.Conn
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

func startRoute() {
	c.ConfigYml.Wait()
	c.Infof("启动消息路由")
	initNats()
	for {
		select {
		case msg := <-msgQueue:
			mess, err := proto.Marshal(msg)
			if err != nil {
				c.Errorf("消息路由失败", c.Json2Str(msg), err)
			} else {
				err := nc.Publish(c.ConfigYml.Channel.Topic, mess)
				if err != nil {
					c.Errorf("消息路由失败 %s", c.Json2Str(msg), err)
				}
			}
		}
	}
}

func initNats() {
	nc, _ = nats.Connect(c.ConfigYml.Channel.Url)
	nc.Subscribe(c.ConfigYml.Channel.Topic, func(msg *nats.Msg) {
		payload := bp.MessagePayload{}
		err := proto.Unmarshal(msg.Data, &payload)
		if err != nil {
			c.Errorf("消息接收失败 >> %s", msg.Header)
			return
		}
		send(&payload)
	})
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
			c.Warn("消息不合法")
		}
	}
	marshal, err := json.Marshal(*msg)
	if err == nil {
		c.Infof("消息发送状态：%t，%s\n", isSend, string(marshal))
	} else {
		panic(err)
	}
}

//初始化
func init() {
	go startRoute()
}
