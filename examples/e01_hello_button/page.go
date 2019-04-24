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
		m.Button("").Variant(m.ButtonVariantOutlined).
			Children(
				ui.StringHTMLComponent(`
				<svg class="mdc-button__icon" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#000000">
                <path fill="none" d="M0 0h24v24H0z"/>
                <path d="M23 12c0-6.07-4.93-11-11-11S1 5.93 1 12s4.93 11 11 11 11-4.93 11-11zM5 17.64C3.75 16.1 3 14.14 3 12c0-2.13.76-4.08 2-5.63v11.27zM17.64 5H6.36C7.9 3.75 9.86 3 12 3s4.1.75 5.64 2zM12 14.53L8.24 7h7.53L12 14.53zM17 9v8h-4l4-8zm-6 8H7V9l4 8zm6.64 2c-1.55 1.25-3.51 2-5.64 2s-4.1-.75-5.64-2h11.28zM21 12c0 2.14-.75 4.1-2 5.64V6.37c1.24 1.55 2 3.5 2 5.63z"/>
			  </svg>
				`),
				ui.StringHTMLComponent("SVG Icon"),
			).
			OnClick(ctx.Hub, "reload7", reload),
	)
	return
}

func reload(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	return
}
