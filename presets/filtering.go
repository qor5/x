package presets

import (
	vt "github.com/sunfmin/bran/vuetify"
)

func (b *ListingBuilder) Filter(v vt.FilterData) {
	b.filterData = v
}
