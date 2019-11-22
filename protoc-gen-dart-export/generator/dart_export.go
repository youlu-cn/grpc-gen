package generator

import (
	"path"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/youlu-cn/grpc-gen/protoc-gen-dart-export/templates"
)

const (
	DartExportGenerator = "dart-export"
	FileParam           = "file"
)

type DartExport struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func New() pgs.Module {
	return &DartExport{
		ModuleBase: &pgs.ModuleBase{},
	}
}

func (m *DartExport) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *DartExport) Name() string {
	return DartExportGenerator
}

func (m *DartExport) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	var (
		outDir  pgs.FilePath
		outFile = m.Parameters().Str(FileParam)
	)

	if outFile == "" {
		outFile = "export.dart"
	} else if path.Ext(outFile) == "" {
		outFile += ".dart"
	}

	tmpl := templates.Template(m.Parameters())
	for _, f := range targets {
		out := templates.FilePathFor(tmpl)(f, m.ctx, tmpl)
		if out != nil {
			outDir = out.Dir()
			break
		}
	}

	// Dart export
	dartOut := pgs.JoinPaths(outDir.String(), outFile)

	m.Push("export")

	m.AddGeneratorTemplateFile(dartOut.String(), tmpl, targets)

	m.Pop()

	return m.Artifacts()
}
