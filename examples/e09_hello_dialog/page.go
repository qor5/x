package e09_hello_dialog

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sunfmin/bran/ui"
	bo "github.com/sunfmin/branoverlay"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	dialogVisible       bool
	dialogVisibleRemote bool
	EditingName         string
	NameError           string
}

var globalState = &struct {
	Name string
}{}

func randStr(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s: %d", prefix, rand.Int31n(100))
}

func overlay(s *mystate, ctx *ui.EventContext) HTMLComponent {
	return Div(
		ui.Bind(Input("").Type("text").Value(s.EditingName)).FieldName("EditingName"),
		Label(s.NameError).Style("color:red"),
		ui.Bind(Button("Update")).OnClick(ctx.Hub, "update", update),
	).Style("padding: 20px; background-color: white;")
}

func HelloDialog(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	s := ctx.StateOrInit(&mystate{EditingName: globalState.Name}).(*mystate)

	pr.Schema = Div(
		H1(globalState.Name),
		bo.Dialog(
			overlay(s, ctx),
		).Trigger(
			A().Text("Edit").Href("#"),
		).DefaultVisible(s.dialogVisible),

		bo.Dialog(
			ui.LazyLoader(ctx.Hub, "remoteOverlay", remoteOverlay).Visible("true"),
		).Trigger(
			A().Text("Remote Loader").Href("#"),
		).DefaultVisible(s.dialogVisibleRemote),
	)
	return
}

func remoteOverlay(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.State.(*mystate)
	r.Schema = overlay(s, ctx)
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	s := ctx.State.(*mystate)
	if len(s.EditingName) < 10 {
		s.NameError = "name is too short"
		s.dialogVisible = true
	} else {
		globalState.Name = s.EditingName
		s.NameError = ""
		s.dialogVisible = false
	}

	return
}
