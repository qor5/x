package e08_hello_popover

import (
	"github.com/sunfmin/bran/web"
	bo "github.com/sunfmin/bran/x/overlay"
	. "github.com/theplant/htmlgo"
)

type popoverState struct {
	EditingName string
	NameError   string
}

type mystate struct {
	Popover1 *popoverState
	Popover2 *popoverState
	Popover3 *popoverState
}

var globalName string

var globalState = &mystate{
	Popover1: &popoverState{EditingName: "popover 1"},
	Popover2: &popoverState{EditingName: "popover 2"},
	Popover3: &popoverState{EditingName: "popover 3"},
}

func HelloPopover(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("remoteOverlay", remoteOverlay)
	ctx.Hub.RegisterEventFunc("update", update)

	pr.Schema = Div(
		H1(globalName),

		H2("Default"),
		bo.Popover(
			web.Portal().EventFunc("remoteOverlay").
				Visible("true").
				ParentForceUpdateAfterLoaded(),
		).TriggerElement(
			A().Text("Edit").Href("#"),
		).Placement("right"),

		H2("Load from remote"),
		bo.Popover(
			web.Portal().EventFunc("remoteOverlay").
				Visible("true").
				ParentForceUpdateAfterLoaded(),
		).TriggerElement(
			A().Text("Remote Loader").Href("#"),
		).Placement("right"),

		H2("Load when mouse over"),
		bo.Popover(
			web.Portal().EventFunc("remoteOverlay").
				Visible("true").
				ParentForceUpdateAfterLoaded(),
		).TriggerElement(
			A().Text("Mouseover").Href("#"),
		).Trigger("hover").Placement("bottom"),
	)
	return
}

func remoteOverlay(ctx *web.EventContext) (r web.EventResponse, err error) {
	ctx.MustUnmarshalForm(&globalState)

	r.Schema = overlay(globalState.Popover2, ctx)
	return
}

func overlay(s *popoverState, ctx *web.EventContext) HTMLComponent {
	return Div(
		web.Bind(Input("").Type("text").Value(s.EditingName)).FieldName("EditingName"),
		Label(s.NameError).Style("color:red"),
		web.Bind(Button("Update")).OnClick("update"),
	).Style("padding: 20px; background-color: white;")
}

func update(ctx *web.EventContext) (r web.EventResponse, err error) {
	var s = &popoverState{}
	ctx.MustUnmarshalForm(s)

	if len(s.EditingName) < 10 {
		s.NameError = "name is too short"
		r.Schema = overlay(s, ctx)
		return
	} else {
		globalName = s.EditingName
		s.NameError = ""
		r.Reload = true
	}
	return
}
