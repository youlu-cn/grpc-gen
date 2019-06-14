package shared

import (
	"fmt"
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
	Name     string
	Comment  string
	Type     string
	Default  string
	Required bool
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
		return fmt.Sprintf("[%s](#%s)", enum.Name(), fn.Anchor(enum.Name()))
	}

	switch el.ProtoType() {
	case pgs.DoubleT, pgs.FloatT:
		return "float"
	case pgs.Int64T, pgs.UInt64T,
		pgs.Int32T, pgs.UInt32T,
		pgs.Fixed32T, pgs.Fixed64T,
		pgs.SInt32, pgs.SInt64,
		pgs.SFixed32, pgs.SFixed64:
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
			Type:     field.Type().ProtoType().Proto().String(),
			Required: field.Required(),
		}
		// Type
		switch field.Type().ProtoType() {
		case pgs.DoubleT, pgs.FloatT:
			doc.Type = "float"
		case pgs.Int64T, pgs.UInt64T,
			pgs.Int32T, pgs.UInt32T,
			pgs.Fixed32T, pgs.Fixed64T,
			pgs.SInt32, pgs.SInt64,
			pgs.SFixed32, pgs.SFixed64:
			doc.Type = "int"
		case pgs.BoolT:
			doc.Type = "bool"
		case pgs.StringT:
			doc.Type = "string"
		case pgs.BytesT:
			doc.Type = "bytes"
		case pgs.EnumT:
			enum := field.Type().Enum()
			doc.Type = fmt.Sprintf("[%s](#%s)", enum.Name(), fn.Anchor(enum.Name()))
		case pgs.MessageT:
			if field.Type().IsMap() {
				key := fn.fieldElementType(field.Type().Key())
				value := fn.fieldElementType(field.Type().Element())
				doc.Type = fmt.Sprintf("map\\<%s, %s\\>", key, value)
			} else if field.Type().IsRepeated() {
				el := fn.fieldElementType(field.Type().Element())
				doc.Type = el
			} else {
				msg := field.Type().Embed()
				doc.Type = fmt.Sprintf("[%s](#%s)", msg.Name(), fn.Anchor(msg.Name()))
			}
		case pgs.GroupT:
			// TODO
		}
		if field.Type().IsRepeated() {
			doc.Type += " array"
		}
		// Comment
		doc.Comment = fn.comment(field.SourceCodeInfo())
		out = append(out, doc)
	}

	return out
}

func (fn Func) EmbedMessages(file pgs.File) map[string]map[string]interface{} {
	out := map[string]map[string]interface{}{
		Enum:    make(map[string]interface{}),
		Message: make(map[string]interface{}),
	}

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
				switch field.Type().ProtoType() {
				case pgs.EnumT:
					name := field.Type().Enum().FullyQualifiedName()
					out[Enum][name] = &EnumDoc{
						Enum:   field.Type().Enum(),
						Values: fn.EnumDoc(field.Type().Enum()),
					}
				case pgs.MessageT:
					if field.Type().IsMap() {
						key := field.Type().Key()
						value := field.Type().Element()
						if key.IsEnum() {
							name := key.Enum().FullyQualifiedName()
							out[Enum][name] = &EnumDoc{
								Enum:   key.Enum(),
								Values: fn.EnumDoc(key.Enum()),
							}
						} else if key.IsEmbed() {
							name := key.Embed().FullyQualifiedName()
							out[Message][name] = &MessageDoc{
								Message: key.Embed(),
								Fields:  fn.MessageDoc(key.Embed()),
							}
						}
						if value.IsEnum() {
							name := value.Enum().FullyQualifiedName()
							out[Enum][name] = &EnumDoc{
								Enum:   value.Enum(),
								Values: fn.EnumDoc(value.Enum()),
							}
						} else if value.IsEmbed() {
							name := value.Embed().FullyQualifiedName()
							out[Message][name] = &MessageDoc{
								Message: value.Embed(),
								Fields:  fn.MessageDoc(value.Embed()),
							}
						}
					} else if field.Type().IsRepeated() {
						el := field.Type().Element()
						if el.IsEnum() {
							name := el.Enum().FullyQualifiedName()
							out[Enum][name] = &EnumDoc{
								Enum:   el.Enum(),
								Values: fn.EnumDoc(el.Enum()),
							}
						} else if el.IsEmbed() {
							name := el.Embed().FullyQualifiedName()
							out[Message][name] = &MessageDoc{
								Message: el.Embed(),
								Fields:  fn.MessageDoc(el.Embed()),
							}
						}
					} else if field.Type().IsEmbed() {
						name := field.Type().Embed().FullyQualifiedName()
						out[Message][name] = &MessageDoc{
							Message: field.Type().Embed(),
							Fields:  fn.MessageDoc(field.Type().Embed()),
						}
					}

				case pgs.GroupT:
					// TODO
				}
			}
		}
	}

	return out
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
