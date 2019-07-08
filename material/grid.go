package material

import (
	"context"
	"fmt"

	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type GridBuilder struct {
	classNames     []string
	children       []h.HTMLComponent
	styles         *ui.Styles
	align          string
	fixColumnWidth string
	innerOnly      bool
}

func Grid(cells ...h.HTMLComponent) (r *GridBuilder) {
	r = &GridBuilder{}
	r.Children(cells...)
	r.styles = &ui.Styles{}
	return
}

type Screen string

const (
	ScreenAll     Screen = ""
	ScreenDesktop Screen = "desktop"
	ScreenTablet  Screen = "tablet"
	ScreenPhone   Screen = "phone"
)

func (b *GridBuilder) Class(names ...string) (r *GridBuilder) {
	b.classNames = names
	return b
}

func (b *GridBuilder) AlignCenter() (r *GridBuilder) {
	b.align = ""
	return b
}

func (b *GridBuilder) FixColumnWidth() (r *GridBuilder) {
	b.fixColumnWidth = "mdc-layout-grid--fixed-column-width"
	return b
}

func (b *GridBuilder) Inner() (r *GridBuilder) {
	b.innerOnly = true
	return b
}

func (b *GridBuilder) AlignLeft() (r *GridBuilder) {
	b.align = "mdc-layout-grid--align-left"
	return b
}

func (b *GridBuilder) AlignRight() (r *GridBuilder) {
	b.align = "mdc-layout-grid--align-right"
	return b
}

func (b *GridBuilder) Margin(v int, screen Screen) (r *GridBuilder) {
	if screen == ScreenAll || screen == ScreenDesktop {
		b.styles.Put("--mdc-layout-grid-margin-desktop", fmt.Sprintf("%dpx", v))
	}
	if screen == ScreenAll || screen == ScreenTablet {
		b.styles.Put("--mdc-layout-grid-margin-tablet", fmt.Sprintf("%dpx", v))
	}
	if screen == ScreenAll || screen == ScreenPhone {
		b.styles.Put("--mdc-layout-grid-margin-phone", fmt.Sprintf("%dpx", v))
	}
	return b
}

func (b *GridBuilder) Gutter(v int, screen Screen) (r *GridBuilder) {
	if screen == ScreenAll || screen == ScreenDesktop {
		b.styles.Put("--mdc-layout-grid-gutter-desktop", fmt.Sprintf("%dpx", v))
	}
	if screen == ScreenAll || screen == ScreenTablet {
		b.styles.Put("--mdc-layout-grid-gutter-tablet", fmt.Sprintf("%dpx", v))
	}
	if screen == ScreenAll || screen == ScreenPhone {
		b.styles.Put("--mdc-layout-grid-gutter-phone", fmt.Sprintf("%dpx", v))
	}
	return b
}

func (b *GridBuilder) Children(comps ...h.HTMLComponent) (r *GridBuilder) {
	b.children = comps
	return b
}

func (b *GridBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {

	inner := h.Div(b.children...).Class("mdc-layout-grid__inner")

	if b.innerOnly {
		return inner.MarshalHTML(ctx)
	}

	root := h.Div(inner).
		Class(append(b.classNames, "mdc-layout-grid", b.align, b.fixColumnWidth)...).
		Style(b.styles.String())

	return root.MarshalHTML(ctx)
}

type CellBuilder struct {
	classNames []string
	children   []h.HTMLComponent
	spans      []string
}

func Cell(children ...h.HTMLComponent) (r *CellBuilder) {
	r = &CellBuilder{}
	r.Children(children...)
	return
}

func (b *CellBuilder) Class(names ...string) (r *CellBuilder) {
	b.classNames = names
	return b
}

func (b *CellBuilder) Span(v int, screen Screen) (r *CellBuilder) {
	var span = fmt.Sprintf("mdc-layout-grid__cell--span-%d", v)
	if screen != ScreenAll {
		span = span + "-" + string(screen)
	}

	for _, s := range b.spans {
		if s == span {
			return
		}
	}

	b.spans = append(b.spans, span)
	return b
}

func (b *CellBuilder) Children(comps ...h.HTMLComponent) (r *CellBuilder) {
	b.children = comps
	return b
}

func (b *CellBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	root := h.Div(b.children...).
		Class(append(append(
			b.classNames,
			"mdc-layout-grid__cell",
		), b.spans...)...)
	return root.MarshalHTML(ctx)
}
