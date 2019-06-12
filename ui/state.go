package ui

import (
	"github.com/sunfmin/reflectutils"
)

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
