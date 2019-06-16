package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"google.golang.org/genproto/googleapis/api/annotations"
)

const (
	Enum    = "enum"
	Message = "message"
)

type Func struct {
	pgsgo.Context
}

type EnumDocValue struct {
	Name    string
	Value   int32
	Comment string
}

type EnumDoc struct {
	pgs.Enum
	Values []*EnumDocValue
}

type MessageDocField struct {
	Name      string
	Comment   string
	ProtoType string
	JsonType  string
	Default   string
	Required  bool
}

type MessageDoc struct {
	pgs.Message
	Fields []*MessageDocField
}

func (fn Func) Anchor(name pgs.Name) string {
	return name.Transform(strings.ToLower, strings.ToLower, "").String()
}

func (fn Func) comment(info pgs.SourceCodeInfo) string {
	comment := "TODO"
	if info.TrailingComments() != "" {
		comment = info.TrailingComments()
	} else if info.LeadingComments() != "" {
		comment = info.LeadingComments()
	}
	comment = strings.Trim(comment, "\n")
	return strings.Replace(comment, "\n", "<br>", -1)
}

func (fn Func) fieldElementType(el pgs.FieldTypeElem) string {
	if el.IsEmbed() {
		msg := el.Embed()
		return fmt.Sprintf("[%s](#%s)", msg.Name(), fn.Anchor(msg.Name()))
	} else if el.IsEnum() {
		enum := el.Enum()
		if enum.FullyQualifiedName() == ".google.protobuf.NullValue" {
			return "null"
		}
		return fmt.Sprintf("[%s](#%s)", enum.Name(), fn.Anchor(enum.Name()))
	}

	switch el.ProtoType() {
	case pgs.DoubleT, pgs.FloatT:
		return "float"
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return "int64"
	case pgs.Int32T, pgs.UInt32T, pgs.Fixed32T, pgs.SInt32, pgs.SFixed32:
		return "int"
	case pgs.BoolT:
		return "bool"
	case pgs.StringT:
		return "string"
	case pgs.BytesT:
		return "bytes"
	default:
		return "UNKNOWN"
	}
}

func (fn Func) internalEmbedJsonType(fullyQualifiedName string) string {
	switch fullyQualifiedName {
	case ".google.protobuf.Timestamp":
		return `string ("1972-01-01T10:00:20.021Z")`
	case ".google.protobuf.Duration":
		return `string ("1.000340012s")`
	case ".google.protobuf.Empty":
		return `object "{}"`
	case ".google.protobuf.FieldMask":
		return "string"
	case ".google.protobuf.ListValue":
		return "array"
	case ".google.protobuf.DoubleValue", ".google.protobuf.FloatValue":
		return "number/string"
	case ".google.protobuf.Int64Value", ".google.protobuf.UInt64Value":
		return "string"
	case ".google.protobuf.Int32Value", ".google.protobuf.UInt32Value":
		return "number/string"
	case ".google.protobuf.BoolValue":
		return "true, false"
	case ".google.protobuf.StringValue":
		return "string"
	case ".google.protobuf.BytesValue":
		return "base64 string"
	default:
		return "object"
	}
}

func (fn Func) fieldType(field pgs.Field) (pbType string, jsonType string) {
	switch field.Type().ProtoType() {
	case pgs.DoubleT, pgs.FloatT:
		pbType = "float"
		jsonType = "number/string"
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		pbType = "int64"
		jsonType = "string"
	case pgs.Int32T, pgs.UInt32T, pgs.Fixed32T, pgs.SInt32, pgs.SFixed32:
		pbType = "int"
		jsonType = "number/string"
	case pgs.BoolT:
		pbType = "bool"
		jsonType = "true, false"
	case pgs.StringT:
		pbType = "string"
		jsonType = "string"
	case pgs.BytesT:
		pbType = "bytes"
		jsonType = "base64 string"
	case pgs.EnumT:
		enum := field.Type().Enum()
		pbType = fmt.Sprintf("enum [%s](#%s)", enum.Name(), fn.Anchor(enum.Name()))
		if enum.FullyQualifiedName() == ".google.protobuf.NullValue" {
			jsonType = "null"
		} else {
			jsonType = "string/integer"
		}
	case pgs.MessageT:
		if field.Type().IsMap() {
			key := fn.fieldElementType(field.Type().Key())
			value := fn.fieldElementType(field.Type().Element())
			pbType = fmt.Sprintf("map\\<%s, %s\\>", key, value)
			jsonType = "object"
		} else if field.Type().IsRepeated() {
			el := fn.fieldElementType(field.Type().Element())
			pbType = el
			jsonType = "array"
		} else {
			msg := field.Type().Embed()
			pbType = fmt.Sprintf("[%s](#%s)", msg.Name(), fn.Anchor(msg.Name()))
			jsonType = fn.internalEmbedJsonType(msg.FullyQualifiedName())
		}
	case pgs.GroupT:
		// TODO
	}
	return
}

