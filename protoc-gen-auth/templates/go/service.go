package golang

const serviceTpl = `
	var _level{{ .Name.UpperCamelCase }} = map[string]auth.AccessLevel{
		{{ range $k, $v := (access .) }}"{{ $k }}": auth.AccessLevel_{{ $v }},
		{{ end }}
	}

	func AccessLevelOf{{ .Name.UpperCamelCase }}(fullPath string) auth.AccessLevel {
		return _level{{ .Name.UpperCamelCase }}[fullPath]
	}

	// Register scoped gRPC server.
	//{{ .SourceCodeInfo.LeadingComments }}
	func Register{{ .Name.UpperCamelCase }}ScopeServer(s auth.Service, srv interface{}) {
		for _, grpc := range s.ScopedGRPCServer(auth.VisibleScope_{{ (scope .) }}) {
			grpc.RegisterService(&_{{ .Name.UpperCamelCase }}_serviceDesc, srv)
		}
	}

	// Register scoped gateway handler.
	//{{ .SourceCodeInfo.LeadingComments }}
	func Register{{ .Name.UpperCamelCase }}ScopeHandler(s auth.Service) error {
		{{ if (hasGw .) }}return s.RegisterGateway(auth.VisibleScope_{{ (scope .) }}, Register{{ .Name.UpperCamelCase }}Handler)
		{{ else }}// No gateway generated.
		return nil
		{{ end }}
	}
`
