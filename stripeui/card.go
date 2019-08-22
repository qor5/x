package stripeui

import (
	"context"

	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

type CardBuilder struct {
	children   []h.HTMLComponent
	systemBar  []h.HTMLComponent
	header     []h.HTMLComponent
	actions    []h.HTMLComponent
	classNames []string
}

func Card(children ...h.HTMLComponent) (r *CardBuilder) {
	r = &CardBuilder{}
	r.Children(children...)
	return
}

func (b *CardBuilder) Children(comps ...h.HTMLComponent) (r *CardBuilder) {
	b.children = comps
	return b
}

func (b *CardBuilder) Actions(actions ...h.HTMLComponent) (r *CardBuilder) {
	b.actions = actions
	return b
}

func (b *CardBuilder) Header(header ...h.HTMLComponent) (r *CardBuilder) {
	b.header = header
	return b
}

func (b *CardBuilder) HeaderTitle(title string) (r *CardBuilder) {
	b.header = []h.HTMLComponent{h.Text(title)}
	return b
}

func (b *CardBuilder) SystemBar(systemBar ...h.HTMLComponent) (r *CardBuilder) {
	b.systemBar = systemBar
	return b
}

func (b *CardBuilder) Class(names ...string) (r *CardBuilder) {
	b.classNames = names
	return b
}

func (b *CardBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	var sb h.HTMLComponent
	if len(b.systemBar) > 0 {
		sb = VSystemBar(b.systemBar...).Class("mx-2 pt-4").Color("white").Height(32)
	}

	return VCard(
		sb,
		VToolbar(
			VToolbarTitle("").Children(b.header...),
			VSpacer(),
		).Flat(true).Dense(true).AppendChildren(b.actions...),
		VDivider(),
	).Class(b.classNames...).AppendChildren(b.children...).MarshalHTML(ctx)
}
