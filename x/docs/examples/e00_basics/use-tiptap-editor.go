package e00_basics

//@snippet_begin(HelloWorldTipTapSample)
import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/tiptap"
	. "github.com/theplant/htmlgo"
	"github.com/yosssi/gohtml"
)

func HelloWorldTipTap(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("refresh", refresh)

	defaultValue := ctx.R.FormValue("Content1")
	if len(defaultValue) == 0 {
		defaultValue = `
			<h1>Hello</h1>
			<p>
				This is a nice editor
			</p>
			<ul>
			  <li>
				<p>
				  123
				</p>
			  </li>
			  <li>
				<p>
				  456
				</p>
			  </li>
			  <li>
				<p>
				  789
				</p>
			  </li>
			</ul>
`
	}

	pr.Body = Div(
		tiptap.TipTapEditor().
			FieldName("Content1").
			Value(defaultValue),
		Hr(),
		Pre(
			gohtml.Format(ctx.R.FormValue("Content1")),
		).Style("background-color: #f8f8f8; padding: 20px;"),
		web.Bind(
			Button("Submit").Style("font-size: 24px"),
		).OnClick("refresh"),
	)

	return
}

func refresh(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.Reload = true
	return
}

//@snippet_end

const HelloWorldTipTapPath = "/samples/hello_world_tiptap"
