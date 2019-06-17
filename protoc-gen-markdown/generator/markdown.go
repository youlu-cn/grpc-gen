package generator

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/youlu-cn/grpc-gen/protoc-gen-markdown/templates"
)

const (
	MarkdownGenerator = "markdown"
)

type Auth struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func New() pgs.Module {
	return &Auth{
		ModuleBase: &pgs.ModuleBase{},
	}
}

func (m *Auth) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *Auth) Name() string {
	return MarkdownGenerator
}

func (m *Auth) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	var (
		outDir pgs.FilePath
	)

	// Process file-level templates
	tpls := templates.Template(m.Parameters())

	for _, f := range targets {
		m.Push(f.Name().String())

		// TODO: check

		for _, tpl := range tpls {
			out := templates.FilePathFor(tpl)(f, m.ctx, tpl)
			// A nil path means no output should be generated for this file - as controlled by
			// implementation-specific FilePathFor implementations.
			// Ex: Don't generate Java validators for files that don't reference PGV.
			if out != nil {
				outDir = out.Dir()
				m.AddGeneratorTemplateFile(out.String(), tpl, f)
			}
		}

		m.Pop()
	}

	// Table of Content
	tocTpl := templates.TOCTemplate(m.Parameters())
	tocOut := pgs.JoinPaths(outDir.String(), "README.md")

	m.Push("toc")

	m.AddGeneratorTemplateFile(tocOut.String(), tocTpl, targets)

	m.Pop()

	return m.Artifacts()
}
