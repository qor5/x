package material

import (
	"fmt"

	h "github.com/sunfmin/bran/html"
	ui "github.com/sunfmin/page"
)

type ButtonBuilder struct {
	tag        *h.HTMLTagBuilder
	classNames []string
	variant    ButtonVariant
	inCard     bool
	disabled   bool
}

func Button(text string) (r *ButtonBuilder) {
	r = &ButtonBuilder{
		tag:     h.Tag("button").Text(text),
		variant: ButtonVariantText,
	}
	return
}

func (b *ButtonBuilder) ClassNames(names ...string) (r *ButtonBuilder) {
	b.classNames = names
	return b
}

func (b *ButtonBuilder) Disabled(v bool) (r *ButtonBuilder) {
	b.disabled = v
	return b
}

func (b *ButtonBuilder) InCard() (r *ButtonBuilder) {
	b.inCard = true
	return b
}

func (b *ButtonBuilder) Children(comps ...ui.HTMLComponent) (r *ButtonBuilder) {
	b.tag.Children(comps...)
	return b
}

func (b *ButtonBuilder) OnClick(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *ButtonBuilder) {
	b.tag.OnClick(hub, eventFuncId, ef, params...)
	return b
}

type ButtonVariant string

const (
	ButtonVariantText       ButtonVariant = "text"
	ButtonVariantOutlined   ButtonVariant = "outlined"
	ButtonVariantRaised     ButtonVariant = "raised"
	ButtonVariantUnelevated ButtonVariant = "unelevated"
)

func (b *ButtonBuilder) Variant(v ButtonVariant) (r *ButtonBuilder) {
	b.variant = v
	return b
}

func (b *ButtonBuilder) MarshalHTML(ctx *ui.EventContext) (r []byte, err error) {
	b.classNames = append(b.classNames, "mdc-button")
	if b.inCard {
		b.classNames = append(b.classNames, "mdc-card__action", "mdc-card__action--button")
	}
	if len(b.variant) > 0 {
		b.classNames = append(b.classNames, fmt.Sprintf("mdc-button--%s", b.variant))
	}
	b.tag.ClassNames(b.classNames...)
	if b.disabled {
		b.tag.Attr("disabled", "true")
	}
	return b.tag.MarshalHTML(ctx)
}
