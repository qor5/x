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

var name string

func HelloDrawer(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("form", form)
	ctx.Hub.RegisterEventFunc("close", close)
	ctx.Hub.RegisterEventFunc("update", update)
	ctx.Hub.RegisterEventFunc("updateForm", updateForm)

	s := ctx.StateOrInit(&mystate{}).(*mystate)

	pr.Schema = Div(
		H1(name),
		bo.Drawer(
			ui.LazyLoader("form", "param1").LoadWhenParentVisible(),
		).TriggerElement(
			A().Text("Edit").Href("#"),
		).Width(500),
		ui.Bind(Input("").Type("text").Value(s.Group)).FieldName("Group"),
		ui.Bind(Button("Check value")).OnClick("update"),
	)
	return
}

func update(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	return
}

func form(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	s := ctx.State.(*mystate)
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
	s := ctx.State.(*mystate)
	if len(s.InputName) < 10 {
		s.NameError = "name is too short"
		r, err = form(ctx)
	} else {
		name = s.InputName
		s.NameError = ""
		s.InputName = ""
		r.Reload = true
	}
	return
}
