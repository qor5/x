package html_test

import (
	"testing"

	. "github.com/sunfmin/bran/html"
	"github.com/sunfmin/bran/ui"
	"github.com/theplant/testingutils"
)

var htmltagCases = []struct {
	name     string
	tag      *HTMLTagBuilder
	expected string
}{
	{
		name: "case 1",
		tag: Div(
			Div().Text("Hello"),
		),
		expected: `<div>
<div>Hello</div>
</div>
`,
	},
}

func TestHtmlTag(t *testing.T) {
	for _, c := range htmltagCases {
		r, err := c.tag.MarshalHTML(new(ui.EventContext))
		if err != nil {
			panic(err)
		}
		diff := testingutils.PrettyJsonDiff(c.expected, string(r))
		if len(diff) > 0 {
			t.Error(c.name, diff)
		}
	}
}
