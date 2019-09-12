package vuetify_test

import (
	"testing"

	"github.com/theplant/testingutils"
)

var setByQueryCases = []struct {
	name             string
	data             FilterData
	qs               string
	expected         FilterData
	expectedSQLConds string
	expectedSQLArgs  []interface{}
}{
	{
		name: "between",
		data: FilterData([]*FilterItem{
			{
				Key:          "created",
				ItemType:     ItemTypeDate,
				SQLCondition: "created_at %s ?",
			},
		}),
		qs: "created.lt=1554912000&created.gte=1554825600",
		expected: FilterData([]*FilterItem{
			{
				Key:       "created",
				ItemType:  ItemTypeDate,
				Modifier:  ModifierBetween,
				Selected:  true,
				ValueFrom: "2019-04-10",
				ValueTo:   "2019-04-10",
			},
		}),
		expectedSQLConds: "created_at >= ? AND created_at < ?",
		expectedSQLArgs:  []interface{}{"1554825600", "1554912000"},
	},
	{
		name: "between2",
		data: FilterData([]*FilterItem{
			{
				Key:          "created",
				Label:        "Created",
				ItemType:     ItemTypeDate,
				SQLCondition: `extract(epoch from created_at) %s ?`,
			},
			{
				Key:          "name",
				Label:        "Name",
				ItemType:     ItemTypeString,
				SQLCondition: `name %s ?`,
			},
		}),
		qs: "created.lt=1565280000&created.gte=1565107200",
		expected: FilterData([]*FilterItem{
			{
				Key:       "created",
				Label:     "Created",
				ItemType:  ItemTypeDate,
				Modifier:  ModifierBetween,
				Selected:  true,
				ValueFrom: "2019-08-07",
				ValueTo:   "2019-08-08",
			},
			{
				Key:          "name",
				Label:        "Name",
				ItemType:     ItemTypeString,
				SQLCondition: `name %s ?`,
			},
		}),
		expectedSQLConds: "extract(epoch from created_at) >= ? AND extract(epoch from created_at) < ?",
		expectedSQLArgs:  []interface{}{"1565107200", "1565280000"},
	},
	{
		name: "equals",
		data: FilterData([]*FilterItem{
			{
				Key:          "created",
				ItemType:     ItemTypeDate,
				SQLCondition: "created_at %s ?",
			},
		}),
		qs: "created=1552320000",
		expected: FilterData([]*FilterItem{
			{
				Key:      "created",
				ItemType: ItemTypeDate,
				Modifier: ModifierEquals,
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
