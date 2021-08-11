package vuetify

import "fmt"

type DType string

const (
	DTypeFlex       DType = "flex"
	DTypeInlineFlex DType = "inline-flex"
	DTypeBlock      DType = "block"
)

type SizeType string

const (
	Xs SizeType = "xs"
	Sm SizeType = "sm"
	Md SizeType = "md"
	Lg SizeType = "lg"
	Xl SizeType = "xl"
)

type AlignType string

const (
	Left    AlignType = "left"
	Center  AlignType = "center"
	Right   AlignType = "right"
	Justify AlignType = "justify"
)

func (b *VContainerBuilder) DType(v DType) (r *VContainerBuilder) {
	b.tag.Attr(fmt.Sprintf(":d-%s", v), fmt.Sprint(true))
	return b
}

func (b *VContainerBuilder) TextAlign(s SizeType, a AlignType) (r *VContainerBuilder) {
	b.tag.Attr(fmt.Sprintf(":text-%s-%s", s, a), fmt.Sprint(true))
	return b
}

func (b *VLayoutBuilder) DType(v DType) (r *VLayoutBuilder) {
	b.tag.Attr(fmt.Sprintf(":d-%s", v), fmt.Sprint(true))
	return b
}

func (b *VContainerBuilder) GridList(s SizeType) (r *VContainerBuilder) {
	b.tag.Attr(fmt.Sprintf(":grid-list-%s", s), fmt.Sprint(true))
	return b
}
