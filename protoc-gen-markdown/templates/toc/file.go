package toc

const fileTpl = `
# 接口列表

{{ range $name, $file := . }}
## [{{ $file.Name }}]({{ $file.Name }}.md)
{{ range $svc := $file.Services }}
### [{{ $svc.Name.UpperCamelCase }}]({{ $file.Name }}.md#{{ anchor $svc.Name }}) - {{ tocComment $svc.SourceCodeInfo }}

> {{ leadingComment $svc.SourceCodeInfo }}
{{ range $method := $svc.Methods }}
{{ $url := (gatewayUrl $method) }}
* [{{ $method.Name.UpperCamelCase }}{{ if $url }} ({{ $url }}){{ end }}]({{ $file.Name }}.md#{{ anchor $method.Name }}) - {{ tocComment $method.SourceCodeInfo }}
{{ end }}
{{ end }}
{{ end }}
`
