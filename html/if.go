package html

import (
	"github.com/sunfmin/bran/ui"
)

type IfBuilder struct {
	comps []ui.HTMLComponent
	set   bool
}

func If(v bool, comps ...ui.HTMLComponent) (r *IfBuilder) {
	r = &IfBuilder{}
	if v {
		r.comps = comps
		r.set = true
	}
	return
}

func (b *IfBuilder) ElseIf(v bool, comps ...ui.HTMLComponent) (r *IfBuilder) {
	if b.set {
		return b
	}
	if v {
		b.comps = comps
		b.set = true
	}
	return b
}

func (b *IfBuilder) Else(comps ...ui.HTMLComponent) (r *IfBuilder) {
	if b.set {
		return b
	}
	b.set = true
	b.comps = comps
	return b
}

func (b *IfBuilder) MarshalHTML(ctx *ui.EventContext) (r []byte, err error) {
	if len(b.comps) == 0 {
		return
	}
	return ui.HTMLComponents(b.comps).MarshalHTML(ctx)
}
