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
		"pkg":           fn.PackageName,
		"anchor":        fn.Anchor,
		"toc":           fn.TableOfContent,
		"gatewayDoc":    fn.GatewayDoc,
		"messageDoc":    fn.MessageDoc,
		"embedMessages": fn.EmbedMessages,
		"jsonDemo":      fn.JSONDemo,
	})
}
