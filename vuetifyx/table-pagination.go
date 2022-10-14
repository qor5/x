package vuetifyx

import (
	"context"
	"fmt"
	"sort"

	"github.com/goplaid/web"
	"github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

type VXTablePaginationBuilder struct {
	total           int64
	currPage        int64
	perPage         int64
	customPerPages  []int64
	noPerPagePart   bool
	onSelectPerPage interface{}
	onPrevPage      interface{}
	onNextPage      interface{}

	perPageText string
}

func VXTablePagination() *VXTablePaginationBuilder {
	return &VXTablePaginationBuilder{}
}

func (tpb *VXTablePaginationBuilder) PerPageText(v string) *VXTablePaginationBuilder {
	tpb.perPageText = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) Total(v int64) *VXTablePaginationBuilder {
	tpb.total = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) CurrPage(v int64) *VXTablePaginationBuilder {
	tpb.currPage = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) PerPage(v int64) *VXTablePaginationBuilder {
	tpb.perPage = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) CustomPerPages(v []int64) *VXTablePaginationBuilder {
	tpb.customPerPages = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) NoPerPagePart(v bool) *VXTablePaginationBuilder {
	tpb.noPerPagePart = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) OnSelectPerPage(v interface{}) *VXTablePaginationBuilder {
	tpb.onSelectPerPage = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) OnPrevPage(v interface{}) *VXTablePaginationBuilder {
	tpb.onPrevPage = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) OnNextPage(v interface{}) *VXTablePaginationBuilder {
	tpb.onNextPage = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	if tpb.onSelectPerPage == nil {
		tpb.OnSelectPerPage(web.Plaid().
			PushState(true).
			Query("per_page", web.Var("[$event]")).
			MergeQuery(true).
			Go())
	}
	if tpb.onPrevPage == nil {
		tpb.OnPrevPage(web.Plaid().
			PushState(true).
			Query("page", tpb.currPage-1).
			MergeQuery(true).
			Go())
	}
	if tpb.onNextPage == nil {
		tpb.OnNextPage(web.Plaid().
			PushState(true).
			Query("page", tpb.currPage+1).
			MergeQuery(true).
			Go())
	}

	var sItems []string
	{
		perPagesM := map[int64]struct{}{
			10:  {},
			15:  {},
			20:  {},
			50:  {},
			100: {},
		}
		if tpb.perPage > 0 {
			perPagesM[tpb.perPage] = struct{}{}
		}
		for _, v := range tpb.customPerPages {
			if v <= 0 {
				continue
			}
			perPagesM[v] = struct{}{}
		}
		perPages := make([]int, 0, len(perPagesM))
		for k, _ := range perPagesM {
			perPages = append(perPages, int(k))
		}
		sort.Ints(perPages)
		for _, v := range perPages {
			sItems = append(sItems, fmt.Sprint(v))
		}
	}

	currPageStart := (tpb.currPage-1)*tpb.perPage + 1
	currPageEnd := tpb.currPage * tpb.perPage
	if currPageEnd > tpb.total {
		currPageEnd = tpb.total
	}

	canNext := false
	canPrev := false
	if tpb.currPage*tpb.perPage < tpb.total {
		canNext = true
	}
	if tpb.currPage > 1 {
		canPrev = true
	}
	var nextIconStyle string
	var prevIconStyle string
	if canNext {
		nextIconStyle = "cursor: pointer;"
	}
	if canPrev {
		prevIconStyle = "cursor: pointer;"
	}

	rowsPerPageText := "Rows per page: "
	if tpb.perPageText != "" {
		rowsPerPageText = tpb.perPageText
	}
	return vuetify.VContainer(vuetify.VRow().Justify("end").Align("center").Class("ma-0").
		Children(
			h.If(!tpb.noPerPagePart,
				h.Div(
					h.Text(rowsPerPageText),
				),
				h.Div(
					vuetify.VSelect().Items(sItems).Value(fmt.Sprint(tpb.perPage)).
						Attach(false).
						Attr("@input", tpb.onSelectPerPage),
				).Style("width: 60px;").Class("ml-6"),
			),
			h.Div(
				h.Text(fmt.Sprintf("%d-%d of %d", currPageStart, currPageEnd, tpb.total)),
			).Class("ml-6"),
			h.Div(
				h.Span("").Style(prevIconStyle).Children(
					vuetify.VIcon("navigate_before").Size(32).Disabled(!canPrev).
						Attr("@click", tpb.onPrevPage),
				),
				h.Span("").Style(nextIconStyle).Children(
					vuetify.VIcon("navigate_next").Size(32).Disabled(!canNext).
						Attr("@click", tpb.onNextPage),
				).Class("ml-3"),
			).Class("ml-6"),
		)).Fluid(true).MarshalHTML(ctx)
}
