package e01_hello_button

import (
	. "github.com/sunfmin/bran/html"
	m "github.com/sunfmin/bran/material"
	ui "github.com/sunfmin/page"
)

type mystate struct {
	Message string
}

func HelloButton(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{}).(*mystate)

	pr.Schema = Tag("div").Children(
		Button("Hello").
			OnClick(ctx.Hub, "reload", reload),
		Tag("input").
			OnInput(ctx.Hub, "reload2", reload).
			Attr("type", "text").
			FieldName("Message").
			Attr("value", s.Message),
		Tag("div").
			Style("font-family: monospace;").
			Text(s.Message),
		m.Button("Reload").Variant(m.ButtonVariantRaised).
			OnClick(ctx.Hub, "reload3", reload),
		m.Button("Reload").Variant(m.ButtonVariantUnelevated).
			OnClick(ctx.Hub, "reload4", reload),
		m.Button("Reload").Variant(m.ButtonVariantOutlined).
			OnClick(ctx.Hub, "reload5", reload),
		m.Button("Reload").Variant(m.ButtonVariantText).
			OnClick(ctx.Hub, "reload6", reload),
	)
	return
}

func reload(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	return
}
