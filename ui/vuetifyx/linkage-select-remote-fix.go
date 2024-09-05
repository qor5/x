package vuetifyx

import (
	h "github.com/theplant/htmlgo"
)

func (b *VXLinkageSelectRemotetBuilber) IsPaging(v bool) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":is-paging", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectRemotetBuilber) ItemTitle(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":item-title", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) ItemValue(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":item-value", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) PageField(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":page-field", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) PagesField(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":pages-field", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) TotalField(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":total-field", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) ItemsField(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":items-field", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) CurrentField(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":current-field", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) SearchField(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":search-field", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectRemotetBuilber) RemoteUrl(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":remote-url", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectRemotetBuilber) Page(v int) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":page", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) PageSize(v int) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":page-size", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectRemotetBuilber) ParentField(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":parent-field", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) ParentIdField(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":parent-id-field", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) LevelField(v string) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":level-field", h.JSONString(v))
	return b
}

func (b *VXLinkageSelectRemotetBuilber) LevelStart(v int) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":level-start", h.JSONString(v))
	return b
}
func (b *VXLinkageSelectRemotetBuilber) LevelStep(v int) (r *VXLinkageSelectRemotetBuilber) {
	b.tag.Attr(":level-step", h.JSONString(v))
	return b
}

type VXLinkageSelectRemoteOptions struct {
	RemoteURL  string `json:"remoteUrl"`
	IsPaging   bool   `json:"isPaging"`
	ItemTitle  string `json:"itemTitle"`
	ItemValue  string `json:"itemValue"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	LevelStart int    `json:"levelStart"`
	LevelStep  int    `json:"levelStep"`
	Separator  string `json:"separator"`

	PageField     string `json:"pageField"`
	PagesField    string `json:"pagesField"`
	PageSizeField string `json:"pageSizeField"`
	TotalField    string `json:"totalField"`
	ItemsField    string `json:"itemsField"`
	CurrentField  string `json:"currentField"`
	SearchField   string `json:"searchField"`
	ParentField   string `json:"parentField"`
	ParentIdField string `json:"parentIdField"`
	LevelField    string `json:"levelField"`
}

func DefaultVXLinkageSelectRemoteOptions(remoteUrl string) *VXLinkageSelectRemoteOptions {
	return &VXLinkageSelectRemoteOptions{
		RemoteURL:     remoteUrl,
		IsPaging:      true,
		ItemTitle:     "Name",
		ItemValue:     "ID",
		PageField:     "page",
		PagesField:    "pages",
		PageSizeField: "pageSize",
		TotalField:    "total",
		ItemsField:    "data",
		CurrentField:  "current",
		SearchField:   "search",
		Page:          1,
		PageSize:      20,
		Separator:     "__",

		ParentField:   "parent",
		ParentIdField: "parentID",
		LevelField:    "level",
		LevelStart:    1,
		LevelStep:     1,
	}
}
