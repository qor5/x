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

func Demo(title string, path string) HTMLComponent {
	if len(title) == 0 {
		title = "Check the demo"
	}
	return Div(
		A().Text(title).Href(path).Target("_blank"),
	).Class("demo")
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