func (fn Func) EnumDoc(enum pgs.Enum) []*EnumDocValue {
	doc := make([]*EnumDocValue, 0, len(enum.Values()))

	for _, v := range enum.Values() {
		doc = append(doc, &EnumDocValue{
			Name:    v.Name().String(),
			Value:   v.Value(),
			Comment: fn.comment(v.SourceCodeInfo()),
		})
	}

	return doc
}

func (fn Func) MessageDoc(msg pgs.Message) []*MessageDocField {
	out := make([]*MessageDocField, 0, len(msg.Fields()))

	for _, field := range msg.Fields() {
		doc := &MessageDocField{
			Name:     field.Name().String(),
			Required: field.Required(),
		}
		// Type
		doc.ProtoType, doc.JsonType = fn.fieldType(field)
		if field.Type().IsRepeated() {
			doc.ProtoType = fmt.Sprintf("array [%s]", doc.ProtoType)
		}
		// Comment
		doc.Comment = fn.comment(field.SourceCodeInfo())
		out = append(out, doc)
	}

	return out
}

func (fn Func) EmbedFields(field pgs.Field, enumDoc map[string]interface{}, msgDoc map[string]interface{}) {
	switch field.Type().ProtoType() {
	case pgs.EnumT:
		name := field.Type().Enum().FullyQualifiedName()
		enumDoc[name] = &EnumDoc{
			Enum:   field.Type().Enum(),
			Values: fn.EnumDoc(field.Type().Enum()),
		}
	case pgs.MessageT:
		if field.Type().IsMap() {
			key := field.Type().Key()
			value := field.Type().Element()
			if key.IsEnum() {
				name := key.Enum().FullyQualifiedName()
				enumDoc[name] = &EnumDoc{
					Enum:   key.Enum(),
					Values: fn.EnumDoc(key.Enum()),
				}
			} else if key.IsEmbed() {
				name := key.Embed().FullyQualifiedName()
				msgDoc[name] = &MessageDoc{
					Message: key.Embed(),
					Fields:  fn.MessageDoc(key.Embed()),
				}
				for _, field := range key.Embed().Fields() {
					fn.EmbedFields(field, enumDoc, msgDoc)
				}
			}
			if value.IsEnum() {
				name := value.Enum().FullyQualifiedName()
				enumDoc[name] = &EnumDoc{
					Enum:   value.Enum(),
					Values: fn.EnumDoc(value.Enum()),
				}
			} else if value.IsEmbed() {
				name := value.Embed().FullyQualifiedName()
				msgDoc[name] = &MessageDoc{
					Message: value.Embed(),
					Fields:  fn.MessageDoc(value.Embed()),
				}
				for _, field := range value.Embed().Fields() {
					fn.EmbedFields(field, enumDoc, msgDoc)
				}
			}
		} else if field.Type().IsRepeated() {
			el := field.Type().Element()
			if el.IsEnum() {
				name := el.Enum().FullyQualifiedName()
				enumDoc[name] = &EnumDoc{
					Enum:   el.Enum(),
					Values: fn.EnumDoc(el.Enum()),
				}
			} else if el.IsEmbed() {
				name := el.Embed().FullyQualifiedName()
				msgDoc[name] = &MessageDoc{
					Message: el.Embed(),
					Fields:  fn.MessageDoc(el.Embed()),
				}
				for _, field := range el.Embed().Fields() {
					fn.EmbedFields(field, enumDoc, msgDoc)
				}
			}
		} else if field.Type().IsEmbed() {
			name := field.Type().Embed().FullyQualifiedName()
			msgDoc[name] = &MessageDoc{
				Message: field.Type().Embed(),
				Fields:  fn.MessageDoc(field.Type().Embed()),
			}
			for _, field := range field.Type().Embed().Fields() {
				fn.EmbedFields(field, enumDoc, msgDoc)
			}
		}

	case pgs.GroupT:
		// TODO
	}
}

