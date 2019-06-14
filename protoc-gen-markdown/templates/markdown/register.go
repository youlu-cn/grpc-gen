package markdown

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	"github.com/youlu-cn/grpc-gen/protoc-gen-markdown/templates/shared"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	shared.Register(tpl, params)
	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("service").Parse(serviceTpl))
	template.Must(tpl.New("method").Parse(methodTpl))
}
