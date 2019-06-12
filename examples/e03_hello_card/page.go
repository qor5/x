package e03_hello_card

import (
	. "github.com/sunfmin/bran/material"
	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type mystate struct {
	Message string
}

func HelloCard(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	pr.Schema = Card(
		h.RawHTML(`Text`),
	).ActionButtons(
		Button("Read").InCard(),
		Button("Bookmark").InCard())
	return
}
