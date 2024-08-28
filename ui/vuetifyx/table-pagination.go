package vuetifyx

import (
	"context"
	"fmt"
	"math"
	"sort"

	"github.com/qor5/web/v3"
	v "github.com/qor5/x/v3/ui/vuetify"
	h "github.com/theplant/htmlgo"
)

type VXTablePaginationBuilder struct {
	total           int64
	currPage        int64
	perPage         int64
	customPerPages  []int64
	noPerPagePart   bool
	noOffsetPart    bool
	totalVisible    int64
	onSelectPerPage string
	onSelectPage    string

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

func (tpb *VXTablePaginationBuilder) NoOffsetPart(v bool) *VXTablePaginationBuilder {
	tpb.noOffsetPart = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) TotalVisible(v int64) *VXTablePaginationBuilder {
	tpb.totalVisible = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) OnSelectPerPage(v string) *VXTablePaginationBuilder {
	tpb.onSelectPerPage = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) OnSelectPage(v string) *VXTablePaginationBuilder {
	tpb.onSelectPage = v
	return tpb
}

func (tpb *VXTablePaginationBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
	if tpb.onSelectPerPage == "" {
		tpb.OnSelectPerPage(web.Plaid().
			PushState(true).
			Query("per_page", web.Var("[$event]")).
			MergeQuery(true).
			Go())
	}
	if tpb.onSelectPage == "" {
		tpb.OnSelectPage(web.Plaid().
			PushState(true).
			Query("page", web.Var("value")).
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
		for k := range perPagesM {
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

	totalPages := int64(math.Ceil(float64(tpb.total) / float64(tpb.perPage)))

	rowsPerPageText := "Rows per page: "
	if tpb.perPageText != "" {
		rowsPerPageText = tpb.perPageText
	}
	pagination := v.VPagination().ShowFirstLastPage(true).ActiveColor(v.ColorPrimary).Density(v.DensityCompact).
		Length(totalPages).
		// https://github.com/vuetifyjs/vuetify/issues/20321
		// https://github.com/vuetifyjs/vuetify/issues/18853
		Attr("v-on-mounted", `({el}) => { 
			const currentWidth = el.offsetWidth + 38; // 37.6;
			el.style.minWidth = currentWidth + "px";
		}`).
		Attr(":model-value", tpb.currPage).
		Attr("@update:model-value", fmt.Sprintf(`(value) => { %s }`, tpb.onSelectPage))
	if tpb.totalVisible > 0 {
		pagination = pagination.TotalVisible(tpb.totalVisible)
	}

	return h.Div(
		v.VRow().Justify("end").Align("center").Class("ma-0").
			Children(
				h.Iff(!tpb.noPerPagePart, func() h.HTMLComponent {
					return h.Components(
						h.Div(
							h.Text(rowsPerPageText),
						),
						h.Div(
							v.VSelect().Items(sItems).Variant("underlined").ModelValue(fmt.Sprint(tpb.perPage)).
								HideDetails(true).Density("compact").Attr("style", "margin-top: -8px").
								Attr("@update:model-value", tpb.onSelectPerPage),
						).Style("width: 64px;").Class("ml-6"),
					)
				}),
				h.Iff(!tpb.noOffsetPart, func() h.HTMLComponent {
					return h.Div(
						h.Text(fmt.Sprintf("%d-%d of %d", currPageStart, currPageEnd, tpb.total)),
					).Class("ml-6")
				}),
				pagination,
			)).MarshalHTML(ctx)
}
