package material

import (
	h "github.com/sunfmin/bran/html"
	ui "github.com/sunfmin/page"
)

type CardBuilder struct {
	children   []ui.HTMLComponent
	classNames []string

	actionButtons []ui.HTMLComponent
	actionIcons   []ui.HTMLComponent
}

func (b *CardBuilder) ClassNames(names ...string) (r *CardBuilder) {
	b.classNames = names
	r = b
	return
}

func Card(children ...ui.HTMLComponent) (r *CardBuilder) {
	r = &CardBuilder{}
	r.children = children
	return
}

func (b *CardBuilder) ActionButtons(buttons ...ui.HTMLComponent) (r *CardBuilder) {
	b.actionButtons = buttons
	r = b
	return
}

func (b *CardBuilder) ActionIcons(icons ...ui.HTMLComponent) (r *CardBuilder) {
	b.actionIcons = icons
	r = b
	return
}

func (b *CardBuilder) MarshalHTML(ctx *ui.EventContext) (r []byte, err error) {
	root := h.Tag("div").
		ClassNames(append([]string{"mdc-card"}, b.classNames...)...).
		Children(
			h.Tag("div").ClassNames("mdc-card__primary-action").Attr("tabindex", "0").
				Children(
					b.children...,
				),
		)

	if len(b.actionButtons) > 0 || len(b.actionIcons) > 0 {
		actions := h.Tag("div").ClassNames("mdc-card__actions")
		root.AddChildren(actions)

		if len(b.actionButtons) > 0 {
			actions.AddChildren(
				h.Tag("div").ClassNames("mdc-card__action-buttons").
					Children(b.actionButtons...),
			)
		}

		if len(b.actionIcons) > 0 {
			actions.AddChildren(
				h.Tag("div").ClassNames("mdc-card__action-icons").
					Children(b.actionIcons...),
			)
		}
	}

	return root.MarshalHTML(ctx)
}
