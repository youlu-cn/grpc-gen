// Code generated by protoc-gen-auth. DO NOT EDIT.
// source: auth.proto
package example

import "github.com/youlu-cn/grpc-gen/protoc-gen-auth/auth"

// Reference imports to suppress errors if they are not otherwise used.
var _ = auth.AccessLevel__NO_LIMIT

var _levelExample = map[string]auth.AccessLevel{
	"/example.Example/test1": auth.AccessLevel_LOW_ACCESS_LEVEL,
}

func AccessLevelOfExample(fullPath string) auth.AccessLevel {
	return _levelExample[fullPath]
}

func RegisterExampleScopeServer(s auth.Service, srv interface{}) {
	for _, grpc := range s.ScopedGRPCServer(auth.VisibleScope_PUBLIC_SCOPE) {
		grpc.RegisterService(&_Example_serviceDesc, srv)
	}
}

func RegisterExampleScopeHandler(s auth.Service) error {
	// No gateway generated.
	return nil

}

var _levelExample2 = map[string]auth.AccessLevel{
	"/example.Example2/test2": auth.AccessLevel_MIDDLE_ACCESS_LEVEL,
	"/example.Example2/test3": auth.AccessLevel_SERVER_INTERNAL,
}

func AccessLevelOfExample2(fullPath string) auth.AccessLevel {
	return _levelExample2[fullPath]
}

func RegisterExample2ScopeServer(s auth.Service, srv interface{}) {
	for _, grpc := range s.ScopedGRPCServer(auth.VisibleScope_ALL_SCOPE) {
		grpc.RegisterService(&_Example2_serviceDesc, srv)
	}
}

func RegisterExample2ScopeHandler(s auth.Service) error {
	return s.RegisterGateway(auth.VisibleScope_ALL_SCOPE, RegisterExample2Handler)

}
