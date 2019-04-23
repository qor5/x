package e01_hello_button

import (
	ui "github.com/sunfmin/page"
	. "github.com/sunfmin/bran/html"
)

func HelloButton(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	pr.Schema = Button("Hello").OnClick(ctx.Hub, "reload", reload)
	return
}

func reload(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	return
}
