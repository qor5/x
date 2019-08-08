package vuetify_test

import (
	"testing"

	"github.com/sunfmin/bran/vuetify"
	"github.com/theplant/testingutils"
)

var setByQueryCases = []struct {
	name             string
	data             vuetify.FilterData
	qs               string
	expected         vuetify.FilterData
	expectedSQLConds string
	expectedSQLArgs  []interface{}
}{
	{
		name: "between",
		data: vuetify.FilterData([]*vuetify.FilterItem{
			{
				Key:          "created",
				ItemType:     vuetify.ItemTypeDate,
				SQLCondition: "created_at %s ?",
			},
		}),
		qs: "created.lt=1554912000&created.gte=1554825600",
		expected: vuetify.FilterData([]*vuetify.FilterItem{
			{
				Key:       "created",
				ItemType:  vuetify.ItemTypeDate,
				Modifier:  vuetify.ModifierBetween,
				Selected:  true,
				ValueFrom: "2019-04-10",
				ValueTo:   "2019-04-10",
			},
		}),
		expectedSQLConds: "created_at < ? AND created_at >= ?",
		expectedSQLArgs:  []interface{}{"1554912000", "1554825600"},
	},
	{
		name: "between2",
		data: vuetify.FilterData([]*vuetify.FilterItem{
			{
				Key:          "created",
				Label:        "Created",
				ItemType:     vuetify.ItemTypeDate,
				SQLCondition: `extract(epoch from created_at) %s ?`,
			},
			{
				Key:          "name",
				Label:        "Name",
				ItemType:     vuetify.ItemTypeString,
				SQLCondition: `name %s ?`,
			},
		}),
		qs: "created.lt=1565280000&created.gte=1565107200",
		expected: vuetify.FilterData([]*vuetify.FilterItem{
			{
				Key:       "created",
				Label:     "Created",
				ItemType:  vuetify.ItemTypeDate,
				Modifier:  vuetify.ModifierBetween,
				Selected:  true,
				ValueFrom: "2019-08-07",
				ValueTo:   "2019-08-08",
			},
			{
				Key:          "name",
				Label:        "Name",
				ItemType:     vuetify.ItemTypeString,
				SQLCondition: `name %s ?`,
			},
		}),
		expectedSQLConds: "extract(epoch from created_at) < ? AND extract(epoch from created_at) >= ?",
		expectedSQLArgs:  []interface{}{"1565280000", "1565107200"},
	},
	{
		name: "equals",
		data: vuetify.FilterData([]*vuetify.FilterItem{
			{
				Key:          "created",
				ItemType:     vuetify.ItemTypeDate,
				SQLCondition: "created_at %s ?",
			},
		}),
		qs: "created=1552320000",
		expected: vuetify.FilterData([]*vuetify.FilterItem{
			{
				Key:      "created",
				ItemType: vuetify.ItemTypeDate,
				Modifier: vuetify.ModifierEquals,
				Selected: true,
				ValueIs:  "2019-03-12",
			},
		}),
		expectedSQLConds: "created_at = ?",
		expectedSQLArgs:  []interface{}{"1552320000"},
	},
}

func TestSetByQueryString(t *testing.T) {
	for _, c := range setByQueryCases {
		conds, args := c.data.SetByQueryString(c.qs)
		diff := testingutils.PrettyJsonDiff(c.expected, c.data)
		if len(diff) > 0 {
			t.Error(c.name, diff)
		}

		diff1 := testingutils.PrettyJsonDiff(c.expectedSQLConds, conds)
		if len(diff1) > 0 {
			t.Error(c.name, "conds", diff1)
		}

		diff2 := testingutils.PrettyJsonDiff(c.expectedSQLArgs, args)
		if len(diff2) > 0 {
			t.Error(c.name, "args", diff2)
		}
	}
}
