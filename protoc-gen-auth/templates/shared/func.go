package shared

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/youlu-cn/grpc-gen/protoc-gen-auth/auth"
	"google.golang.org/genproto/googleapis/api/annotations"
)

type Func struct {
	pgsgo.Context
}

func (fn Func) Access(svc pgs.Service) map[string]auth.AccessLevel {
	out := make(map[string]auth.AccessLevel)

	for _, method := range svc.Methods() {
		fullPath := fmt.Sprintf("/%s.%s/%s", svc.Package().ProtoName(), svc.Name(), method.Name().UpperCamelCase())
		out[fullPath] = auth.AccessLevel__NO_LIMIT

		opts := method.Descriptor().GetOptions()
		descs, _ := proto.ExtensionDescs(opts)

		for _, desc := range descs {
			if desc.Field == 2360 {
				ext, _ := proto.GetExtension(opts, desc)
				if access, ok := ext.(*auth.Access); ok {
					out[fullPath] = access.Level
					break
				}
			}
		}
	}

	return out
}

func (fn Func) Scope(svc pgs.Service) auth.VisibleScope {
	opts := svc.Descriptor().GetOptions()
	descs, _ := proto.ExtensionDescs(opts)

	for _, desc := range descs {
		if desc.Field == 1360 {
			ext, _ := proto.GetExtension(opts, desc)
			if visible, ok := ext.(*auth.Visible); ok {
				return visible.Scope
			}
		}
	}

	return auth.VisibleScope_PUBLIC_SCOPE
}

func (fn Func) GatewayDefined(svc pgs.Service) bool {
	for _, method := range svc.Methods() {
		opts := method.Descriptor().GetOptions()
		descs, _ := proto.ExtensionDescs(opts)

		for _, desc := range descs {
			if desc.Field == 72295728 {
				ext, _ := proto.GetExtension(opts, desc)
				if _, ok := ext.(*annotations.HttpRule); ok {
					return true
				}
			}
		}
	}

	return false
}
