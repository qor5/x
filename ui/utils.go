package ui

import (
	"bytes"
	"encoding/json"

	"github.com/sunfmin/reflectutils"
)

type RawHTML string

func (s RawHTML) MarshalHTML(ctx *EventContext) (r []byte, err error) {
	r = []byte(s)
	return
}

type RawSchema string

func (s RawSchema) MarshalSchema(ctx *EventContext) (r []byte, err error) {
	r = []byte(s)
	return
}

type ComponentFunc func(ctx *EventContext) (r []byte, err error)

func (f ComponentFunc) MarshalHTML(ctx *EventContext) (r []byte, err error) {
	return f(ctx)
}

func (f ComponentFunc) MarshalSchema(ctx *EventContext) (r []byte, err error) {
	return f(ctx)
}

type HTMLComponents []HTMLComponent

func (hcs HTMLComponents) MarshalHTML(ctx *EventContext) (r []byte, err error) {
	buf := bytes.NewBuffer(nil)
	for _, h := range hcs {
		var b []byte
		b, err = h.MarshalHTML(ctx)
		if err != nil {
			return
		}
		buf.Write(b)
	}
	r = buf.Bytes()
	return
}

func (ctx *EventContext) StateOrInit(v PageState) (r PageState) {
	if ctx.State == nil {
		ctx.State = v
	}
	r = ctx.State
	return
}

func (ctx *EventContext) SubStateOrInit(reflectPath string, v interface{}) (r interface{}) {

	r = reflectutils.MustGet(ctx.State, reflectPath)
	if r == nil {
		err := reflectutils.Set(ctx.State, reflectPath, v)
		if err != nil {
			panic(err)
		}
		r = reflectutils.MustGet(ctx.State, reflectPath)
	}

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
