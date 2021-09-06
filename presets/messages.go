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
}
