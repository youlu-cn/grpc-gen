package shared

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	fn := Func{
		Context: pgsgo.InitContext(params),
	}

	tpl.Funcs(map[string]interface{}{
		"pkg":    fn.PackageName,
		"scope":  fn.Scope,
		"hasGw":  fn.GatewayDefined,
		"access": fn.Access,
	})
}
