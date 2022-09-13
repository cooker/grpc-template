package core

import (
	"github.com/bwmarrin/snowflake"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NowTime() *timestamppb.Timestamp {
	return timestamppb.Now()
}

func CreateMsgId(clientId string) string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "0" + "-" + clientId
	}
	id := node.Generate()
	return id.String() + "-" + clientId
}
