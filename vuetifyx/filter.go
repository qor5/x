package vuetifyx

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/goplaid/web"
	h "github.com/theplant/htmlgo"
)

type VXFilterBuilder struct {
	value FilterData
	tag   *h.HTMLTagBuilder
}

func VXFilter(value FilterData) (r *VXFilterBuilder) {
	r = &VXFilterBuilder{
		value: value,
		tag:   h.Tag("vx-filter"),
	}

	var visibleFilterData FilterData
	for _, v := range value {
		if !v.Invisible {
			visibleFilterData = append(visibleFilterData, v)
		}
	}

	//	$plaid().stringLocation(qs).mergeQueryWithoutParams(keysInFilterData).url(window.location.href).pushState(true).go()
	r.Value(visibleFilterData).Attr("@change", web.GET().
		StringQuery(web.Var("$event.encodedFilterData")).
		ClearMergeQuery(web.Var("$event.filterKeys")).
		PushState(true).
		Go())

	return
}

func (b *VXFilterBuilder) Attr(vs ...interface{}) (r *VXFilterBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXFilterBuilder) Value(v FilterData) (r *VXFilterBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VXFilterBuilder) Translations(v FilterTranslations) (r *VXFilterBuilder) {
	b.tag.Attr(":translations", h.JSONString(v))
	return b
}

func (b *VXFilterBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}

/*
	translations: {
		type: Object,
		default: () => {
			return {
				date: {
					inTheLast: 'is in the last',
					equals: 'is equal to',
					between: 'is between',
					isAfter: 'is after',
					isAfterOrOn: 'is on or after',
					isBefore: 'is before',
					isBeforeOrOn: 'is before or on',
					days: 'days',
					months: 'months',
					and: 'and',
				},
				number: {
					equals: 'is equal to',
					between: 'between',
					greaterThan: 'is greater than',
					lessThan: 'is less than',
					and: 'and',
				},
				string: {
					equals: 'is equal to',
					contains: 'contains',
				},
				clear: 'Clear',
				filters: 'Filters',
				filter: 'Filter',
				done: 'Done',
			};
		},
*/
type FilterTranslations struct {
	Clear   string `json:"clear,omitempty"`
	Done    string `json:"done,omitempty"`
	Filters string `json:"filters,omitempty"`
	Filter  string `json:"filter,omitempty"`
	Date    struct {
		InTheLast    string `json:"inTheLast,omitempty"`
		Equals       string `json:"equals,omitempty"`
		Between      string `json:"between,omitempty"`
		IsAfter      string `json:"isAfter,omitempty"`
		IsAfterOrOn  string `json:"isAfterOrOn,omitempty"`
		IsBefore     string `json:"isBefore,omitempty"`
		IsBeforeOrOn string `json:"isBeforeOrOn,omitempty"`
		Days         string `json:"days,omitempty"`
		Months       string `json:"months,omitempty"`
		And          string `json:"and,omitempty"`
		To           string `json:"to,omitempty"`
	} `json:"date,omitempty"`

	Number struct {
		Equals      string `json:"equals,omitempty"`
		Between     string `json:"between,omitempty"`
		GreaterThan string `json:"greaterThan,omitempty"`
		LessThan    string `json:"lessThan,omitempty"`
		And         string `json:"and,omitempty"`
	} `json:"number,omitempty"`

	String struct {
		Equals   string `json:"equals,omitempty"`
		Contains string `json:"contains,omitempty"`
	} `json:"string,omitempty"`

	MultipleSelect struct {
		In    string `json:"in,omitempty"`
		NotIn string `json:"notIn,omitempty"`
	} `json:"multipleSelect,omitempty"`
}

type FilterItemType string

const (
	ItemTypeDate           FilterItemType = "DateItem"
	ItemTypeSelect         FilterItemType = "SelectItem"
	ItemTypeMultipleSelect FilterItemType = "MultipleSelectItem"
	ItemTypeLinkageSelect  FilterItemType = "LinkageSelectItem"
	ItemTypeNumber         FilterItemType = "NumberItem"
	ItemTypeString         FilterItemType = "StringItem"
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
	ModifierIn           FilterItemModifier = "in"           // String
	ModifierNotIn        FilterItemModifier = "notIn"        // String
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
	Text         string `json:"text,omitempty"`
	Value        string `json:"value,omitempty"`
	SQLCondition string `json:"-"`
}

