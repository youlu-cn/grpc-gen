package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/youlu-cn/grpc-gen/protoc-gen-auth/generator"
)

func main() {
	pgs.
		Init(pgs.DebugEnv("DEBUG_GRPC_GEN")).
		RegisterModule(generator.New()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
