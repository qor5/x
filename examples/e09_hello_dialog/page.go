package e09_hello_dialog

import (
	bo "github.com/sunfmin/bran/overlay"
	"github.com/sunfmin/bran/ui"
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

func HelloDialog(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("remoteOverlay", remoteOverlay)
	ctx.Hub.RegisterEventFunc("update", update)

	var s = &mystate{}

	pr.Schema = Div(
		H1(globalState.Name),
		bo.Dialog(
			ui.Portal().EventFunc("remoteOverlay").Visible("true"),
		).TriggerElement(
			A().Text("Edit").Href("#"),
		).DefaultVisible(s.dialogVisible),

		bo.Dialog(
			ui.Portal().EventFunc("remoteOverlay").Visible("true"),
		).TriggerElement(
			A().Text("Remote Loader").Href("#"),
		).DefaultVisible(s.dialogVisibleRemote),

		bo.Dialog(
			ui.Portal().EventFunc("remoteOverlay").Visible("true"),
		).TriggerElement(
			A().Text("Mouseover").Href("#"),
		).DefaultVisible(s.dialogVisible).
			Trigger("mouseover"),
	)
	return
}

func overlay(s *mystate, ctx *ui.EventContext) HTMLComponent {
	return Div(
		ui.Bind(Input("").Type("text").Value(s.EditingName)).FieldName("EditingName"),
		Label(s.NameError).Style("color:red"),
		ui.Bind(Button("Update")).OnClick("update"),
	)
}

func remoteOverlay(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)

	r.Schema = overlay(s, ctx)
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)

	if len(s.EditingName) < 10 {
		s.NameError = "name is too short"
		r.Schema = overlay(s, ctx)
	} else {
		globalState.Name = s.EditingName
		s.NameError = ""
		r.Reload = true
	}

	return
}
