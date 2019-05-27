package auth

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type GatewayHandler func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error

type Service interface {
	ScopedGRPCServer(scope VisibleScope) []*grpc.Server
	RegisterGateway(scope VisibleScope, handlers ...GatewayHandler) error
}
