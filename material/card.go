package material

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type CardBuilder struct {
	children   []h.HTMLComponent
	classNames []string

	actionButtons []h.HTMLComponent
	actionIcons   []h.HTMLComponent
}

func (b *CardBuilder) Class(names ...string) (r *CardBuilder) {
	b.classNames = names
	return b
}

func Card(children ...h.HTMLComponent) (r *CardBuilder) {
	r = &CardBuilder{}
	r.children = children
	return
}

func (b *CardBuilder) ActionButtons(buttons ...h.HTMLComponent) (r *CardBuilder) {
	b.actionButtons = buttons
	return b
}

func (b *CardBuilder) ActionIcons(icons ...h.HTMLComponent) (r *CardBuilder) {
	b.actionIcons = icons
	return b
}

func (b *CardBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	root := h.Div(
		h.Div(b.children...).
			Class("mdc-card__primary-action").
			Attr("tabindex", "0"),
	).Class(append([]string{"mdc-card"}, b.classNames...)...)

	if len(b.actionButtons) > 0 || len(b.actionIcons) > 0 {
		actions := h.Div().Class("mdc-card__actions")
		root.AppendChildren(actions)

		if len(b.actionButtons) > 0 {
			actions.AppendChildren(
				h.Div().Class("mdc-card__action-buttons").
					Children(b.actionButtons...),
			)
		}

		if len(b.actionIcons) > 0 {
			actions.AppendChildren(
				h.Div().Class("mdc-card__action-icons").
					Children(b.actionIcons...),
			)
		}
	}

	return root.MarshalHTML(ctx)
}
