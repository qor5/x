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
}

func Button(text string) (r *ButtonBuilder) {
	r = &ButtonBuilder{
		tag: h.Tag("button").Text(text),
	}
	return
}

func (b *ButtonBuilder) ClassNames(names ...string) (r *ButtonBuilder) {
	b.tag.ClassNames(names...)
	r = b
	return
}

func (b *ButtonBuilder) Children(comps ...ui.HTMLComponent) (r *ButtonBuilder) {
	b.tag.Children(comps...)
	r = b
	return
}

func (b *ButtonBuilder) OnClick(hub ui.EventFuncHub, eventFuncId string, ef ui.EventFunc, params ...string) (r *ButtonBuilder) {
	b.tag.OnClick(hub, eventFuncId, ef, params...)
	r = b
	return
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
	r = b
	return
}

func (b *ButtonBuilder) MarshalHTML(ctx *ui.EventContext) (r []byte, err error) {
	b.classNames = append(b.classNames, "mdc-button")
	if len(b.variant) > 0 {
		b.classNames = append(b.classNames, fmt.Sprintf("mdc-button--%s", b.variant))
	}
	b.tag.ClassNames(b.classNames...)

	return b.tag.MarshalHTML(ctx)
}
