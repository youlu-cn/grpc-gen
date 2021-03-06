
# auth.proto 接口文档

*Document generated by protoc-gen-markdown. DO NOT EDIT.*

> 接口列表


* [Example](#example) - 服务说明，支持 markdown 语法


	* [Test (/example/test)](#test) - 方法说明，支持 markdown





<h2 id="example">Example</h2>

>  服务说明，支持 markdown 语法
> 
>  1. First item
>  2. Second item
>  3. Third item
> 
>  ```json
>  {
>    "firstName": "John",
>    "lastName": "Smith",
>    "age": 25
>  }
>  ```



<h3 id="test">Test</h3>

>  方法说明，支持 markdown
> 
>  - [x] Write the press release
>  - [ ] Update the website
>  - [ ] Contact the media



* HTTP Gateway

	* URL: `/example/test`
	* Method: `POST`
	* Content-Type: `application/json`

* 请求类型: ***ExampleMessage***

>  消息类型注释，支持多行，
>  支持 markdown 语法：
> 
>  > blockquote
> 
>  | Syntax | Description |
>  | ----------- | ----------- |
>  | Header | Title |
>  | Paragraph | Text |

|字段|protobuf 类型|json 类型|说明|默认值|是否必传|
|---|---|---|---|---|---|
|id|int64|string| 字段注释，简洁|-|false|




> JSON 示例

```json
{
  "id": "string($int64)"
}
```



* 返回类型: ***ExampleMessage***

>  消息类型注释，支持多行，
>  支持 markdown 语法：
> 
>  > blockquote
> 
>  | Syntax | Description |
>  | ----------- | ----------- |
>  | Header | Title |
>  | Paragraph | Text |

|字段|protobuf 类型|json 类型|说明|默认值|是否必传|
|---|---|---|---|---|---|
|id|int64|string| 字段注释，简洁|-|false|



> JSON 示例

```json
{
  "id": "string($int64)"
}
```






********

## *Embed Messages*














