package markdown

const methodTpl = `
<h3 id="{{ anchor .Name }}">{{ .Name.UpperCamelCase }}</h3>

> {{ .SourceCodeInfo.LeadingComments }}

* 请求类型: ***{{ .Input.Name.UpperCamelCase }}***

|字段|类型|说明|默认值|是否必传|
|---|---|---|---|---|
{{ range $v := (messageDoc .Input) }}|{{ $v.Name }}|{{ $v.Type }}|{{ $v.Comment }}|-|{{ $v.Required }}|
{{ end }}

* 返回类型: ***{{ .Output.Name.UpperCamelCase }}***

|字段|类型|说明|默认值|是否必传|
|---|---|---|---|---|
{{ range $v := (messageDoc .Output) }}|{{ $v.Name }}|{{ $v.Type }}|{{ $v.Comment }}|-|{{ $v.Required }}|
{{ end }}
`
