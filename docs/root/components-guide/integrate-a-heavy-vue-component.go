package components_guide

import (
	"github.com/goplaid/x/docs/examples"
	"github.com/goplaid/x/docs/examples/e00_basics"
	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	"github.com/theplant/docgo/ch"
)

var IntegrateAHeavyVueComponent = Doc(
	Markdown(`
We can abstract any complicated of server side render component with [htmlgo](https://github.com/theplant/htmlgo).
But a lots of components in the modern web have done many things on the client side. means there are many logic
happens before the it interact with server side.

Here is an example, a rich text editor. you have a toolbar of buttons that you can interact, most of them won't
need to communicate with server. We are going to integrate the fantastic rich text editor [tiptap](https://tiptap.scrumpy.io/)
to be used as any ~htmlgo.HTMLComponent~.

**Step 1**: [Create a @vue/cli project](https://cli.vuejs.org/guide/creating-a-project.html):

~~~
$ vue create tiptapjs
~~~

Modify or add a separate ~vue.config.js~ config file, 

`),
	ch.Code(examples.TipTapVueConfig).Language("javascript"),

	Markdown(`
- Enable ~runtimeCompiler~ so that vue can parse template html generate from server.
- Made ~Vue~ as externals so that it won't be packed to the dist production js file, 
  Since we will be sharing one Vue.js for in one page with other libraries.
- Config svg module to inline the svg icons used by tiptap

**Step 2**: Create a vue component that use tiptap

Install ~tiptap~ and ~tiptap-extensions~ first
~~~
$ yarn add tiptap tiptap-extensions
~~~

And write the ~editor.vue~ something like this, We omitted the template at here. 

`),
	ch.Code(examples.TipTapEditorVueComponent).Language("javascript"),
	Markdown(`
We injected the ~this.$plaid()~. that is from ~web/corejs~, Which you will need to use
For every Go Plaid web applications. Here we uses one function ~fieldValue~ from it.
It set the form value when the rich text editor changes. So that later when you call
~EventFunc~ it the value will be posted to the server side. Here we will post the html value.
Also allow component user to set ~fieldName~, which is important when posting the value to the
server.

**Step 3**: At ~main.js~, Use a special hook to register the component to ~web/corejs~

`),
	ch.Code(examples.GoPlaidRegisterVueComponentSample).Language("go"),
	Markdown(`
**Step 4**: Test the component in a simple html

We edited the ~index.html~ inside public to be the following:

`),
	ch.Code(examples.TipTapDemoHTML).Language("html"),
	Markdown(`
- For ~http://localhost:3500/app.js~ to be able to serve. you have to run ~yarn serve~ in 
tiptapjs directory. 
- ~http://localhost:3100/app.js~ is goplaid web corejs vue project.
  So go to that directory and run ~yarn serve~ to start it. and then in
- Run a web server inside tiptapjs directory like ~python -m SimpleHTTPServer~ and point your
  Browser to the index.html file, and see if your vue component can render and behave correctly.

**Step 5**: Use [packr](https://github.com/gobuffalo/packr) to pack the dist folder

We write a packr box inside ~tiptapjs.go~ along side the tiptapjs folder.
`),
	ch.Code(examples.TipTapPackrSample).Language("go"),
	Markdown(`
And write a ~build.sh~ to build the javascript to production version, and run packr to pack
them into ~a_tiptap-packr.go~ file.
`),
	ch.Code(examples.TiptapBuilderSH).Language("bash"),

	Markdown(`
**Step 6**: Write a Go wrapper to wrap it to be a ~HTMLComponent~ 
`),
	ch.Code(examples.TipTapEditorHTMLComponent).Language("go"),

	Markdown(`
**Step 7**: Use it in your web app

To use it, first we have to mount the assets into our app
`),
	ch.Code(examples.TipTapComponentsPackSample).Language("go"),
	Markdown(`
And reference them in our layout function.
`),
	ch.Code(examples.TipTapLayoutSample).Language("go"),

	Markdown(`
And we write a page func to use it like any other component:
`),
	ch.Code(examples.HelloWorldTipTapSample).Language("go"),

	Markdown(`
And now let's check out our fruits:
`),
	utils.Demo("Integrate a Heavy Vue Component", e00_basics.HelloWorldTipTapPath, "e00_basics/use-tiptap-editor.go"),
).Title("Integrate a heavy Vue Component").
	Slug("components-guide/integrate-a-heavy-vue-component")