type FilterLinkageSelectData struct {
	Items            [][]*LinkageSelectItem `json:"items,omitempty"`
	Labels           []string               `json:"labels,omitempty"`
	SelectOutOfOrder bool                   `json:"selectOutOfOrder,omitempty"`
	SQLConditions    []string               `json:"-"`
}

type FilterItem struct {
	Key               string                  `json:"key,omitempty"`
	Label             string                  `json:"label,omitempty"`
	ItemType          FilterItemType          `json:"itemType,omitempty"`
	Selected          bool                    `json:"selected,omitempty"`
	Modifier          FilterItemModifier      `json:"modifier,omitempty"`
	ValueIs           string                  `json:"valueIs,omitempty"`
	ValuesAre         []string                `json:"valuesAre,omitempty"`
	ValueFrom         string                  `json:"valueFrom,omitempty"`
	ValueTo           string                  `json:"valueTo,omitempty"`
	InTheLastValue    string                  `json:"inTheLastValue,omitempty"`
	InTheLastUnit     FilterItemInTheLastUnit `json:"inTheLastUnit,omitempty"`
	Timezone          FilterItemTimezone      `json:"timezone,omitempty"`
	SQLCondition      string                  `json:"-"`
	Options           []*SelectItem           `json:"options,omitempty"`
	LinkageSelectData FilterLinkageSelectData `json:"linkageSelectData,omitempty"`
	Invisible         bool                    `json:"invisible,omitempty"`
}

func (fd FilterData) Clone() (r FilterData) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(fd)
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&r)
	if err != nil {
		panic(err)
	}
	return
}

func (fd FilterData) getSQLCondition(key string, val string) string {
	it := fd.getFilterItem(key)
	if it == nil {
		return ""
	}

	// If item type is ItemTypeSelect and value is not nil, we use option's SQLCondition instead of item SQLCondition if option's SQLCondition present.
	if it.ItemType == ItemTypeSelect && val != "" {
		for _, option := range it.Options {
			if option.Value == val && option.SQLCondition != "" {
				return option.SQLCondition
			}
		}
	}

	return it.SQLCondition
}

func (fd FilterData) getFilterItem(key string) *FilterItem {
	for _, it := range fd {
		if it.Key == key {
			return it
		}
	}

	return nil
}

var sqlOps = map[string]string{
	"":      "=",
	"gte":   ">=",
	"lte":   "<=",
	"gt":    ">",
	"lt":    "<",
	"ilike": "ILIKE",
	"in":    "IN",
	"notIn": "NOT IN",
}

const SQLOperatorPlaceholder = "{op}"

