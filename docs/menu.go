package docs

import (
	"net/http"

	"github.com/goplaid/x/docs/root"
	advanced_functions "github.com/goplaid/x/docs/root/advanced-functions"
	"github.com/goplaid/x/docs/root/basics"
	digging_deeper "github.com/goplaid/x/docs/root/digging-deeper"
	getting_started "github.com/goplaid/x/docs/root/getting-started"
	"github.com/goplaid/x/docs/utils"
	"github.com/theplant/docgo"
)

func DocMenu(prefix string) http.Handler {
	return docgo.New().
		SitePrefix(prefix).
		Assets("/assets/", root.Assets).
		MainPageTitle("GoPlaid Document").
		DocTree(
			root.Home,
			&docgo.DocsGroup{
				Title: "Getting Started",
				Docs: []*docgo.DocBuilder{
					getting_started.OneMinuteQuickStart,
				},
			},
			&docgo.DocsGroup{
				Title: "Basics",
				Docs: []*docgo.DocBuilder{
					basics.FormHandling,
					basics.BasicInputs,
					basics.ListingCustomizations,
					basics.EditingCustomizations,
					basics.Filter,
					basics.AutoComplete,
					basics.ShortCut,
					basics.VariantSubForm,
					basics.NavigationDrawer,
					basics.LinkageSelect,
					basics.Permissions,
					basics.NotificationCenter,
				},
			},

			&docgo.DocsGroup{
				Title: "Advanced Functions",
				Docs: []*docgo.DocBuilder{
					advanced_functions.PageFuncAndEventFunc,
					advanced_functions.TheGoHTMLBuilder,
					advanced_functions.ATasteOfUsingVuetifyInGo,
					advanced_functions.ItsTheWholeHouse,

					advanced_functions.LazyPortalsAndReload,
					advanced_functions.LayoutFunctionAndPageInjector,
					advanced_functions.SwitchPagesWithPushState,
					advanced_functions.ReloadPageWithAFlash,
					advanced_functions.PartialRefreshWithPortal,
					advanced_functions.ManipulatePageURLInEventFunc,
					advanced_functions.SummaryOfEventResponse,
					advanced_functions.WebScope,
					advanced_functions.EventHandling,
					advanced_functions.DetailPageForComplexObject,
				},
			},
			&docgo.DocsGroup{
				Title: "Digging Deeper",
				Docs: []*docgo.DocBuilder{
					digging_deeper.CompositeNewComponentWithGo,
					digging_deeper.IntegrateAHeavyVueComponent,
				},
			},
			&docgo.DocsGroup{
				Title: "Appendix",
				Docs: []*docgo.DocBuilder{
					docgo.Doc(utils.ExamplesDoc()).
						Title("All Demo Examples").
						Slug("appendix/all-demo-examples"),
				},
			},
		).
		Build()
}
