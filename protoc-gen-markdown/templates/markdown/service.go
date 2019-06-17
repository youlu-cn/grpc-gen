package markdown

const serviceTpl = `
<h2 id="{{ anchor .Name }}">{{ .Name.UpperCamelCase }}</h2>

> {{ leadingComment .SourceCodeInfo }}

{{ range .Methods }}
{{ template "method" . }}
{{ end }}
`
