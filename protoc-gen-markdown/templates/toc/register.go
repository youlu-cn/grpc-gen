package toc

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	"github.com/youlu-cn/grpc-gen/protoc-gen-markdown/templates/shared"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	shared.Register(tpl, params)
	template.Must(tpl.Parse(fileTpl))
}
