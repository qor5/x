package e03_hello_card

import (
	. "github.com/sunfmin/bran/material"
	ui "github.com/sunfmin/page"
)

type mystate struct {
	Message string
}

func HelloCard(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	pr.Schema = Card(
		ui.RawHTML(`Text`),
	).ActionButtons(
		Button("Read").InCard(),
		Button("Bookmark").InCard())
	return
}
