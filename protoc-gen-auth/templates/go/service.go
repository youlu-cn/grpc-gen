package golang

const serviceTpl = `
	var _level{{ .Name.UpperCamelCase }} = map[string]int32 {
		{{ range $k, $v := (auth .) }}"{{ $k }}": {{ $v }},
		{{ end }}
	}

	func AccessLevelOf{{ .Name.UpperCamelCase }}(fullPath string) int32 {
		return _level{{ .Name.UpperCamelCase }}[fullPath]
	}
`
