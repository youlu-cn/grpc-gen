package grpc

const serviceTpl = `
<h2 id="{{ .FullyQualifiedName }}"> {{ .Name.UpperCamelCase }} </h2>

> {{ .SourceCodeInfo.LeadingComments }}

{{ range .Methods }}
{{ template "method" . }}
{{ end }}
`
