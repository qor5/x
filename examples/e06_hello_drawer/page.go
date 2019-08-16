package e06_hello_drawer

import (
	bo "github.com/sunfmin/bran/overlay"
	"github.com/sunfmin/bran/ui"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	InputName string
	NameError string
	Group     string
}

var globalName string

func HelloDrawer(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("form", form)
	ctx.Hub.RegisterEventFunc("close", close)
	ctx.Hub.RegisterEventFunc("update", update)
	ctx.Hub.RegisterEventFunc("updateForm", updateForm)

	var s = &mystate{}
	if ctx.Flash != nil {
		s = ctx.Flash.(*mystate)
	}

	pr.Schema = Div(
		H1(globalName),
		bo.Drawer(
			ui.LazyPortal().EventFunc("form", "param1").LoadWhenParentVisible(),
		).TriggerElement(
			A().Text("Edit").Href("#"),
		).Width(500),
		ui.Bind(Input("").Type("text").Value(s.Group)).FieldName("Group"),
		ui.Bind(Button("Check value")).OnClick("update"),
	)
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)
	ctx.Flash = s

	r.Reload = true
	return
}

func form(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	var s = &mystate{InputName: globalName}

	if ctx.Flash != nil {
		s = ctx.Flash.(*mystate)
	}

	r.Schema = Div(
		ui.Bind(Button("Close")).OnClick("close"),
		ui.Bind(Input("").Type("text").Value(s.InputName)).FieldName("InputName"),
		Label(s.NameError).Style("color:red"),
		ui.Bind(Button("Update")).OnClick("updateForm"),
	)
	return
}

func close(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	return
}

func updateForm(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)

	if len(s.InputName) < 10 {
		s.NameError = "is too short"
		ctx.Flash = s
		r, err = form(ctx)
	} else {
		globalName = s.InputName
		r.Reload = true
	}
	return
}
