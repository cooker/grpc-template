package core

import bp "grpc-template/proto/generate"

type Server struct {
	bp.UnimplementedGreeterServer
}