func (fd FilterData) SetByQueryString(qs string) (sqlCondition string, sqlArgs []interface{}) {
	queryMap, err := url.ParseQuery(qs)

	if err != nil {
		panic(err)
	}

	var conds []string

	var keys = make([]string, len(queryMap))
	i := 0
	for k := range queryMap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	var keyModValueMap = map[string]map[string]string{}
	for _, k := range keys {
		v := queryMap[k]
		segs := strings.Split(k, ".")
		var mod = ""
		key := k
		val := v[0]
		if len(segs) > 1 {
			key = segs[0]
			mod = segs[1]
		}
		it := fd.getFilterItem(key)
		if it == nil {
			continue
		}

		if _, ok := keyModValueMap[key]; !ok {
			keyModValueMap[key] = map[string]string{}
		}

		keyModValueMap[key][mod] = v[0]

		if it.ItemType == ItemTypeLinkageSelect {
			vals := strings.Split(val, ",")
			for i, v := range vals {
				if v != "" {
					conds = append(conds, it.LinkageSelectData.SQLConditions[i])
					sqlArgs = append(sqlArgs, v)
				}
			}
		} else {
			sqlc := fd.getSQLCondition(key, v[0])
			if len(sqlc) > 0 {
				if it.ItemType == ItemTypeDate {
					val = unixToDatetime(val, it.Timezone == TimezoneUTC, 0)
				}

				// Compose operator into sql condition. If you want to use multiple operators you have to use {op}, '%s' is not supported
				// e.g.
				// "source_b %s ?"                        ==> "source_b = ?"
				// "source_b {op} ?"                      ==> "source_b = ?"
				// "source_b {op} ? AND source_c {op} ?"  ==> "source_b = ? AND source_c = ?"
				if strings.Contains(sqlc, "%s") {
					// This is for backward compatibility
					conds = append(conds, fmt.Sprintf(sqlc, sqlOps[mod]))
				} else {
					conds = append(conds, strings.NewReplacer(SQLOperatorPlaceholder, sqlOps[mod]).Replace(sqlc))
				}

				// Prepare value Args for sql condition.
				// e.g.  assume value is "1"
				// "source_b = ?"                           ==>   []interface{}{"1"}
				// "source_b = ? OR source_c = ?"           ==>   []interface{}{"1", "1"}
				// "source_b ilike ? OR source_c ilike ?"   ==>   []interface{}{"%1%", "%1%"}
				valCount := strings.Count(sqlc, "?")
				for i := 0; i < valCount; i++ {
					switch mod {
					case "ilike":
						sqlArgs = append(sqlArgs, fmt.Sprintf("%%%s%%", val))
					case "in", "notIn":
						sqlArgs = append(sqlArgs, strings.Split(val, ","))
					default:
						sqlArgs = append(sqlArgs, val)
					}
				}
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
					it.ValueFrom = unixToDatetimeWithFormat(mv["gte"], it.Timezone == TimezoneUTC, 0, "2006-01-02 15:04")
					it.ValueTo = unixToDatetimeWithFormat(mv["lt"], it.Timezone == TimezoneUTC, 0, "2006-01-02 15:04")
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

					if it.ItemType == ItemTypeMultipleSelect {
						switch mod {
						case "in":
							it.Modifier = ModifierIn
						case "notIn":
							it.Modifier = ModifierNotIn
						default:
							it.Modifier = ModifierIn
						}
						if v != "" {
							it.ValuesAre = strings.Split(v, ",")
						}
						continue
					}
					if it.ItemType == ItemTypeLinkageSelect {
						if v != "" {
							it.ValuesAre = strings.Split(v, ",")
						}
						continue
					}

					if it.ItemType == ItemTypeDate {
						it.ValueIs = unixToDate(v, it.Timezone == TimezoneUTC, 0)
						if mod == "gte" {
							it.Modifier = ModifierIsAfterOrOn
							it.ValueIs = unixToDate(v, it.Timezone == TimezoneUTC, 0)
						}
						if mod == "gt" {
							it.Modifier = ModifierIsAfter
							it.ValueIs = unixToDate(v, it.Timezone == TimezoneUTC, -1)
						}
						if mod == "lt" {
							it.Modifier = ModifierIsBefore
							it.ValueIs = unixToDate(v, it.Timezone == TimezoneUTC, 0)
						}
						if mod == "lte" {
							it.Modifier = ModifierIsBeforeOrOn
							it.ValueIs = unixToDate(v, it.Timezone == TimezoneUTC, -1)
						}
						continue
					}

					it.ValueIs = v
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
	return unixToDatetimeWithFormat(u, utc, sub, "2006-01-02")
}

func unixToDatetime(u string, utc bool, sub int) string {
	return unixToDatetimeWithFormat(u, utc, sub, time.RFC3339)
}

func unixToDatetimeWithFormat(u string, utc bool, sub int, format string) string {
	return unixToTime(u, utc, sub).Format(format)
}

func unixToTime(u string, utc bool, sub int) time.Time {
	unix, _ := strconv.ParseInt(u, 10, 64)
	d := time.Unix(unix, 0).Add(time.Duration(24 * sub * int(time.Hour)))
	if utc {
		d = d.UTC()
	}

	return d
}
