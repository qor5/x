package e06_hello_drawer

import (
	"github.com/sunfmin/bran/web"
	bo "github.com/sunfmin/bran/x/overlay"
	. "github.com/theplant/htmlgo"
)

type mystate struct {
	InputName string
	NameError string
	Group     string
}

var globalName string

func HelloDrawer(ctx *web.EventContext) (pr web.PageResponse, err error) {
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
			web.Portal().EventFunc("form", "param1").LoadWhenParentVisible(),
		).TriggerElement(
			A().Text("Edit").Href("#"),
		).Width(500),
		web.Bind(Input("").Type("text").Value(s.Group)).FieldName("Group"),
		web.Bind(Button("Check value")).OnClick("update"),
	)
	return
}

func update(ctx *web.EventContext) (r web.EventResponse, err error) {
	var s = &mystate{}
	ctx.MustUnmarshalForm(s)
	ctx.Flash = s

	r.Reload = true
	return
}

func form(ctx *web.EventContext) (r web.EventResponse, err error) {
	var s = &mystate{InputName: globalName}

	if ctx.Flash != nil {
		s = ctx.Flash.(*mystate)
	}

	r.Schema = Div(
		web.Bind(Button("Close")).OnClick("close"),
		web.Bind(Input("").Type("text").Value(s.InputName)).FieldName("InputName"),
		Label(s.NameError).Style("color:red"),
		web.Bind(Button("Update")).OnClick("updateForm"),
	)
	return
}

func close(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.Reload = true
	return
}

func updateForm(ctx *web.EventContext) (r web.EventResponse, err error) {
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