func (fn Func) EmbedMessages(file pgs.File) map[string]map[string]interface{} {
	enumDoc := make(map[string]interface{})
	msgDoc := make(map[string]interface{})

	for _, svc := range file.Services() {
		for _, method := range svc.Methods() {
			fields := make([]pgs.Field, 0, len(method.Input().Fields())+len(method.Output().Fields()))
			for _, field := range method.Input().Fields() {
				fields = append(fields, field)
			}
			for _, field := range method.Output().Fields() {
				fields = append(fields, field)
			}
			for _, field := range fields {
				fn.EmbedFields(field, enumDoc, msgDoc)
			}
		}
	}

	return map[string]map[string]interface{}{
		Enum:    enumDoc,
		Message: msgDoc,
	}
}

type TOCElement struct {
	Interface bool
	Name      pgs.Name
	Gateway   string
	Comment   string
}

func (fn Func) TableOfContent(file pgs.File) []*TOCElement {
	var out []*TOCElement

	for _, svc := range file.Services() {
		out = append(out, &TOCElement{
			Interface: false,
			Name:      svc.Name(),
			Comment:   fn.comment(svc.SourceCodeInfo()),
		})
		for _, method := range svc.Methods() {
			el := &TOCElement{
				Interface: true,
				Name:      method.Name(),
				Comment:   fn.comment(method.SourceCodeInfo()),
			}

			opts := method.Descriptor().GetOptions()
			descs, _ := proto.ExtensionDescs(opts)

			for _, desc := range descs {
				// 72295728 gRPC gateway
				if desc.Field == 72295728 {
					ext, _ := proto.GetExtension(opts, desc)
					if rule, ok := ext.(*annotations.HttpRule); ok {
						switch p := rule.Pattern.(type) {
						case *annotations.HttpRule_Get:
							el.Gateway = p.Get
						case *annotations.HttpRule_Put:
							el.Gateway = p.Put
						case *annotations.HttpRule_Post:
							el.Gateway = p.Post
						case *annotations.HttpRule_Delete:
							el.Gateway = p.Delete
						case *annotations.HttpRule_Patch:
							el.Gateway = p.Patch
						case *annotations.HttpRule_Custom:
							el.Gateway = p.Custom.Path
						}
						break
					}
				}
			}

			out = append(out, el)
		}
	}

	return out
}

type GatewayDoc struct {
	URL          string
	Method       string
	ContentType  string
	JsonRequired bool
}

func (fn Func) GatewayDoc(method pgs.Method) *GatewayDoc {
	opts := method.Descriptor().GetOptions()
	descs, _ := proto.ExtensionDescs(opts)

	for _, desc := range descs {
		// 72295728 gRPC gateway
		if desc.Field == 72295728 {
			ext, _ := proto.GetExtension(opts, desc)
			if rule, ok := ext.(*annotations.HttpRule); ok {
				doc := &GatewayDoc{
					ContentType:  "`application/json`",
				}
				switch p := rule.Pattern.(type) {
				case *annotations.HttpRule_Get:
					doc.Method = fmt.Sprintf("`%s`", http.MethodGet)
					doc.URL = fmt.Sprintf("`%s`", p.Get)
				case *annotations.HttpRule_Put:
					doc.Method = fmt.Sprintf("`%s`", http.MethodPut)
					doc.URL = fmt.Sprintf("`%s`", p.Put)
					doc.JsonRequired = true
				case *annotations.HttpRule_Post:
					doc.Method = fmt.Sprintf("`%s`", http.MethodPost)
					doc.URL = fmt.Sprintf("`%s`", p.Post)
					doc.JsonRequired = true
				case *annotations.HttpRule_Delete:
					doc.Method = fmt.Sprintf("`%s`", http.MethodDelete)
					doc.URL = fmt.Sprintf("`%s`", p.Delete)
				case *annotations.HttpRule_Patch:
					doc.Method = fmt.Sprintf("`%s`", http.MethodPatch)
					doc.URL = fmt.Sprintf("`%s`", p.Patch)
				case *annotations.HttpRule_Custom:
					doc.Method = fmt.Sprintf("`%s`", p.Custom.Kind)
					doc.URL = fmt.Sprintf("`%s`", p.Custom.Path)
				}
				return doc
			}
		}
	}

	return nil
}

func (fn Func) enumJson(enum pgs.Enum) string {
	var val []string

	for _, v := range enum.Values() {
		val = append(val, v.Name().String())
	}

	return strings.Join(val, " | ")
}

