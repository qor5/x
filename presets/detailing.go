package presets

import (
	"github.com/sunfmin/bran/ui"
)

type DetailingBuilder struct {
	mb         *ModelBuilder
	fieldNames []string
	fields     []*FieldBuilder
	actions    []*ActionBuilder
	pageFunc   ui.PageFunc
	fetcher    FetchOpFunc
}

func (b *ModelBuilder) Detailing(vs ...string) (r *DetailingBuilder) {
	r = b.detailing
	r.fieldNames = vs
	var newfields []*FieldBuilder
	for _, f := range vs {
		newfields = append(newfields, r.Field(f))
	}
	r.fields = newfields
	return r
}

func (b *DetailingBuilder) Field(name string) (r *FieldBuilder) {
	for _, f := range b.fields {
		if f.name == name {
			return f
		}
	}
	r = &FieldBuilder{name: name}
	b.fields = append(b.fields, r)
	return
}

func (b *DetailingBuilder) PageFunc(pf ui.PageFunc) (r *DetailingBuilder) {
	b.pageFunc = pf
	return b
}

func (b *DetailingBuilder) Fetcher(v FetchOpFunc) (r *DetailingBuilder) {
	b.fetcher = v
	return b
}

func (b *DetailingBuilder) GetPageFunc() ui.PageFunc {
	if b.pageFunc != nil {
		return b.pageFunc
	}
	return b.defaultPageFunc
}

func (b *DetailingBuilder) defaultPageFunc(ctx *ui.EventContext) (r ui.PageResponse, err error) {
	return
}
