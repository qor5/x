package root

import (
	"embed"

	"github.com/goplaid/x/docs/utils"
	. "github.com/theplant/docgo"
	. "github.com/theplant/htmlgo"
)

var Home = Doc(
	Markdown(`
GoPlaid is yet another Go library to build web applications. We aim to accelerate the development speed and make the website highly customizable.

- It prefers writing HTML in [static typing Go language](/advanced-functions/the-go-html-builder.html), rather than a certain type of template language, Not even go template.
- It try to minify the needs to write any JavaScript/Typescript for building interactive web applications
- It maximize the reusability of Components. since it uses Go to write components, You can abstract component very easy, and use component from a third party Go package is also like using normal Go packages.
	`),

	utils.Anchor(H2(""), "How is this document organized"),
	Markdown(`
Most of latter examples are based on the initial sample project. In another word, we will demonstrate how to build a rich functioned website by this document.

- First, we will start with a quick sample project that would give you a rough but visual idea of what GoPlaid can do.
- Second, we will introduce the basic functions, The sequence is from listing page to editing page. You can find all commonly used Admin website features in this section.
- Third, we will introduce the essentials of GoPlaid and advanced functions, You would understand how GoPlaid render a page and advanced features like "how to partially refresh a page".
- At last, the digging deeper part, you would learn how to create new component for GoPlaid

**Join the Discord community**: https://discord.gg/76YPsVBE4E
`)).Title("Introduction").
	Slug("/")

//go:embed assets/**.*
var Assets embed.FS
