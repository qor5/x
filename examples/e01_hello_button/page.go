package e01_hello_button

import (
	ui "github.com/sunfmin/page"
	. "github.com/sunfmin/vuibuilder/html"
)

func HelloButton(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	pr.Schema = Button()
	return
}
