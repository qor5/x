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
		ui.StringHTMLComponent(`Text`),
	).ActionButtons(ui.StringHTMLComponent(`
			<button class="mdc-button mdc-card__action mdc-card__action--button">Read</button>
			<button class="mdc-button mdc-card__action mdc-card__action--button">Bookmark</button>
		`))
	return
}
