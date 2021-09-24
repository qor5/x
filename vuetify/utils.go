package vuetify

import (
	"strings"

	h "github.com/theplant/htmlgo"
)

func stringsTrim(vs ...string) (r []string) {
	for _, v := range vs {
		if cv := strings.TrimSpace(v); len(cv) > 0 {
			r = append(r, cv)
		}
	}
	return
}

func SetErrorMessages(t h.MutableAttrHTMLComponent, vs []string) {
	cvs := stringsTrim(vs...)
	if len(cvs) == 0 {
		return
	}
	t.SetAttr(":error-messages", h.JSONString(cvs))
}
