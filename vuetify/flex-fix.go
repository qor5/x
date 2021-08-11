package vuetify

import "fmt"

func (b *VFlexBuilder) Col(size SizeType, columns int) (r *VFlexBuilder) {
	b.tag.Attr(fmt.Sprintf(":%s%d", size, columns), fmt.Sprint(true))
	return b
}

func (b *VFlexBuilder) Offset(size SizeType, columns int) (r *VFlexBuilder) {
	b.tag.Attr(fmt.Sprintf(":offset-%s%d", size, columns), fmt.Sprint(true))
	return b
}

func (b *VFlexBuilder) Order(size SizeType, columns int) (r *VFlexBuilder) {
	b.tag.Attr(fmt.Sprintf(":order-%s%d", size, columns), fmt.Sprint(true))

	return b
}
