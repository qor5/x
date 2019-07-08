package vuetify

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type TextValue struct {
	Text  string `json:"text,omitempty"`
	Value string `json:"value,omitempty"`
}

type VAutoCompleteBuilder struct {
	tag *h.HTMLTagBuilder
}

func VAutoComplete(v []*TextValue) (r *VAutoCompleteBuilder) {
	r = &VAutoCompleteBuilder{
		tag: h.Tag("vw-autocomplete"),
	}
	r.Items(v).
		Multiple(true)
	return
}

func (b *VAutoCompleteBuilder) Items(v []*TextValue) (r *VAutoCompleteBuilder) {
	b.tag.Attr(":items", v)
	return b
}

func (b *VAutoCompleteBuilder) FieldName(v string) (r *VAutoCompleteBuilder) {
	b.tag.Attr("field-name", v)
	return b
}

func (b *VAutoCompleteBuilder) Multiple(v bool) (r *VAutoCompleteBuilder) {
	b.tag.Attr("multiple", v)
	return b
}

// func (b *VAutoCompleteBuilder) Chips(v bool) (r *VAutoCompleteBuilder) {
// 	b.tag.Attr("chips", v)
// 	return b
// }

// func (b *VAutoCompleteBuilder) DeletableChips(v bool) (r *VAutoCompleteBuilder) {
// 	b.tag.Attr("deletable-chips", v)
// 	return b
// }

// func (b *VAutoCompleteBuilder) Solo(v bool) (r *VAutoCompleteBuilder) {
// 	b.tag.Attr("solo", v)
// 	return b
// }

// func (b *VAutoCompleteBuilder) Clearable(v bool) (r *VAutoCompleteBuilder) {
// 	b.tag.Attr("clearable", v)
// 	return b
// }

func (b *VAutoCompleteBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
