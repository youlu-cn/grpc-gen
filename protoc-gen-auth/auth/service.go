package auth

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GatewayHandler func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error

type Implementor interface {
	ScopedGRPCServer(scope VisibleScope) []*grpc.Server
	RegisterGateway(scope VisibleScope, handlers GatewayHandler) error
}

// Server authenticator interface.
type Authenticator interface {
	SetAccessLevel(map[string]AccessLevel)
	Authenticate(ctx context.Context, fullMethod string) error
}

// UnaryServerInterceptor returns a new unary server interceptor that authenticates incoming messages.
//
// Invalid messages will be rejected with `PermissionDenied` before reaching any userspace handlers.
func UnaryServerInterceptor(v Authenticator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := v.Authenticate(ctx, info.FullMethod); err != nil {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that authenticates incoming messages.
//
// The stage at which unauthenticated messages will be rejected with `PermissionDenied` varies based on the
// type of the RPC. For `ServerStream` (1:m) requests, it will happen before reaching any userspace
// handlers. For `ClientStream` (n:1) or `BidiStream` (n:m) RPCs, the messages will be rejected on
// calls to `stream.Recv()`.
func StreamServerInterceptor(v Authenticator) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if err := v.Authenticate(stream.Context(), info.FullMethod); err != nil {
			return status.Errorf(codes.Unauthenticated, err.Error())
		}
		return handler(srv, stream)
	}
}
