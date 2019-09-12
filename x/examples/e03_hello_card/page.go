package e03_hello_card

import (
	"github.com/sunfmin/bran/web"
	. "github.com/sunfmin/bran/x/material"
	h "github.com/theplant/htmlgo"
)

func HelloCard(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Schema = Card(
		h.RawHTML(`Text`),
	).ActionButtons(
		Button("Read").InCard(),
		Button("Bookmark").InCard())
	return
}
