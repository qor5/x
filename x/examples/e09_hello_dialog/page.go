package e09_hello_dialog

import (
	"github.com/goplaid/web"
	bo "github.com/goplaid/x/overlay"
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

func HelloDialog(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("remoteOverlay", remoteOverlay)
	ctx.Hub.RegisterEventFunc("update", update)

	var s = &mystate{}

	pr.Schema = Div(
		H1(globalState.Name),
		bo.Dialog(
			web.Portal().EventFunc("remoteOverlay").Visible("true"),
		).TriggerElement(
			A().Text("Edit").Href("#"),
		).DefaultVisible(s.dialogVisible),

		bo.Dialog(
			web.Portal().EventFunc("remoteOverlay").Visible("true"),
		).TriggerElement(
			A().Text("Remote Loader").Href("#"),
		).DefaultVisible(s.dialogVisibleRemote),

		bo.Dialog(
			web.Portal().EventFunc("remoteOverlay").Visible("true"),
		).TriggerElement(
			A().Text("Mouseover").Href("#"),
		).DefaultVisible(s.dialogVisible).
			Trigger("mouseover"),
	)
	return
}

func overlay(s *mystate, ctx *web.EventContext) HTMLComponent {
	return Div(
		web.Bind(Input("").Type("text").Value(s.EditingName)).FieldName("EditingName"),
		Label(s.NameError).Style("color:red"),
		web.Bind(Button("Update")).OnClick("update"),
	)
}

func remoteOverlay(ctx *web.EventContext) (r web.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)

	r.Schema = overlay(s, ctx)
	return
}

func update(ctx *web.EventContext) (r web.EventResponse, err error) {
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
