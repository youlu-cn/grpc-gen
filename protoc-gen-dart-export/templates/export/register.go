package export

import (
	pgs "github.com/lyft/protoc-gen-star"
	"text/template"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	template.Must(tpl.Parse(fileTpl))
}
