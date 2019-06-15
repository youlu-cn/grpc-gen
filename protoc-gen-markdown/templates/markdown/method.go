package markdown

const methodTpl = `
<h3 id="{{ anchor .Name }}">{{ .Name.UpperCamelCase }}</h3>

> {{ .SourceCodeInfo.LeadingComments }}

{{ $gateway := (gatewayDoc .) }}
{{ if $gateway }}
* HTTP Gateway

	* URL: {{ $gateway.URL }}
	* Method: {{ $gateway.Method }}
	* Content-Type: {{ $gateway.ContentType }}
{{ end }}
* 请求参数

> gRPC object 类型: ***{{ .Input.Name.UpperCamelCase }}***

|字段|protobuf 类型|json 类型|说明|默认值|是否必传|
|---|---|---|---|---|---|
{{ range $v := (messageDoc .Input) }}|{{ $v.Name }}|{{ $v.ProtoType }}|{{ $v.JsonType }}|{{ $v.Comment }}|-|{{ $v.Required }}|
{{ end }}

> JSON 示例

{{ (jsonDemo .Input) }}

* 返回值

> gRPC object 类型: ***{{ .Output.Name.UpperCamelCase }}***

|字段|protobuf 类型|json 类型|说明|默认值|是否必传|
|---|---|---|---|---|---|
{{ range $v := (messageDoc .Output) }}|{{ $v.Name }}|{{ $v.ProtoType }}|{{ $v.JsonType }}|{{ $v.Comment }}|-|{{ $v.Required }}|
{{ end }}

> JSON 示例

{{ (jsonDemo .Output) }}
`
