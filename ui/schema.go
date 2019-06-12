package ui

import (
	"encoding/json"
)

type SchemaComponent interface {
	MarshalSchema(ctx *EventContext) ([]byte, error)
}

type RawSchema string

func (s RawSchema) MarshalSchema(ctx *EventContext) (r []byte, err error) {
	r = []byte(s)
	return
}

func WithContext(ctx *EventContext, comp SchemaComponent) json.Marshaler {
	return &withCtx{ctx, comp}
}

type withCtx struct {
	ctx  *EventContext
	body SchemaComponent
}

func (wc *withCtx) MarshalJSON() ([]byte, error) {
	return wc.body.MarshalSchema(wc.ctx)
}
