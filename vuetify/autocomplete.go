package vuetify

import (
	"context"

	"github.com/sunfmin/bran/ui"

	h "github.com/theplant/htmlgo"
)

type TextValue struct {
	Text  string `json:"text,omitempty"`
	Value string `json:"value,omitempty"`
}

type VAutoCompleteBuilder struct {
	tag     *h.HTMLTagBuilder
	dirData *ui.BranDirectiveData
}

func VAutoComplete(v []*TextValue) (r *VAutoCompleteBuilder) {
	r = &VAutoCompleteBuilder{
		tag: h.Tag("v-autocomplete"),
		dirData: &ui.BranDirectiveData{
			SetupFunc: "vuetifyVSelect",
		},
	}
	r.Items(v).
		Multiple(true).
		Chips(true).
		DeletableChips(true).
		Solo(true)
	return
}

func (b *VAutoCompleteBuilder) Items(v []*TextValue) (r *VAutoCompleteBuilder) {
	b.tag.Attr(":items", v)
	return b
}

func (b *VAutoCompleteBuilder) FieldName(v string) (r *VAutoCompleteBuilder) {
	b.dirData.FieldName = v
	return b
}

func (b *VAutoCompleteBuilder) Multiple(v bool) (r *VAutoCompleteBuilder) {
	b.tag.Attr("multiple", v)
	return b
}

func (b *VAutoCompleteBuilder) Chips(v bool) (r *VAutoCompleteBuilder) {
	b.tag.Attr("chips", v)
	return b
}

func (b *VAutoCompleteBuilder) DeletableChips(v bool) (r *VAutoCompleteBuilder) {
	b.tag.Attr("deletable-chips", v)
	return b
}

func (b *VAutoCompleteBuilder) Solo(v bool) (r *VAutoCompleteBuilder) {
	b.tag.Attr("solo", v)
	return b
}

func (b *VAutoCompleteBuilder) Clearable(v bool) (r *VAutoCompleteBuilder) {
	b.tag.Attr("clearable", v)
	return b
}

func (b *VAutoCompleteBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	b.tag.Attr("v-bran", b.dirData)
	return b.tag.MarshalHTML(ctx)
}
