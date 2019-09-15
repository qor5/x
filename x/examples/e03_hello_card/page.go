package e03_hello_card

import (
	"github.com/goplaid/web"
	. "github.com/goplaid/x/material"
	h "github.com/theplant/htmlgo"
)

func HelloCard(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = Card(
		h.RawHTML(`Text`),
	).ActionButtons(
		Button("Read").InCard(),
		Button("Bookmark").InCard())
	return
}
