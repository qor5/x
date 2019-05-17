package e08_hello_popover

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sunfmin/bran/ui"
	bo "github.com/sunfmin/branoverlay"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	popoverVisible       bool
	popoverVisibleRemote bool
	Name                 string
	NameError            string
}

func randStr(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s: %d", prefix, rand.Int31n(100))
}

func overlay(s *mystate, ctx *ui.EventContext) HTMLComponent {
	return Div(
		ui.Bind(Input("").Type("text").Value(s.Name)).FieldName("Name"),
		Label(s.NameError).Style("color:red"),
		ui.Bind(Button("Update")).OnClick(ctx.Hub, "update", update),
	)
}

func HelloPopover(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{}).(*mystate)

	pr.Schema = Div(
		H1(s.Name),
		bo.Popover(
			A().Text("Edit").Href("#"),
		).Overlay(overlay(s, ctx)).DefaultVisible(s.popoverVisible),

		bo.Popover(
			A().Text("Remote Loader").Href("#"),
		).Overlay(
			ui.LazyLoader(ctx.Hub, "remoteOverlay", remoteOverlay).Visible("true"),
		).DefaultVisible(s.popoverVisibleRemote),
	)
	return
}

func remoteOverlay(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.StateOrInit(&mystate{}).(*mystate)
	r.Schema = overlay(s, ctx)
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	s := ctx.State.(*mystate)
	if len(s.Name) < 10 {
		s.NameError = "name is too short"
		s.popoverVisible = true
		s.Name = ""
	} else {
		s.NameError = ""
		s.popoverVisible = false
	}
	return
}
