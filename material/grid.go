package material

import (
	"fmt"

	h "github.com/sunfmin/bran/html"
	ui "github.com/sunfmin/pagui"
)

type GridBuilder struct {
	classNames     []string
	children       []ui.HTMLComponent
	styles         *h.Styles
	align          string
	fixColumnWidth string
	innerOnly      bool
}

func Grid(cells ...ui.HTMLComponent) (r *GridBuilder) {
	r = &GridBuilder{}
	r.Children(cells...)
	r.styles = &h.Styles{}
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

func (b *GridBuilder) Children(comps ...ui.HTMLComponent) (r *GridBuilder) {
	b.children = comps
	return b
}

func (b *GridBuilder) MarshalHTML(ctx *ui.EventContext) (r []byte, err error) {
	ctx.Head.PutStyle(gridcss)

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
	children   []ui.HTMLComponent
	spans      []string
}

func Cell(children ...ui.HTMLComponent) (r *CellBuilder) {
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

func (b *CellBuilder) Children(comps ...ui.HTMLComponent) (r *CellBuilder) {
	b.children = comps
	return b
}

func (b *CellBuilder) MarshalHTML(ctx *ui.EventContext) (r []byte, err error) {
	root := h.Div(b.children...).
		Class(append(append(
			b.classNames,
			"mdc-layout-grid__cell",
		), b.spans...)...)
	return root.MarshalHTML(ctx)
}

const gridcss = `
/*!
 Material Components for the Web
 Copyright (c) 2019 Google Inc.
 License: MIT
*/
:root {
  --mdc-layout-grid-margin-desktop: 24px;
  --mdc-layout-grid-gutter-desktop: 24px;
  --mdc-layout-grid-column-width-desktop: 72px;
  --mdc-layout-grid-margin-tablet: 16px;
  --mdc-layout-grid-gutter-tablet: 16px;
  --mdc-layout-grid-column-width-tablet: 72px;
  --mdc-layout-grid-margin-phone: 16px;
  --mdc-layout-grid-gutter-phone: 16px;
  --mdc-layout-grid-column-width-phone: 72px;
}

@media (min-width: 840px) {
  .mdc-layout-grid {
    box-sizing: border-box;
    margin: 0 auto;
    padding: 24px;
    padding: var(--mdc-layout-grid-margin-desktop, 24px);
  }
}
@media (min-width: 480px) and (max-width: 839px) {
  .mdc-layout-grid {
    box-sizing: border-box;
    margin: 0 auto;
    padding: 16px;
    padding: var(--mdc-layout-grid-margin-tablet, 16px);
  }
}
@media (max-width: 479px) {
  .mdc-layout-grid {
    box-sizing: border-box;
    margin: 0 auto;
    padding: 16px;
    padding: var(--mdc-layout-grid-margin-phone, 16px);
  }
}

@media (min-width: 840px) {
  .mdc-layout-grid__inner {
    display: flex;
    flex-flow: row wrap;
    align-items: stretch;
    margin: -12px;
    margin: calc(var(--mdc-layout-grid-gutter-desktop, 24px) / 2 * -1);
  }
  @supports (display: grid) {
    .mdc-layout-grid__inner {
      display: grid;
      margin: 0;
      grid-gap: 24px;
      grid-gap: var(--mdc-layout-grid-gutter-desktop, 24px);
      grid-template-columns: repeat(12, minmax(0, 1fr));
    }
  }
}
@media (min-width: 480px) and (max-width: 839px) {
  .mdc-layout-grid__inner {
    display: flex;
    flex-flow: row wrap;
    align-items: stretch;
    margin: -8px;
    margin: calc(var(--mdc-layout-grid-gutter-tablet, 16px) / 2 * -1);
  }
  @supports (display: grid) {
    .mdc-layout-grid__inner {
      display: grid;
      margin: 0;
      grid-gap: 16px;
      grid-gap: var(--mdc-layout-grid-gutter-tablet, 16px);
      grid-template-columns: repeat(8, minmax(0, 1fr));
    }
  }
}
@media (max-width: 479px) {
  .mdc-layout-grid__inner {
    display: flex;
    flex-flow: row wrap;
    align-items: stretch;
    margin: -8px;
    margin: calc(var(--mdc-layout-grid-gutter-phone, 16px) / 2 * -1);
  }
  @supports (display: grid) {
    .mdc-layout-grid__inner {
      display: grid;
      margin: 0;
      grid-gap: 16px;
      grid-gap: var(--mdc-layout-grid-gutter-phone, 16px);
      grid-template-columns: repeat(4, minmax(0, 1fr));
    }
  }
}

@media (min-width: 840px) {
  .mdc-layout-grid__cell {
    width: calc(33.3333333333% - 24px);
    width: calc(33.3333333333% - var(--mdc-layout-grid-gutter-desktop, 24px));
    box-sizing: border-box;
    margin: 12px;
    margin: calc(var(--mdc-layout-grid-gutter-desktop, 24px) / 2);
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell {
      width: auto;
      grid-column-end: span 4;
    }
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell {
      margin: 0;
    }
  }
  .mdc-layout-grid__cell--span-1,
.mdc-layout-grid__cell--span-1-desktop {
    width: calc(8.3333333333% - 24px);
    width: calc(8.3333333333% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-1,
.mdc-layout-grid__cell--span-1-desktop {
      width: auto;
      grid-column-end: span 1;
    }
  }

  .mdc-layout-grid__cell--span-2,
.mdc-layout-grid__cell--span-2-desktop {
    width: calc(16.6666666667% - 24px);
    width: calc(16.6666666667% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-2,
.mdc-layout-grid__cell--span-2-desktop {
      width: auto;
      grid-column-end: span 2;
    }
  }

  .mdc-layout-grid__cell--span-3,
.mdc-layout-grid__cell--span-3-desktop {
    width: calc(25% - 24px);
    width: calc(25% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-3,
.mdc-layout-grid__cell--span-3-desktop {
      width: auto;
      grid-column-end: span 3;
    }
  }

  .mdc-layout-grid__cell--span-4,
.mdc-layout-grid__cell--span-4-desktop {
    width: calc(33.3333333333% - 24px);
    width: calc(33.3333333333% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-4,
.mdc-layout-grid__cell--span-4-desktop {
      width: auto;
      grid-column-end: span 4;
    }
  }

  .mdc-layout-grid__cell--span-5,
.mdc-layout-grid__cell--span-5-desktop {
    width: calc(41.6666666667% - 24px);
    width: calc(41.6666666667% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-5,
.mdc-layout-grid__cell--span-5-desktop {
      width: auto;
      grid-column-end: span 5;
    }
  }

  .mdc-layout-grid__cell--span-6,
.mdc-layout-grid__cell--span-6-desktop {
    width: calc(50% - 24px);
    width: calc(50% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-6,
.mdc-layout-grid__cell--span-6-desktop {
      width: auto;
      grid-column-end: span 6;
    }
  }

  .mdc-layout-grid__cell--span-7,
.mdc-layout-grid__cell--span-7-desktop {
    width: calc(58.3333333333% - 24px);
    width: calc(58.3333333333% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-7,
.mdc-layout-grid__cell--span-7-desktop {
      width: auto;
      grid-column-end: span 7;
    }
  }

  .mdc-layout-grid__cell--span-8,
.mdc-layout-grid__cell--span-8-desktop {
    width: calc(66.6666666667% - 24px);
    width: calc(66.6666666667% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-8,
.mdc-layout-grid__cell--span-8-desktop {
      width: auto;
      grid-column-end: span 8;
    }
  }

  .mdc-layout-grid__cell--span-9,
.mdc-layout-grid__cell--span-9-desktop {
    width: calc(75% - 24px);
    width: calc(75% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-9,
.mdc-layout-grid__cell--span-9-desktop {
      width: auto;
      grid-column-end: span 9;
    }
  }

  .mdc-layout-grid__cell--span-10,
.mdc-layout-grid__cell--span-10-desktop {
    width: calc(83.3333333333% - 24px);
    width: calc(83.3333333333% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-10,
.mdc-layout-grid__cell--span-10-desktop {
      width: auto;
      grid-column-end: span 10;
    }
  }

  .mdc-layout-grid__cell--span-11,
.mdc-layout-grid__cell--span-11-desktop {
    width: calc(91.6666666667% - 24px);
    width: calc(91.6666666667% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-11,
.mdc-layout-grid__cell--span-11-desktop {
      width: auto;
      grid-column-end: span 11;
    }
  }

  .mdc-layout-grid__cell--span-12,
.mdc-layout-grid__cell--span-12-desktop {
    width: calc(100% - 24px);
    width: calc(100% - var(--mdc-layout-grid-gutter-desktop, 24px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-12,
.mdc-layout-grid__cell--span-12-desktop {
      width: auto;
      grid-column-end: span 12;
    }
  }
}
@media (min-width: 480px) and (max-width: 839px) {
  .mdc-layout-grid__cell {
    width: calc(50% - 16px);
    width: calc(50% - var(--mdc-layout-grid-gutter-tablet, 16px));
    box-sizing: border-box;
    margin: 8px;
    margin: calc(var(--mdc-layout-grid-gutter-tablet, 16px) / 2);
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell {
      width: auto;
      grid-column-end: span 4;
    }
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell {
      margin: 0;
    }
  }
  .mdc-layout-grid__cell--span-1,
.mdc-layout-grid__cell--span-1-tablet {
    width: calc(12.5% - 16px);
    width: calc(12.5% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-1,
.mdc-layout-grid__cell--span-1-tablet {
      width: auto;
      grid-column-end: span 1;
    }
  }

  .mdc-layout-grid__cell--span-2,
.mdc-layout-grid__cell--span-2-tablet {
    width: calc(25% - 16px);
    width: calc(25% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-2,
.mdc-layout-grid__cell--span-2-tablet {
      width: auto;
      grid-column-end: span 2;
    }
  }

  .mdc-layout-grid__cell--span-3,
.mdc-layout-grid__cell--span-3-tablet {
    width: calc(37.5% - 16px);
    width: calc(37.5% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-3,
.mdc-layout-grid__cell--span-3-tablet {
      width: auto;
      grid-column-end: span 3;
    }
  }

  .mdc-layout-grid__cell--span-4,
.mdc-layout-grid__cell--span-4-tablet {
    width: calc(50% - 16px);
    width: calc(50% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-4,
.mdc-layout-grid__cell--span-4-tablet {
      width: auto;
      grid-column-end: span 4;
    }
  }

  .mdc-layout-grid__cell--span-5,
.mdc-layout-grid__cell--span-5-tablet {
    width: calc(62.5% - 16px);
    width: calc(62.5% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-5,
.mdc-layout-grid__cell--span-5-tablet {
      width: auto;
      grid-column-end: span 5;
    }
  }

  .mdc-layout-grid__cell--span-6,
.mdc-layout-grid__cell--span-6-tablet {
    width: calc(75% - 16px);
    width: calc(75% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-6,
.mdc-layout-grid__cell--span-6-tablet {
      width: auto;
      grid-column-end: span 6;
    }
  }

  .mdc-layout-grid__cell--span-7,
.mdc-layout-grid__cell--span-7-tablet {
    width: calc(87.5% - 16px);
    width: calc(87.5% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-7,
.mdc-layout-grid__cell--span-7-tablet {
      width: auto;
      grid-column-end: span 7;
    }
  }

  .mdc-layout-grid__cell--span-8,
.mdc-layout-grid__cell--span-8-tablet {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-8,
.mdc-layout-grid__cell--span-8-tablet {
      width: auto;
      grid-column-end: span 8;
    }
  }

  .mdc-layout-grid__cell--span-9,
.mdc-layout-grid__cell--span-9-tablet {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-9,
.mdc-layout-grid__cell--span-9-tablet {
      width: auto;
      grid-column-end: span 8;
    }
  }

  .mdc-layout-grid__cell--span-10,
.mdc-layout-grid__cell--span-10-tablet {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-10,
.mdc-layout-grid__cell--span-10-tablet {
      width: auto;
      grid-column-end: span 8;
    }
  }

  .mdc-layout-grid__cell--span-11,
.mdc-layout-grid__cell--span-11-tablet {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-11,
.mdc-layout-grid__cell--span-11-tablet {
      width: auto;
      grid-column-end: span 8;
    }
  }

  .mdc-layout-grid__cell--span-12,
.mdc-layout-grid__cell--span-12-tablet {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-tablet, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-12,
.mdc-layout-grid__cell--span-12-tablet {
      width: auto;
      grid-column-end: span 8;
    }
  }
}
@media (max-width: 479px) {
  .mdc-layout-grid__cell {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-phone, 16px));
    box-sizing: border-box;
    margin: 8px;
    margin: calc(var(--mdc-layout-grid-gutter-phone, 16px) / 2);
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell {
      width: auto;
      grid-column-end: span 4;
    }
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell {
      margin: 0;
    }
  }
  .mdc-layout-grid__cell--span-1,
.mdc-layout-grid__cell--span-1-phone {
    width: calc(25% - 16px);
    width: calc(25% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-1,
.mdc-layout-grid__cell--span-1-phone {
      width: auto;
      grid-column-end: span 1;
    }
  }

  .mdc-layout-grid__cell--span-2,
.mdc-layout-grid__cell--span-2-phone {
    width: calc(50% - 16px);
    width: calc(50% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-2,
.mdc-layout-grid__cell--span-2-phone {
      width: auto;
      grid-column-end: span 2;
    }
  }

  .mdc-layout-grid__cell--span-3,
.mdc-layout-grid__cell--span-3-phone {
    width: calc(75% - 16px);
    width: calc(75% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-3,
.mdc-layout-grid__cell--span-3-phone {
      width: auto;
      grid-column-end: span 3;
    }
  }

  .mdc-layout-grid__cell--span-4,
.mdc-layout-grid__cell--span-4-phone {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-4,
.mdc-layout-grid__cell--span-4-phone {
      width: auto;
      grid-column-end: span 4;
    }
  }

  .mdc-layout-grid__cell--span-5,
.mdc-layout-grid__cell--span-5-phone {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-5,
.mdc-layout-grid__cell--span-5-phone {
      width: auto;
      grid-column-end: span 4;
    }
  }

  .mdc-layout-grid__cell--span-6,
.mdc-layout-grid__cell--span-6-phone {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-6,
.mdc-layout-grid__cell--span-6-phone {
      width: auto;
      grid-column-end: span 4;
    }
  }

  .mdc-layout-grid__cell--span-7,
.mdc-layout-grid__cell--span-7-phone {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-7,
.mdc-layout-grid__cell--span-7-phone {
      width: auto;
      grid-column-end: span 4;
    }
  }

  .mdc-layout-grid__cell--span-8,
.mdc-layout-grid__cell--span-8-phone {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-8,
.mdc-layout-grid__cell--span-8-phone {
      width: auto;
      grid-column-end: span 4;
    }
  }

  .mdc-layout-grid__cell--span-9,
.mdc-layout-grid__cell--span-9-phone {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-9,
.mdc-layout-grid__cell--span-9-phone {
      width: auto;
      grid-column-end: span 4;
    }
  }

  .mdc-layout-grid__cell--span-10,
.mdc-layout-grid__cell--span-10-phone {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-10,
.mdc-layout-grid__cell--span-10-phone {
      width: auto;
      grid-column-end: span 4;
    }
  }

  .mdc-layout-grid__cell--span-11,
.mdc-layout-grid__cell--span-11-phone {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-11,
.mdc-layout-grid__cell--span-11-phone {
      width: auto;
      grid-column-end: span 4;
    }
  }

  .mdc-layout-grid__cell--span-12,
.mdc-layout-grid__cell--span-12-phone {
    width: calc(100% - 16px);
    width: calc(100% - var(--mdc-layout-grid-gutter-phone, 16px));
  }
  @supports (display: grid) {
    .mdc-layout-grid__cell--span-12,
.mdc-layout-grid__cell--span-12-phone {
      width: auto;
      grid-column-end: span 4;
    }
  }
}
.mdc-layout-grid__cell--order-1 {
  order: 1;
}
.mdc-layout-grid__cell--order-2 {
  order: 2;
}
.mdc-layout-grid__cell--order-3 {
  order: 3;
}
.mdc-layout-grid__cell--order-4 {
  order: 4;
}
.mdc-layout-grid__cell--order-5 {
  order: 5;
}
.mdc-layout-grid__cell--order-6 {
  order: 6;
}
.mdc-layout-grid__cell--order-7 {
  order: 7;
}
.mdc-layout-grid__cell--order-8 {
  order: 8;
}
.mdc-layout-grid__cell--order-9 {
  order: 9;
}
.mdc-layout-grid__cell--order-10 {
  order: 10;
}
.mdc-layout-grid__cell--order-11 {
  order: 11;
}
.mdc-layout-grid__cell--order-12 {
  order: 12;
}
.mdc-layout-grid__cell--align-top {
  align-self: flex-start;
}
@supports (display: grid) {
  .mdc-layout-grid__cell--align-top {
    align-self: start;
  }
}
.mdc-layout-grid__cell--align-middle {
  align-self: center;
}
.mdc-layout-grid__cell--align-bottom {
  align-self: flex-end;
}
@supports (display: grid) {
  .mdc-layout-grid__cell--align-bottom {
    align-self: end;
  }
}

@media (min-width: 840px) {
  .mdc-layout-grid--fixed-column-width {
    width: 1176px;
    width: calc( var(--mdc-layout-grid-column-width-desktop, 72px) * 12 + var(--mdc-layout-grid-gutter-desktop, 24px) * 11 + var(--mdc-layout-grid-margin-desktop, 24px) * 2 );
  }
}
@media (min-width: 480px) and (max-width: 839px) {
  .mdc-layout-grid--fixed-column-width {
    width: 720px;
    width: calc( var(--mdc-layout-grid-column-width-tablet, 72px) * 8 + var(--mdc-layout-grid-gutter-tablet, 16px) * 7 + var(--mdc-layout-grid-margin-tablet, 16px) * 2 );
  }
}
@media (max-width: 479px) {
  .mdc-layout-grid--fixed-column-width {
    width: 368px;
    width: calc( var(--mdc-layout-grid-column-width-phone, 72px) * 4 + var(--mdc-layout-grid-gutter-phone, 16px) * 3 + var(--mdc-layout-grid-margin-phone, 16px) * 2 );
  }
}

.mdc-layout-grid--align-left {
  margin-right: auto;
  margin-left: 0;
}

.mdc-layout-grid--align-right {
  margin-right: 0;
  margin-left: auto;
}

`