func (fn Func) fieldElementJson(el pgs.FieldTypeElem, mapKey bool) string {
	if el.IsEmbed() {
		return fn.messageJson(el.Embed())
	} else if el.IsEnum() {
		if el.Enum().FullyQualifiedName() == ".google.protobuf.NullValue" {
			return "null"
		}
		return fmt.Sprintf(`"%s"`, fn.enumJson(el.Enum()))
	}

	switch el.ProtoType() {
	case pgs.DoubleT, pgs.FloatT:
		if mapKey {
			return `"3.1415926"`
		} else {
			return "3.1415926"
		}
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return `"string($int64)"`
	case pgs.Int32T, pgs.UInt32T, pgs.Fixed32T, pgs.SInt32, pgs.SFixed32:
		if mapKey {
			return `"0"`
		} else {
			return "0"
		}
	case pgs.BoolT:
		if mapKey {
			return `"true"`
		} else {
			return "true"
		}
	case pgs.StringT:
		return `"string"`
	case pgs.BytesT:
		return `"YmFzZTY0IHN0cmluZw=="`
	default:
		return `"UNKNOWN"`
	}
}

func (fn Func) embedJson(message pgs.Message) string {
	switch message.FullyQualifiedName() {
	case ".google.protobuf.Timestamp":
		return `"1972-01-01T10:00:20.021Z"`
	case ".google.protobuf.Duration":
		return `"1.000340012s"`
	case ".google.protobuf.Empty":
		return `"{}"`
	case ".google.protobuf.FieldMask":
		return `"f.fooBar,h"`
	case ".google.protobuf.ListValue":
		return `["foo", "bar"]`
	case ".google.protobuf.DoubleValue", ".google.protobuf.FloatValue":
		return `3.1415926`
	case ".google.protobuf.Int64Value", ".google.protobuf.UInt64Value":
		return `"string($int64)"`
	case ".google.protobuf.Int32Value", ".google.protobuf.UInt32Value":
		return "0"
	case ".google.protobuf.BoolValue":
		return "true"
	case ".google.protobuf.StringValue":
		return `"string"`
	case ".google.protobuf.BytesValue":
		return `"YmFzZTY0IHN0cmluZw=="`
	default:
		return fn.messageJson(message)
	}
}

func (fn Func) fieldJson(field pgs.Field) string {
	switch field.Type().ProtoType() {
	case pgs.DoubleT, pgs.FloatT:
		return `3.1415926`
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return `"string($int64)"`
	case pgs.Int32T, pgs.UInt32T, pgs.Fixed32T, pgs.SInt32, pgs.SFixed32:
		return "0"
	case pgs.BoolT:
		return "true"
	case pgs.StringT:
		return `"string"`
	case pgs.BytesT:
		return `"YmFzZTY0IHN0cmluZw=="`
	case pgs.EnumT:
		enum := field.Type().Enum()
		if enum.FullyQualifiedName() == ".google.protobuf.NullValue" {
			return "null"
		}
		return fmt.Sprintf(`"%s"`, fn.enumJson(enum))
	case pgs.MessageT:
		if field.Type().IsMap() {
			key := fn.fieldElementJson(field.Type().Key(), true)
			value := fn.fieldElementJson(field.Type().Element(), false)
			return fmt.Sprintf(`{%s:%s}`, key, value)
		} else if field.Type().IsRepeated() {
			return fn.fieldElementJson(field.Type().Element(), false)
		} else {
			return fn.embedJson(field.Type().Embed())
		}
	// TODO
	//case pgs.GroupT:
	default:
		return ""
	}
}

func (fn Func) messageJson(message pgs.Message) string {
	var lines []string

	for _, field := range message.Fields() {
		val := fn.fieldJson(field)
		if field.Type().IsRepeated() {
			val = fmt.Sprintf(`[%s]`, val)
		}
		lines = append(lines, fmt.Sprintf(`"%s":%s`, field.Name(), val))
	}

	return fmt.Sprintf(`{%s}`, strings.Join(lines, ","))
}

func (fn Func) JSONDemo(message pgs.Message) string {
	var prettyJSON bytes.Buffer
	jsonVal := fn.messageJson(message)
	if err := json.Indent(&prettyJSON, []byte(jsonVal), "", "\t"); err != nil {
		return "json.Indent err:" + err.Error()
	}
	return fmt.Sprintf("```json\n%s\n```", string(prettyJSON.Bytes()))
}
