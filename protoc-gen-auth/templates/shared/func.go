package shared

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/youlu-cn/grpc-gen/protoc-gen-auth/options"
)

type Func struct {
	pgsgo.Context
}

func (fn Func) Auth(svc pgs.Service) map[string]int32 {
	out := make(map[string]int32)

	for _, method := range svc.Methods() {
		fullPath := fmt.Sprintf("/%s.%s/%s", svc.Package().ProtoName(), svc.Name(), method.Name())

		opts := method.Descriptor().GetOptions()
		descs, _ := proto.ExtensionDescs(opts)

		for _, desc := range descs {
			if desc.Field == 1042 {
				ext, _ := proto.GetExtension(opts, desc)
				if access, ok := ext.(*options.Access); ok {
					out[fullPath] = access.Level
					break
				}
			}
		}
	}

	return out
}
