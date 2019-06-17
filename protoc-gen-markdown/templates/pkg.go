package templates

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/youlu-cn/grpc-gen/protoc-gen-markdown/templates/markdown"
	"github.com/youlu-cn/grpc-gen/protoc-gen-markdown/templates/toc"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)
type FilePathFn func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath

func Template(params pgs.Parameters) []*template.Template {
	return []*template.Template{
		makeTemplate("markdown", markdown.Register, params),
	}
}

func TOCTemplate(params pgs.Parameters) *template.Template {
	return makeTemplate("toc", toc.Register, params)
}

func FilePathFor(tpl *template.Template) FilePathFn {
	switch tpl.Name() {
	default:
		return func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
			out := ctx.OutputPath(f)
			out = pgs.JoinPaths(out.Dir().String(), f.Name().String()+".md")
			return &out
		}
	}
}

func makeTemplate(typ string, fn RegisterFn, params pgs.Parameters) *template.Template {
	tpl := template.New(typ)
	fn(tpl, params)
	return tpl
}
