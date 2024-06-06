package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type KeyFieldBuilder struct {
	label    string
	icon     h.HTMLComponent
	children []h.HTMLComponent
}

func KeyField(children ...h.HTMLComponent) (r *KeyFieldBuilder) {
	r = &KeyFieldBuilder{}
	r.Children(children...)
	return
}

func (b *KeyFieldBuilder) Label(v string) (r *KeyFieldBuilder) {
	b.label = v
	return b
}

func (b *KeyFieldBuilder) Icon(v h.HTMLComponent) (r *KeyFieldBuilder) {
	b.icon = v
	return b
}

func (b *KeyFieldBuilder) Children(comps ...h.HTMLComponent) (r *KeyFieldBuilder) {
	b.children = comps
	return b
}

func (b *KeyFieldBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return h.Div(
		h.Label(b.label).Class("blue-grey--text lighten-3"),
		h.Div(b.children...).PrependChildren(b.icon),
	).Class("px-4 my-4").Style("border-right: 1px solid #E0E0E0").
		MarshalHTML(ctx)
}

type KeyInfoBuilder struct {
	children []h.HTMLComponent
}

func KeyInfo(children ...h.HTMLComponent) (r *KeyInfoBuilder) {
	r = &KeyInfoBuilder{}
	r.Children(children...)
	return
}

func (b *KeyInfoBuilder) Children(comps ...h.HTMLComponent) (r *KeyInfoBuilder) {
	b.children = comps
	return b
}

func (b *KeyInfoBuilder) Append(label string, comp h.HTMLComponent) (r *KeyInfoBuilder) {
	b.children = append(b.children, KeyField(comp).Label(label))
	return b
}

func (b *KeyInfoBuilder) AppendIcon(label string, icon h.HTMLComponent, comp h.HTMLComponent) (r *KeyInfoBuilder) {
	b.children = append(b.children, KeyField(comp).Label(label).Icon(icon))
	return b
}

func (b *KeyInfoBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return h.Div(b.children...).
		Class("grey lighten-5 d-flex").
		MarshalHTML(ctx)
}
