package vuetifyx

import (
	h "github.com/theplant/htmlgo"
)

// SafeText renders s as plain text that Vue will not compile.
//
// QOR5 pages are compiled as Vue templates in the browser, and h.Text only
// escapes HTML special characters, not "{{ }}". Rendering user-controlled data
// with h.Text therefore lets "{{ ... }}" run as JavaScript in the viewer's
// session (client-side template injection). Use SafeText for any value that
// may come from users, customers or other tenants.
func SafeText(s string) h.HTMLComponent {
	return h.Span(s).Attr("v-pre", true)
}
