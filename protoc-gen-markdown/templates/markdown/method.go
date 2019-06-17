package markdown

const methodTpl = `
<h3 id="{{ anchor .Name }}">{{ .Name.UpperCamelCase }}</h3>

> {{ leadingComment .SourceCodeInfo }}

{{ $gateway := (gatewayDoc .) }}
{{ if $gateway }}
* HTTP Gateway

	* URL: {{ $gateway.URL }}
	* Method: {{ $gateway.Method }}
	* Content-Type: {{ $gateway.ContentType }}
{{ end }}
* 请求类型: ***{{ .Input.Name.UpperCamelCase }}***

> {{ leadingComment .Input.SourceCodeInfo }}

|字段|protobuf 类型|json 类型|说明|默认值|是否必传|
|---|---|---|---|---|---|
{{ range $v := (messageDoc .Input) }}|{{ $v.Name }}|{{ $v.ProtoType }}|{{ $v.JsonType }}|{{ $v.Comment }}|{{ $v.Default }}|{{ $v.Required }}|
{{ end }}

{{ if $gateway }}
{{ if $gateway.JsonRequired }}
> JSON 示例

{{ (jsonDemo .Input) }}
{{ end }}
{{ end }}

* 返回类型: ***{{ .Output.Name.UpperCamelCase }}***

> {{ leadingComment .Output.SourceCodeInfo }}

|字段|protobuf 类型|json 类型|说明|默认值|是否必传|
|---|---|---|---|---|---|
{{ range $v := (messageDoc .Output) }}|{{ $v.Name }}|{{ $v.ProtoType }}|{{ $v.JsonType }}|{{ $v.Comment }}|{{ $v.Default }}|{{ $v.Required }}|
{{ end }}

{{ if $gateway }}
> JSON 示例

{{ (jsonDemo .Output) }}
{{ end }}
`
