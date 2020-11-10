package utils

import (
	"encoding/json"
	"fmt"

	"github.com/goplaid/web"
	"github.com/shurcooL/sanitized_anchor_name"
	. "github.com/theplant/htmlgo"
)

func Anchor(h *HTMLTagBuilder, text string) HTMLComponent {
	anchorName := sanitized_anchor_name.Create(text)
	return h.Children(
		Text(text),
		A().Class("anchor").Href(fmt.Sprintf("#%s", anchorName)),
	).Id(anchorName)
}

type Example struct {
	Title      string
	DemoPath   string
	SourcePath string
}

var LiveExamples []*Example

func Demo(title string, demoPath string, sourcePath string) HTMLComponent {
	ex := &Example{
		Title:      title,
		DemoPath:   demoPath,
		SourcePath: fmt.Sprintf("https://github.com/goplaid/x/tree/master/docs/examples/%s", sourcePath),
	}

	LiveExamples = append(LiveExamples, ex)

	return Div(
		Div(
			A().Text("Check the demo").Href(ex.DemoPath).Target("_blank"),
			Text(" | "),
			A().Text("Source on GitHub").
				Href(ex.SourcePath).
				Target("_blank"),
		).Class("demo"),
	)
}

func ExamplesDoc() HTMLComponent {
	u := Ul()
	for _, le := range LiveExamples {
		u.AppendChildren(
			Li(
				A().Href(le.DemoPath).Text(le.Title).Target("_blank"),
				Text(" | "),
				A().Href(le.SourcePath).Text("Source").Target("_blank"),
			),
		)
	}
	return u
}

func PrettyFormAsJSON(ctx *web.EventContext) HTMLComponent {
	if ctx.R.Form == nil {
		return nil
	}

	formData, err := json.MarshalIndent(ctx.R.Form, "", "\t")
	if err != nil {
		panic(err)
	}

	return Pre(
		string(formData),
	)
}
