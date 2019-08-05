package filterpanel

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/sunfmin/bran/ui"
	. "github.com/sunfmin/bran/vuetify"
	h "github.com/theplant/htmlgo"
)

type FilterBuilder struct {
	data FilterData
}

func Filter(data FilterData) (r *FilterBuilder) {
	r = &FilterBuilder{
		data: data,
	}

	return
}

func (b *FilterBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {

	var vals []int

	panels := VExpansionPanels().Multiple(true).Focusable(true)

	for i, fi := range b.data {
		switch fi.ItemType {
		case ItemTypeSelect:

		}

		if fi.Selected {
			vals = append(vals, i)
		}

		panels.AppendChildren(
			VExpansionPanel(
				VExpansionPanelHeader(
					h.Text(fi.Label)),
				VExpansionPanelContent(
					//VSelect().Items(nil),
					h.Text("Hello"),
				),
			),
		)
	}

	panels.Value(vals)

	return VMenu(
		ui.Slot(
			VBtn("Filter").Attr("v-on", "on"),
		).Name("activator").Scope("{ on }"),
		VCard(
			panels,
		).Flat(true).Width(440),
	).OffsetY(true).CloseOnContentClick(false).MarshalHTML(ctx)

}

type FilterItemType string

const (
	ItemTypeDate   FilterItemType = "DateItem"
	ItemTypeSelect FilterItemType = "SelectItem"
	ItemTypeNumber FilterItemType = "NumberItem"
	ItemTypeString FilterItemType = "StringItem"
)

type FilterItemModifier string

const (
	ModifierInTheLast    FilterItemModifier = "inTheLast"    // Date
	ModifierEquals       FilterItemModifier = "equals"       // Date, String, Number
	ModifierBetween      FilterItemModifier = "between"      // Date, Number
	ModifierIsAfter      FilterItemModifier = "isAfter"      // Date
	ModifierIsAfterOrOn  FilterItemModifier = "isAfterOrOn"  // Date
	ModifierIsBefore     FilterItemModifier = "isBefore"     // Date
	ModifierIsBeforeOrOn FilterItemModifier = "isBeforeOrOn" // Date
	ModifierGreaterThan  FilterItemModifier = "greaterThan"  // Number
	ModifierLessThan     FilterItemModifier = "lessThan"     // Number
	ModifierContains     FilterItemModifier = "contains"     // String
)

type FilterItemInTheLastUnit string

const (
	InTheLastUnitDays   FilterItemInTheLastUnit = "days"
	InTheLastUnitMonths FilterItemInTheLastUnit = "months"
)

type FilterItemTimezone string

const (
	TimezoneLocal FilterItemTimezone = "local"
	TimezoneUTC   FilterItemTimezone = "utc"
)

type FilterData []*FilterItem

type SelectItem struct {
	Label string `json:"label,omitempty"`
	Key   string `json:"key,omitempty"`
}

type FilterItem struct {
	Key            string                  `json:"key,omitempty"`
	Label          string                  `json:"label,omitempty"`
	ItemType       FilterItemType          `json:"itemType,omitempty"`
	Selected       bool                    `json:"selected,omitempty"`
	Modifier       FilterItemModifier      `json:"modifier,omitempty"`
	Value          string                  `json:"value,omitempty"`
	ValueFrom      string                  `json:"valueFrom,omitempty"`
	ValueTo        string                  `json:"valueTo,omitempty"`
	InTheLastValue string                  `json:"inTheLastValue,omitempty"`
	InTheLastUnit  FilterItemInTheLastUnit `json:"inTheLastUnit,omitempty"`
	Timezone       FilterItemTimezone      `json:"timezone,omitempty"`
	SQLCondition   string                  `json:"-"`
	Options        []*SelectItem           `json:"options,omitempty"`
}

func (fd FilterData) getSQLCondition(key string) string {
	for _, it := range fd {
		if it.Key == key {
			return it.SQLCondition
		}
	}
	return ""
}

var sqlOps = map[string]string{
	"":      "=",
	"gte":   ">=",
	"lte":   "<=",
	"gt":    ">",
	"lt":    "<",
	"ilike": "ILIKE",
}

func (fd FilterData) SetByQueryString(qs string) (sqlCondition string, sqlArgs []interface{}) {
	m, _ := url.ParseQuery(qs)

	conds := []string{}

	var keyModValueMap = map[string]map[string]string{}
	for k, v := range m {
		segs := strings.Split(k, ".")
		var mod = ""
		key := k
		if len(segs) > 1 {
			key = segs[0]
			mod = segs[1]
		}

		if _, ok := keyModValueMap[key]; !ok {
			keyModValueMap[key] = map[string]string{}
		}

		keyModValueMap[key][mod] = v[0]

		sqlc := fd.getSQLCondition(key)
		if len(sqlc) > 0 {
			conds = append(conds, fmt.Sprintf(sqlc, sqlOps[mod]))
			if mod == "ilike" {
				sqlArgs = append(sqlArgs, fmt.Sprintf("%%%s%%", v[0]))
			} else {
				sqlArgs = append(sqlArgs, v[0])
			}
		}
	}

	sqlCondition = strings.Join(conds, " AND ")

	for k, mv := range keyModValueMap {
		for _, it := range fd {
			if it.Key != k {
				continue
			}

			if len(mv) == 2 {
				it.Selected = true
				it.Modifier = ModifierBetween
				if it.ItemType == ItemTypeDate {
					it.ValueFrom = unixToDate(mv["gte"], it.Timezone == TimezoneUTC, 0)
					it.ValueTo = unixToDate(mv["lt"], it.Timezone == TimezoneUTC, -1)
				}

				if it.ItemType == ItemTypeNumber {
					it.ValueFrom = mv["gte"]
					it.ValueTo = mv["lte"]
				}

			} else if len(mv) == 1 {
				it.Selected = true
				for mod, v := range mv {
					if mod == "" {
						it.Modifier = ModifierEquals
					}

					if it.ItemType == ItemTypeDate {
						it.Value = unixToDate(v, it.Timezone == TimezoneUTC, 0)
						if mod == "gte" {
							it.Modifier = ModifierIsAfterOrOn
							it.Value = unixToDate(v, it.Timezone == TimezoneUTC, 0)
						}
						if mod == "gt" {
							it.Modifier = ModifierIsAfter
							it.Value = unixToDate(v, it.Timezone == TimezoneUTC, -1)
						}
						if mod == "lt" {
							it.Modifier = ModifierIsBefore
							it.Value = unixToDate(v, it.Timezone == TimezoneUTC, 0)
						}
						if mod == "lte" {
							it.Modifier = ModifierIsBeforeOrOn
							it.Value = unixToDate(v, it.Timezone == TimezoneUTC, -1)
						}
						continue
					}

					it.Value = v
					if it.ItemType == ItemTypeNumber {
						if mod == "gte" {
							it.Modifier = ModifierBetween
							it.ValueFrom = v
						}
						if mod == "lte" {
							it.Modifier = ModifierBetween
							it.ValueTo = v
						}
						if mod == "gt" {
							it.Modifier = ModifierGreaterThan
						}
						if mod == "lt" {
							it.Modifier = ModifierLessThan
						}
						continue
					}

					if it.ItemType == ItemTypeString {
						if mod == "ilike" {
							it.Modifier = ModifierContains
						}
						continue
					}

				}
			}
		}
	}

	return
}

func unixToDate(u string, utc bool, sub int) string {
	unix, _ := strconv.ParseInt(u, 10, 64)
	d := time.Unix(unix, 0).Add(time.Duration(24 * sub * int(time.Hour)))
	if utc {
		d = d.UTC()
	}
	return d.Format("01/02/2006")
}
