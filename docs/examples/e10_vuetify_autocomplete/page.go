package e10_vuetify_autocomplete

// @snippet_begin(VuetifyAutoCompleteSample)

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/goplaid/web"
	. "github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

type myFormValue struct {
	Values1 []string
	Values2 []string
	Value3  string
}

type User struct {
	Login string
	Name  string
}

var selectedItems1 = []*User{
	{Login: "sam", Name: "Sam"},
	{Login: "charles", Name: "Charles"},
}

var options1 = []*User{
	{Login: "sam", Name: "Sam"},
	{Login: "john", Name: "John"},
	{Login: "charles", Name: "Charles"},
}

var selectedItems2 = []*User{
	{Login: "charles", Name: "Charles"},
}

var selectedItems3 = []*User{
	{Login: "charles", Name: "Charles"},
}

var options2 = []*User{
	{Login: "sam", Name: "Sam"},
	{Login: "john", Name: "John"},
	{Login: "charles", Name: "Charles"},
}

var globalState = &myFormValue{
	Values1: []string{
		"sam",
		"charles",
	},
	Values2: []string{
		"charles",
	},
	Value3: "charles",
}

func VuetifyAutocomplete(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("update", update)

	result := h.Ul()
	for _, v := range globalState.Values1 {
		result.AppendChildren(h.Li().Text(v))
	}
	result.AppendChildren(h.Li().Text("======"))
	for _, v := range globalState.Values2 {
		result.AppendChildren(h.Li().Text(v))
	}

	pr.Body = VContainer(
		h.H1("VAutocomplete"),
		VAutocomplete().
			Items(options1).
			FieldName("Values1").
			ItemText("Name").
			ItemValue("Login").
			Label("Static Options").
			Value(globalState.Values1),

		VAutocomplete().
			ItemsEventFunc(ctx.Hub, "users", users).
			ItemText("Name").
			ItemValue("Login").
			SelectedItems(selectedItems2).
			FieldName("Values2").
			Label("Load Options from Remote").
			Value(globalState.Values2),

		result,
		h.H1("VSelect"),
		VSelect().
			Items(options1).
			ItemText("Name").
			ItemValue("Login").
			FieldName("Value3").
			Solo(true).
			Value(globalState.Value3),
		h.Pre(globalState.Value3),
		VBtn("Update").
			Color("success").
			OnClick("update"),
	)
	return
}

func users(ctx *web.EventContext) (r web.EventResponse, err error) {
	us := []*User{}
	for _, u := range options1 {
		us = append(us, u)
	}
	if len(options2) <= 100 {
		for i := 0; i < 200; i++ {
			us = append(us, &User{
				Login: randomdata.Email(),
				Name:  randomdata.SillyName(),
			})
		}
		options2 = us
	}

	r.Data = options2
	return
}

func update(ctx *web.EventContext) (r web.EventResponse, err error) {
	globalState = &myFormValue{}
	ctx.MustUnmarshalForm(globalState)

	selectedItems1 = []*User{}
	for _, login := range globalState.Values1 {
		for _, u := range options1 {
			if u.Login == login {
				selectedItems1 = append(selectedItems1, u)
			}
		}
	}

	selectedItems2 = []*User{}
	for _, login := range globalState.Values2 {
		for _, u := range options2 {
			if u.Login == login {
				selectedItems2 = append(selectedItems2, u)
			}
		}
	}

	selectedItems3 = []*User{}
	for _, u := range options1 {
		if u.Login == globalState.Value3 {
			selectedItems3 = append(selectedItems3, u)
		}
	}
	r.Reload = true

	return
}

// @snippet_end

const VuetifyAutoCompletePath = "/samples/vuetify-auto-complete"
