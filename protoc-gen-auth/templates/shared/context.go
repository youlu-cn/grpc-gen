package shared

import (
	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
)

type RuleContext struct {
	Field pgs.Field
	Rules proto.Message
	//Gogo  Gogo

	Typ        string
	WrapperTyp string

	OnKey            bool
	Index            string
	AccessorOverride string
}
