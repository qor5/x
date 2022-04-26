package vuetifyx_test

import (
	"testing"

	. "github.com/goplaid/x/vuetifyx"
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
				ValueFrom: "2019-04-10 00:00",
				ValueTo:   "2019-04-11 00:00",
			},
		}),
		expectedSQLConds: "created_at >= ? AND created_at < ?",
		expectedSQLArgs:  []interface{}{"2019-04-10T00:00:00+08:00", "2019-04-11T00:00:00+08:00"},
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
				ValueFrom: "2019-08-07 00:00",
				ValueTo:   "2019-08-09 00:00",
			},
			{
				Key:          "name",
				Label:        "Name",
				ItemType:     ItemTypeString,
				SQLCondition: `name %s ?`,
			},
		}),
		expectedSQLConds: "extract(epoch from created_at) >= ? AND extract(epoch from created_at) < ?",
		expectedSQLArgs:  []interface{}{"2019-08-07T00:00:00+08:00", "2019-08-09T00:00:00+08:00"},
	},
	{
		name: "customize SQLCondition",
		data: FilterData([]*FilterItem{
			{
				Key:      "missing-source-url",
				ItemType: ItemTypeSelect,
				Options: []*SelectItem{
					{Text: "Yes", Value: "1", SQLCondition: "source_a IS NULL AND source_b IS NULL"},
					{Text: "No", Value: "0", SQLCondition: "source_a IS NOT NULL OR source_b IS NOT NULL"},
				},
			},
		}),
		qs: "missing-source-url=1",
		expected: FilterData([]*FilterItem{
			{
				Key:      "missing-source-url",
				ItemType: ItemTypeSelect,
				Modifier: ModifierEquals,
				Selected: true,
				ValueIs:  "1",
				Options: []*SelectItem{
					{Text: "Yes", Value: "1", SQLCondition: "source_a IS NULL AND source_b IS NULL"},
					{Text: "No", Value: "0", SQLCondition: "source_a IS NOT NULL OR source_b IS NOT NULL"},
				},
			},
		}),
		expectedSQLConds: "source_a IS NULL AND source_b IS NULL",
		expectedSQLArgs:  nil,
	},
	{
		name: "customize SQLCondition with single value",
		data: FilterData([]*FilterItem{
			{
				Key:      "missing-source-url",
				ItemType: ItemTypeSelect,
				Options: []*SelectItem{
					{Text: "Yes", Value: "1", SQLCondition: "source_a IS NULL AND source_b = ?"},
					{Text: "No", Value: "0", SQLCondition: "source_a IS NOT NULL OR source_b IS NOT NULL"},
				},
			},
		}),
		qs: "missing-source-url=1",
		expected: FilterData([]*FilterItem{
			{
				Key:      "missing-source-url",
				ItemType: ItemTypeSelect,
				Modifier: ModifierEquals,
				Selected: true,
				ValueIs:  "1",
				Options: []*SelectItem{
					{Text: "Yes", Value: "1", SQLCondition: "source_a IS NULL AND source_b = ?"},
					{Text: "No", Value: "0", SQLCondition: "source_a IS NOT NULL OR source_b IS NOT NULL"},
				},
			},
		}),
		expectedSQLConds: "source_a IS NULL AND source_b = ?",
		expectedSQLArgs:  []interface{}{"1"},
	},
	{
		name: "customize SQLCondition with multiple value",
		data: FilterData([]*FilterItem{
			{
				Key:      "missing-source-url",
				ItemType: ItemTypeSelect,
				Options: []*SelectItem{
					{Text: "Yes", Value: "1", SQLCondition: "source_a IS NULL AND (source_b = ? OR source_c = ?)"},
					{Text: "No", Value: "0", SQLCondition: "source_a IS NOT NULL OR source_b IS NOT NULL"},
				},
			},
		}),
		qs: "missing-source-url=1",
		expected: FilterData([]*FilterItem{
			{
				Key:      "missing-source-url",
				ItemType: ItemTypeSelect,
				Modifier: ModifierEquals,
				Selected: true,
				ValueIs:  "1",
				Options: []*SelectItem{
					{Text: "Yes", Value: "1", SQLCondition: "source_a IS NULL AND (source_b = ? OR source_c = ?)"},
					{Text: "No", Value: "0", SQLCondition: "source_a IS NOT NULL OR source_b IS NOT NULL"},
				},
			},
		}),
		expectedSQLConds: "source_a IS NULL AND (source_b = ? OR source_c = ?)",
		expectedSQLArgs:  []interface{}{"1", "1"},
	},
	{
		name: "customize SQLCondition with default operator and single value",
		data: FilterData([]*FilterItem{
			{
				Key:      "missing-source-url",
				ItemType: ItemTypeSelect,
				Options: []*SelectItem{
					{Text: "Yes", Value: "1", SQLCondition: "source_a IS NULL AND source_b %s ?"},
				},
			},
		}),
		qs: "missing-source-url=1",
		expected: FilterData([]*FilterItem{
			{
				Key:      "missing-source-url",
				ItemType: ItemTypeSelect,
				Modifier: ModifierEquals,
				Selected: true,
				ValueIs:  "1",
				Options: []*SelectItem{
					{Text: "Yes", Value: "1", SQLCondition: "source_a IS NULL AND source_b %s ?"},
				},
			},
		}),
		expectedSQLConds: "source_a IS NULL AND source_b = ?",
		expectedSQLArgs:  []interface{}{"1"},
	},
	{
		name: "customize SQLCondition with extra operator and multiple value",
		data: FilterData([]*FilterItem{
			{
				Key:      "missing-source-url",
				ItemType: ItemTypeSelect,
				Options: []*SelectItem{
					{Text: "Yes", Value: "1", SQLCondition: "source_a IS NULL AND (source_b {op} ? OR source_c {op} ?)"},
				},
			},
		}),
		qs: "missing-source-url.gt=1",
		expected: FilterData([]*FilterItem{
			{
				Key:      "missing-source-url",
				ItemType: ItemTypeSelect,
				Selected: true,
				ValueIs:  "1",
				Options: []*SelectItem{
					{Text: "Yes", Value: "1", SQLCondition: "source_a IS NULL AND (source_b {op} ? OR source_c {op} ?)"},
				},
			},
		}),
		expectedSQLConds: "source_a IS NULL AND (source_b > ? OR source_c > ?)",
		expectedSQLArgs:  []interface{}{"1", "1"},
	},
	{
		name: "ItemTypeMultipleSelect",
		data: FilterData([]*FilterItem{
			{
				Key:          "state",
				ItemType:     ItemTypeMultipleSelect,
				SQLCondition: "state %s ?",
				Options: []*SelectItem{
					{Text: "Draft", Value: "draft"},
					{Text: "Approved", Value: "approved"},
					{Text: "Rejected", Value: "rejected"},
				},
			},
		}),
		qs: "state.in=draft,rejected",
		expected: FilterData([]*FilterItem{
			{
				Key:      "state",
				ItemType: ItemTypeMultipleSelect,
				Selected: true,
				Modifier: ModifierIn,
				ValuesAre: []string{
					"draft",
					"rejected",
				},
				Options: []*SelectItem{
					{Text: "Draft", Value: "draft"},
					{Text: "Approved", Value: "approved"},
					{Text: "Rejected", Value: "rejected"},
				},
			},
		}),
		expectedSQLConds: "state IN ?",
		expectedSQLArgs:  []interface{}{[]string{"draft", "rejected"}},
	},
	{
		name: "ItemTypeLinkageSelect",
		data: FilterData([]*FilterItem{
			{
				Key:      "province_city_district",
				ItemType: ItemTypeLinkageSelect,
				LinkageSelectData: FilterLinkageSelectData{
					Items: [][]*LinkageSelectItem{
						{
							{ID: "1", Name: "浙江", ChildrenIDs: []string{"1", "2"}},
							{ID: "2", Name: "江苏", ChildrenIDs: []string{"3", "4"}},
						},
						{
							{ID: "1", Name: "杭州", ChildrenIDs: []string{"1", "2"}},
							{ID: "2", Name: "宁波", ChildrenIDs: []string{"3", "4"}},
							{ID: "3", Name: "南京", ChildrenIDs: []string{"5", "6"}},
							{ID: "4", Name: "苏州", ChildrenIDs: []string{"7", "8"}},
						},
						{

							{ID: "1", Name: "拱墅区"},
							{ID: "2", Name: "西湖区"},
							{ID: "3", Name: "镇海区"},
							{ID: "4", Name: "鄞州区"},
							{ID: "5", Name: "鼓楼区"},
							{ID: "6", Name: "玄武区"},
							{ID: "7", Name: "常熟区"},
							{ID: "8", Name: "吴江区"},
						},
					},
					Labels: []string{"Province", "City", "District"},
					SQLConditions: []string{
						"province_id = ?",
						"city_id = ?",
						"district_id = ?",
					},
				},
			},
		}),
		qs: "province_city_district=2,3,7",
		expected: FilterData([]*FilterItem{
			{
				Key:      "province_city_district",
				ItemType: ItemTypeLinkageSelect,
				Selected: true,
				Modifier: ModifierEquals,
				ValuesAre: []string{
					"2",
					"3",
					"7",
				},
				LinkageSelectData: FilterLinkageSelectData{
					Items: [][]*LinkageSelectItem{
						{
							{ID: "1", Name: "浙江", ChildrenIDs: []string{"1", "2"}},
							{ID: "2", Name: "江苏", ChildrenIDs: []string{"3", "4"}},
						},
						{
							{ID: "1", Name: "杭州", ChildrenIDs: []string{"1", "2"}},
							{ID: "2", Name: "宁波", ChildrenIDs: []string{"3", "4"}},
							{ID: "3", Name: "南京", ChildrenIDs: []string{"5", "6"}},
							{ID: "4", Name: "苏州", ChildrenIDs: []string{"7", "8"}},
						},
						{
							{ID: "1", Name: "拱墅区"},
							{ID: "2", Name: "西湖区"},
							{ID: "3", Name: "镇海区"},
							{ID: "4", Name: "鄞州区"},
							{ID: "5", Name: "鼓楼区"},
							{ID: "6", Name: "玄武区"},
							{ID: "7", Name: "常熟区"},
							{ID: "8", Name: "吴江区"},
						},
					},
					Labels: []string{"Province", "City", "District"},
				},
			},
		}),
		expectedSQLConds: "province_id = ? AND city_id = ? AND district_id = ?",
		expectedSQLArgs:  []interface{}{"2", "3", "7"},
	},
}

func TestSetByQueryString(t *testing.T) {
	for _, c := range setByQueryCases {
		t.Run(c.name, func(t *testing.T) {
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
		})
	}
}
