package markdown

const serviceTpl = `
<h2 id="{{ anchor .Name }}">{{ .Name.UpperCamelCase }}</h2>

> {{ .SourceCodeInfo.LeadingComments }}

{{ range .Methods }}
{{ template "method" . }}
{{ end }}
`
