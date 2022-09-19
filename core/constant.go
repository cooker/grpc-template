package core

import bp "grpc-template/proto/generate"

const (
	GATEWAY = iota
	ROUTE
)

const AUTH_HEADER = "authorization"

var SUCCESS_STATE = bp.State{
	Code:    "200",
	Message: "success",
}
