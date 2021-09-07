package presets

import (
	"strings"
)

type Messages struct {
	SuccessfullyUpdated            string
	New                            string
	Update                         string
	Delete                         string
	Edit                           string
	OK                             string
	Cancel                         string
	Create                         string
	DeleteConfirmationTextTemplate string
	CreatingObjectTitleTemplate    string
	EditingObjectTitleTemplate     string
	ListingObjectTitleTemplate     string
	DetailingObjectTitleTemplate   string
	FiltersClear                   string
	FiltersDone                    string
	Filters                        string
	Filter                         string
	FiltersDateInTheLast           string
	FiltersDateEquals              string
	FiltersDateBetween             string
	FiltersDateIsAfter             string
	FiltersDateIsAfterOrOn         string
	FiltersDateIsBefore            string
	FiltersDateIsBeforeOrOn        string
	FiltersDateDays                string
	FiltersDateMonths              string
	FiltersDateAnd                 string
	FiltersNumberEquals            string
	FiltersNumberBetween           string
	FiltersNumberGreaterThan       string
	FiltersNumberLessThan          string
	FiltersNumberAnd               string
	FiltersStringEquals            string
	FiltersStringContains          string
}

func (msgr *Messages) DeleteConfirmationText(id string) string {
	return strings.NewReplacer("{id}", id).
		Replace(msgr.DeleteConfirmationTextTemplate)
}

func (msgr *Messages) CreatingObjectTitle(modelName string) string {
	return strings.NewReplacer("{modelName}", modelName).
		Replace(msgr.CreatingObjectTitleTemplate)
}

func (msgr *Messages) EditingObjectTitle(label string, name string) string {
	return strings.NewReplacer("{id}", name, "{modelName}", label).
		Replace(msgr.EditingObjectTitleTemplate)
}
func (msgr *Messages) ListingObjectTitle(label string) string {
	return strings.NewReplacer("{modelName}", label).
		Replace(msgr.ListingObjectTitleTemplate)
}
func (msgr *Messages) DetailingObjectTitle(label string, name string) string {
	return strings.NewReplacer("{id}", name, "{modelName}", label).
		Replace(msgr.DetailingObjectTitleTemplate)
}

var Messages_en_US = &Messages{
	DeleteConfirmationTextTemplate: "Are you sure you want to delete object with id: {id}?",
	CreatingObjectTitleTemplate:    "New {modelName}",
	EditingObjectTitleTemplate:     "Editing {modelName} {id}",
	ListingObjectTitleTemplate:     "Listing {modelName}",
	DetailingObjectTitleTemplate:   "{modelName} {id}",
	SuccessfullyUpdated:            "Successfully Updated",
	New:                            "New",
	Update:                         "Update",
	Delete:                         "Delete",
	Edit:                           "Edit",
	OK:                             "OK",
	Cancel:                         "Cancel",
	Create:                         "Create",
	Filters:                        "Filters",
	Filter:                         "Filter",
	FiltersClear:                   "Clear",
	FiltersDone:                    "Done",
	FiltersDateInTheLast:           "is in the last",
	FiltersDateEquals:              "is equal to",
	FiltersDateBetween:             "is between",
	FiltersDateIsAfter:             "is after",
	FiltersDateIsAfterOrOn:         "is on or after",
	FiltersDateIsBefore:            "is before",
	FiltersDateIsBeforeOrOn:        "is before or on",
	FiltersDateDays:                "days",
	FiltersDateMonths:              "months",
	FiltersDateAnd:                 "and",
	FiltersNumberEquals:            "is equal to",
	FiltersNumberBetween:           "between",
	FiltersNumberGreaterThan:       "is greater than",
	FiltersNumberLessThan:          "is less than",
	FiltersNumberAnd:               "and",
	FiltersStringEquals:            "is equal to",
	FiltersStringContains:          "contains",
}

var Messages_zh_CN = &Messages{
	DeleteConfirmationTextTemplate: "你确定你要删除这个对象吗，对象ID: {id}?",
	CreatingObjectTitleTemplate:    "新建{modelName}",
	EditingObjectTitleTemplate:     "编辑{modelName} {id}",
	ListingObjectTitleTemplate:     "{modelName}列表",
	DetailingObjectTitleTemplate:   "{modelName} {id}",
	SuccessfullyUpdated:            "成功更新了",
	New:                            "新建",
	Update:                         "更新",
	Delete:                         "删除",
	Edit:                           "编辑",
	OK:                             "确定",
	Cancel:                         "取消",
	Create:                         "创建",
	Filters:                        "筛选",
	Filter:                         "筛选",
	FiltersClear:                   "清除",
	FiltersDone:                    "确定",
	FiltersDateInTheLast:           "过去",
	FiltersDateEquals:              "等于",
	FiltersDateBetween:             "之间",
	FiltersDateIsAfter:             "之后",
	FiltersDateIsAfterOrOn:         "当天或之后",
	FiltersDateIsBefore:            "之前",
	FiltersDateIsBeforeOrOn:        "当天或之前",
	FiltersDateDays:                "天",
	FiltersDateMonths:              "月",
	FiltersDateAnd:                 "和",
	FiltersNumberEquals:            "等于",
	FiltersNumberBetween:           "之间",
	FiltersNumberGreaterThan:       "大于",
	FiltersNumberLessThan:          "小于",
	FiltersNumberAnd:               "和",
	FiltersStringEquals:            "等于",
	FiltersStringContains:          "包含",
}
