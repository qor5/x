package e04_hello_material_grid

import (
	"github.com/goplaid/web"
	. "github.com/goplaid/x/material"
	h "github.com/theplant/htmlgo"
)

func HelloGrid(ctx *web.EventContext) (pr web.PageResponse, err error) {

	pr.Schema = h.Div(
		h.Div().Text("Grid of default wide (4 columns) items").
			Class("presets-example-main-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("presets-example-main-cell"),
			Cell(h.Text("2")).Class("presets-example-main-cell"),
			Cell(h.Text("3")).Class("presets-example-main-cell"),
		).Class("presets-example-main-grid"),

		h.Div().Text("Grid of 1 column wide items").
			Class("presets-example-main-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
		).Class("presets-example-main-grid"),

		h.Div().Text("Grid of differently sized items").
			Class("presets-example-main-grid-legend"),
		Grid(
			Cell(h.Text("6")).Class("presets-example-main-cell").Span(6, ScreenAll),
			Cell(h.Text("4")).Class("presets-example-main-cell").Span(4, ScreenAll),
			Cell(h.Text("2")).Class("presets-example-main-cell").Span(2, ScreenAll),
		).Class("presets-example-main-grid").Margin(12, ScreenAll),

		h.Div().Text("Grid of items with tweaks at different screen sizes").
			Class("presets-example-main-grid-legend"),
		Grid(
			Cell(h.Text("6 (8 tablet)")).Class("presets-example-main-cell").
				Span(6, ScreenAll).Span(8, ScreenTablet),
			Cell(h.Text("4 (6 tablet)")).Class("presets-example-main-cell").
				Span(4, ScreenAll).Span(4, ScreenTablet),
			Cell(h.Text("2 (4 phone)")).Class("presets-example-main-cell").
				Span(2, ScreenAll).Span(4, ScreenPhone),
		).Class("presets-example-main-grid"),

		h.Div().Text("Grid nested within parent grid cell").
			Class("presets-example-main-grid-legend"),
		Grid(
			Cell(
				Grid(
					Cell(h.Tag("span").Text("Child 4")).Class("presets-example-main-child-cell", "presets-example-main-cell"),
					Cell(h.Tag("span").Text("Child 4")).Class("presets-example-main-child-cell", "presets-example-main-cell"),
					Cell(h.Tag("span").Text("Child 4")).Class("presets-example-main-child-cell", "presets-example-main-cell"),
				).Inner(),
				h.Tag("span").Text("Parent 4"),
			).Class("presets-example-main-parent-cell").
				Span(4, ScreenAll),
			Cell(h.Text("4")).Class("presets-example-main-cell").
				Span(4, ScreenAll),
			Cell(h.Text("4")).Class("presets-example-main-cell").
				Span(4, ScreenAll),
		).Class("presets-example-main-grid"),

		h.Div().Text("Grid with max width (1280px) and center alignment by default").
			Class("presets-example-main-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("presets-example-main-cell"),
			Cell(h.Text("2")).Class("presets-example-main-cell"),
			Cell(h.Text("3")).Class("presets-example-main-cell"),
		).Class("presets-example-main-grid", "max-width"),

		h.Div().Text("Grid with max width (1280px) and left alignment").
			Class("presets-example-main-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("presets-example-main-cell"),
			Cell(h.Text("2")).Class("presets-example-main-cell"),
			Cell(h.Text("3")).Class("presets-example-main-cell"),
		).Class("presets-example-main-grid", "max-width").AlignLeft(),

		h.Div().Text("Fixed column width layout grid and center alignment by default").
			Class("presets-example-main-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("2")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("3")).Class("presets-example-main-cell").Span(1, ScreenAll),
		).Class("presets-example-main-grid", "max-width").FixColumnWidth(),

		h.Div().Text("Fixed column width layout grid and right alignment").
			Class("presets-example-main-grid-legend"),
		Grid(
			Cell(h.Text("1")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("2")).Class("presets-example-main-cell").Span(1, ScreenAll),
			Cell(h.Text("3")).Class("presets-example-main-cell").Span(1, ScreenAll),
		).Class("presets-example-main-grid", "max-width").FixColumnWidth().AlignRight(),
	)

	styles(ctx)
	return
}

func reload(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.Reload = true
	return
}

func styles(ctx *web.EventContext) {
	ctx.Injector.HeadHTML(`
    <style>
      .presets-example-main-grid {
        background-color: #DDDDDD;
        margin-bottom: 32px;
      }

      .presets-example-main-cell {
        box-sizing: border-box;
        background-color: #666666;
        height: 200px;
        padding: 8px;
        color: white;

        font-size: 1.5em;
      }

      .presets-example-main-parent-cell {
        position: relative;
        background-color: #aaaaaa;
      }

      .presets-example-main-parent-cell>span{
        position: absolute;
        top: 8px;
        left: 8px;
        font-size: 1.5em;
        color: white;
      }

      .presets-example-main-child-cell {
        position: relative;
      }

      .presets-example-main-child-cell>span{
        position: absolute;
        bottom: 8px;
        right: 8px;
        color: #ddd;
      }

      .presets-example-main-grid.max-width {
        max-width: 1280px;
      }

      .presets-example-main-grid-legend {
        margin: 16px 0 8px 0;
      }
    </style>
	`)
}
