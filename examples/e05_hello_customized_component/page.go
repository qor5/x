package e05_hello_customized_component

import (
	"fmt"

	. "github.com/sunfmin/bran/html"
	"github.com/sunfmin/pagui/ui"
)

type TagsInputBuilder struct {
	classNames   []string
	selectedKeys []string
	options      []*TagsInputOption
}

type TagsInputOption struct {
	Key   string
	Label ui.HTMLComponent
}

func TagsInput() (r *TagsInputBuilder) {
	r = &TagsInputBuilder{}
	return
}

func (b *TagsInputBuilder) Class(names ...string) (r *TagsInputBuilder) {
	b.classNames = names
	return b
}

func (b *TagsInputBuilder) Selected(keys []string) (r *TagsInputBuilder) {
	b.selectedKeys = keys
	return b
}

func (b *TagsInputBuilder) Options(options ...*TagsInputOption) (r *TagsInputBuilder) {
	b.options = options
	return b
}

func contains(k string, in []string) bool {
	for _, i := range in {
		if k == i {
			return true
		}
	}
	return false
}

func (b *TagsInputBuilder) MarshalHTML(ctx *ui.EventContext) (r []byte, err error) {
	ctx.Injector.PutScript(tagsInputScript)
	ctx.Injector.PutStyle(tagsInputStyles)

	selectedComps := []ui.HTMLComponent{}
	optionComps := []ui.HTMLComponent{}
	for _, op := range b.options {
		optionComps = append(optionComps, op.Label)
		if contains(op.Key, b.selectedKeys) {
			selectedComps = append(selectedComps, op.Label)
		}
	}

	root := Tag("tags-input").
		Class("tagsInput").
		Attr("v-slot", "{ parent }").
		Children(
			Div(
				Div().Class("tagsInputSelected").Children(
					selectedComps...,
				),
				Tag("button").Text("Toggle").Attr("v-on:click", "parent.toggle()"),
			),
			Div().
				Class("tagsInputOptions").
				Attr("v-bind:class", "{tagsInputOptionsOpen: parent.isOpen}").Children(
				optionComps...,
			),
		)
	return root.MarshalHTML(ctx)
}

const tagsInputScript = `
	(window.vueComps = (window.vueComps || [])).push(function(Vue){
		Vue.component("tags-input", {
			data: function() {
				return {
					isOpen: false
				};
			},
			methods: {
				toggle: function() {
					this.isOpen = !this.isOpen;
				}
			},
			template: "<div><slot v-bind:parent='this'/></div>"
		});
	})
`

const tagsInputStyles = `
.tagsInput .tagsInputOptions {
	display: none;
}

.tagsInput .tagsInputOptions.tagsInputOptionsOpen {
	display: block;
}
`

// Above is component code

type mystate struct {
	Message string
}

func HelloCustomziedComponent(ctx *ui.EventContext) (pr ui.PageResponse, err error) {
	// s := ctx.StateOrInit(&mystate{}).(*mystate)

	opts := []*TagsInputOption{}
	for i := 1; i < 11; i++ {
		opts = append(opts, &TagsInputOption{
			Key:   fmt.Sprint(i),
			Label: Div().Text(fmt.Sprintf("label %d", i)),
		})
	}

	pr.Schema = Div(
		TagsInput().Selected([]string{"1", "2", "3"}).Options(opts...),
		Button("Refresh").OnClick(ctx.Hub, "refresh", reload),
	)
	return
}

func reload(ctx *ui.EventContext) (r ui.EventResponse, err error) {
	r.Reload = true
	return
}
