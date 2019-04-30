package e04_hello_material_grid

import (
	h "github.com/sunfmin/bran/html"
	. "github.com/sunfmin/bran/material"
	"github.com/sunfmin/pagui/ui"
)

func HelloGrid(ctx *ui.EventContext) (pr ui.PageResponse, err error) {

	pr.Schema = h.Div(
		h.Div().Text("Grid of default wide (4 columns) items").
			Class("demo-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("demo-cell"),
			Cell(h.Text("2")).Class("demo-cell"),
			Cell(h.Text("3")).Class("demo-cell"),
		).Class("demo-grid"),

		h.Div().Text("Grid of 1 column wide items").
			Class("demo-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
		).Class("demo-grid"),

		h.Div().Text("Grid of differently sized items").
			Class("demo-grid-legend"),
		Grid(
			Cell(h.Text("6")).Class("demo-cell").Span(6, ScreenAll),
			Cell(h.Text("4")).Class("demo-cell").Span(4, ScreenAll),
			Cell(h.Text("2")).Class("demo-cell").Span(2, ScreenAll),
		).Class("demo-grid").Margin(12, ScreenAll),

		h.Div().Text("Grid of items with tweaks at different screen sizes").
			Class("demo-grid-legend"),
		Grid(
			Cell(h.Text("6 (8 tablet)")).Class("demo-cell").
				Span(6, ScreenAll).Span(8, ScreenTablet),
			Cell(h.Text("4 (6 tablet)")).Class("demo-cell").
				Span(4, ScreenAll).Span(4, ScreenTablet),
			Cell(h.Text("2 (4 phone)")).Class("demo-cell").
				Span(2, ScreenAll).Span(4, ScreenPhone),
		).Class("demo-grid"),

		h.Div().Text("Grid nested within parent grid cell").
			Class("demo-grid-legend"),
		Grid(
			Cell(
				Grid(
					Cell(h.Tag("span").Text("Child 4")).Class("demo-child-cell", "demo-cell"),
					Cell(h.Tag("span").Text("Child 4")).Class("demo-child-cell", "demo-cell"),
					Cell(h.Tag("span").Text("Child 4")).Class("demo-child-cell", "demo-cell"),
				).Inner(),
				h.Tag("span").Text("Parent 4"),
			).Class("demo-parent-cell").
				Span(4, ScreenAll),
			Cell(h.Text("4")).Class("demo-cell").
				Span(4, ScreenAll),
			Cell(h.Text("4")).Class("demo-cell").
				Span(4, ScreenAll),
		).Class("demo-grid"),

		h.Div().Text("Grid with max width (1280px) and center alignment by default").
			Class("demo-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("demo-cell"),
			Cell(h.Text("2")).Class("demo-cell"),
			Cell(h.Text("3")).Class("demo-cell"),
		).Class("demo-grid", "max-width"),

		h.Div().Text("Grid with max width (1280px) and left alignment").
			Class("demo-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("demo-cell"),
			Cell(h.Text("2")).Class("demo-cell"),
			Cell(h.Text("3")).Class("demo-cell"),
		).Class("demo-grid", "max-width").AlignLeft(),

		h.Div().Text("Fixed column width layout grid and center alignment by default").
			Class("demo-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("2")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("3")).Class("demo-cell").Span(1, ScreenAll),
		).Class("demo-grid", "max-width").FixColumnWidth(),

		h.Div().Text("Fixed column width layout grid and right alignment").
			Class("demo-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("2")).Class("demo-cell").Span(1, ScreenAll),
			Cell(h.Text("3")).Class("demo-cell").Span(1, ScreenAll),
		).Class("demo-grid", "max-width").FixColumnWidth().AlignRight(),
	)

	styles(ctx)
	return
}

func reload(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	return
}

func styles(ctx *ui.EventContext) {
	ctx.Injector.PutStyle(`
      .demo-grid {
        background-color: #DDDDDD;
        margin-bottom: 32px;
      }

      .demo-cell {
        box-sizing: border-box;
        background-color: #666666;
        height: 200px;
        padding: 8px;
        color: white;

        font-size: 1.5em;
      }

      .demo-parent-cell {
        position: relative;
        background-color: #aaaaaa;
      }

      .demo-parent-cell>span{
        position: absolute;
        top: 8px;
        left: 8px;
        font-size: 1.5em;
        color: white;
      }

      .demo-child-cell {
        position: relative;
      }

      .demo-child-cell>span{
        position: absolute;
        bottom: 8px;
        right: 8px;
        color: #ddd;
      }

      .demo-grid.max-width {
        max-width: 1280px;
	  }

      .demo-grid-legend {
        margin: 16px 0 8px 0;
      }
	`)
}
