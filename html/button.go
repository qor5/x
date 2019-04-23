package html

import (
	ui "github.com/sunfmin/page"
)

type ButtonBuilder struct {
	tag *HTMLTagBuilder
}

func Button(label string) (r *ButtonBuilder) {
	r = &ButtonBuilder{
		tag: Tag("button").Text(label),
	}
	return
}

func (b *ButtonBuilder) OnClick(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *ButtonBuilder) {
	b.tag.OnClick(hub, eventFuncId, ef, params...)
	r = b
	return
}

func (b *ButtonBuilder) MarshalHTML(ctx *ui.EventContext) (r []byte, err error) {
	ctx.Head.PutStyle(`
	button {
		color: red;
	}
	`)

	ctx.Head.PutScript(`
		console.log("run")
	`)

	return b.tag.MarshalHTML(ctx)
}
