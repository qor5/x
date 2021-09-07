package vuetify

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

	h "github.com/theplant/htmlgo"
)

type FilterBuilder struct {
	value FilterData
	tag   *h.HTMLTagBuilder
}

func Filter(value FilterData) (r *FilterBuilder) {
	r = &FilterBuilder{
		value: value,
		tag:   h.Tag("vw-filter"),
	}

	r.Value(value).ReplaceWindowLocation(true)

	return
}

func (b *FilterBuilder) Value(v FilterData) (r *FilterBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *FilterBuilder) ReplaceWindowLocation(v bool) (r *FilterBuilder) {
	b.tag.Attr(":replace-window-location", fmt.Sprint(v))
	return b
}

func (b *FilterBuilder) Translations(v FilterTranslations) (r *FilterBuilder) {
	b.tag.Attr(":translations", h.JSONString(v))
	return b
}

func (b *FilterBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
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
	Text  string `json:"text,omitempty"`
	Value string `json:"value,omitempty"`
}

type FilterItem struct {
	Key            string                  `json:"key,omitempty"`
	Label          string                  `json:"label,omitempty"`
	ItemType       FilterItemType          `json:"itemType,omitempty"`
	Selected       bool                    `json:"selected,omitempty"`
	Modifier       FilterItemModifier      `json:"modifier,omitempty"`
	ValueIs        string                  `json:"valueIs,omitempty"`
	ValueFrom      string                  `json:"valueFrom,omitempty"`
	ValueTo        string                  `json:"valueTo,omitempty"`
	InTheLastValue string                  `json:"inTheLastValue,omitempty"`
	InTheLastUnit  FilterItemInTheLastUnit `json:"inTheLastUnit,omitempty"`
	Timezone       FilterItemTimezone      `json:"timezone,omitempty"`
	SQLCondition   string                  `json:"-"`
	Options        []*SelectItem           `json:"options,omitempty"`
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
	m, err := url.ParseQuery(qs)

	if err != nil {
		panic(err)
	}

	var conds []string

	var keys = make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	var keyModValueMap = map[string]map[string]string{}
	for _, k := range keys {
		v := m[k]
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
	unix, _ := strconv.ParseInt(u, 10, 64)
	d := time.Unix(unix, 0).Add(time.Duration(24 * sub * int(time.Hour)))
	if utc {
		d = d.UTC()
	}
	return d.Format("2006-01-02")
}
