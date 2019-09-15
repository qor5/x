package e05_hello_customized_component

import (
	"context"
	"fmt"

	"github.com/goplaid/web"
	. "github.com/theplant/htmlgo"
)

type TagsInputBuilder struct {
	classNames   []string
	selectedKeys []string
	options      []*TagsInputOption
}

type TagsInputOption struct {
	Key   string
	Label HTMLComponent
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

func (b *TagsInputBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	//ui.Injector(ctx).PutScript(tagsInputScript)
	//ui.Injector(ctx).PutStyle(tagsInputStyles)

	selectedComps := []HTMLComponent{}
	optionComps := []HTMLComponent{}
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

//const tagsInputScript = `
//	(window.__branVueComponentRegisters = (window.__branVueComponentRegisters || [])).push(function(Vue){
//		Vue.component("tags-input", {
//			data: function() {
//				return {
//					isOpen: false
//				};
//			},
//			methods: {
//				toggle: function() {
//					this.isOpen = !this.isOpen;
//				}
//			},
//			template: "<div><slot v-bind:parent='this'/></div>"
//		});
//	})
//`
//
//const tagsInputStyles = `
//.tagsInput .tagsInputOptions {
//	display: none;
//}
//
//.tagsInput .tagsInputOptions.tagsInputOptionsOpen {
//	display: block;
//}
//`

// Above is component code

type mystate struct {
	Message string
}

func HelloCustomziedComponent(ctx *web.EventContext) (pr web.PageResponse, err error) {
	// s := ctx.StateOrInit(&mystate{}).(*mystate)

	opts := []*TagsInputOption{}
	for i := 1; i < 11; i++ {
		opts = append(opts, &TagsInputOption{
			Key:   fmt.Sprint(i),
			Label: Div().Text(fmt.Sprintf("label %d", i)),
		})
	}

	pr.Body = Div(
		TagsInput().Selected([]string{"1", "2", "3"}).Options(opts...),
		web.Bind(Button("Refresh")).OnClick("refresh"),
	)
	return
}

func reload(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.Reload = true
	return
}
