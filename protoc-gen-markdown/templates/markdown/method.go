package markdown

const methodTpl = `
<h3 id="{{ anchor .Name }}">{{ .Name.UpperCamelCase }}</h3>

> {{ .SourceCodeInfo.LeadingComments }}

* 请求类型: ***{{ .Input.Name.UpperCamelCase }}***

|字段|protobuf 类型|json 类型|说明|默认值|是否必传|
|---|---|---|---|---|---|
{{ range $v := (messageDoc .Input) }}|{{ $v.Name }}|{{ $v.ProtoType }}|{{ $v.JsonType }}|{{ $v.Comment }}|-|{{ $v.Required }}|
{{ end }}

* 返回类型: ***{{ .Output.Name.UpperCamelCase }}***

|字段|protobuf 类型|json 类型|说明|默认值|是否必传|
|---|---|---|---|---|---|
{{ range $v := (messageDoc .Output) }}|{{ $v.Name }}|{{ $v.ProtoType }}|{{ $v.JsonType }}|{{ $v.Comment }}|-|{{ $v.Required }}|
{{ end }}
`
