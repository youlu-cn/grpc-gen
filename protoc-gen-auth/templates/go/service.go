package golang

const serviceTpl = `
	var _level{{ .Name.UpperCamelCase }} = map[string]auth.AccessLevel{
		{{ range $k, $v := (access .) }}"{{ $k }}": auth.AccessLevel_{{ $v }},
		{{ end }}
	}

	// Register scoped gRPC server.
	func Register{{ .Name.UpperCamelCase }}ScopeServer(a auth.Authenticator, s auth.Implementor, srv {{ .Name.UpperCamelCase }}Server) error {
		// Set service access level.
		a.SetAccessLevel(_level{{ .Name.UpperCamelCase }})

		// Register scoped gRPC server.
		for _, grpc := range s.ScopedGRPCServer(auth.VisibleScope_{{ (scope .) }}) {
			Register{{ .Name.UpperCamelCase }}Server(grpc, srv)
		}

		// Register scoped gateway handler.
		{{ if (hasGw .) }}return s.RegisterGateway(auth.VisibleScope_{{ (scope .) }}, Register{{ .Name.UpperCamelCase }}Handler)
		{{ else }}// No gateway generated.
		return nil
		{{ end }}
	}
`
