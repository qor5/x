package vuetifyx

import (
	"context"
	"fmt"

	"github.com/qor5/web/v3"
	h "github.com/theplant/htmlgo"
)

type (
	VXChartOptionTitle struct {
		Text string `json:"text,omitempty"`
	}
	VXChartOptionXAxis struct {
		Data []string `json:"data,omitempty"`
	}
	VXChartOptionSeries struct {
		Name   string        `json:"name,omitempty"`
		Type   string        `json:"type,omitempty"`
		Smooth    bool          `json:"smooth,omitempty"`
		LineColor string        `json:"lineColor,omitempty"`
		Data      []interface{} `json:"data,omitempty"`
		IsDisabled bool          `json:"isDisabled,omitempty"`
	}
	VXChartOption struct {
		Title  *VXChartOptionTitle    `json:"title,omitempty"`
		XAxis  *VXChartOptionXAxis    `json:"xAxis,omitempty"`
		Series *[]VXChartOptionSeries `json:"series,omitempty"`
	}
)

type VXChartBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXChart() (r *VXChartBuilder) {
	r = &VXChartBuilder{
		tag: h.Tag("vx-chart"),
	}

	return
}

func (b *VXChartBuilder) Presets(v string) (r *VXChartBuilder) {
	b.tag.Attr(":presets", h.JSONString(v))
	return b
}

func (b *VXChartBuilder) Options(v interface{}) (r *VXChartBuilder) {
	b.tag.Attr(":options", h.JSONString(v))
	return b
}

func (b *VXChartBuilder) MergeOptionsCallback(v string) (r *VXChartBuilder) {
	b.tag.Attr(":merge-options-callback", v)
	return b
}

func (b *VXChartBuilder) SetAttr(k string, v interface{}) {
	b.tag.SetAttr(k, v)
}

func (b *VXChartBuilder) Attr(vs ...interface{}) (r *VXChartBuilder) {
	b.tag.Attr(vs...)
	return b
}

func (b *VXChartBuilder) Children(children ...h.HTMLComponent) (r *VXChartBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VXChartBuilder) AppendChildren(children ...h.HTMLComponent) (r *VXChartBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VXChartBuilder) PrependChildren(children ...h.HTMLComponent) (r *VXChartBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VXChartBuilder) Class(names ...string) (r *VXChartBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VXChartBuilder) ClassIf(name string, add bool) (r *VXChartBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VXChartBuilder) On(name string, value string) (r *VXChartBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VXChartBuilder) Bind(name string, value string) (r *VXChartBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VXChartBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}

func (b *VXChartBuilder) OnClick(eventFuncId string) (r *VXChartBuilder) {
	b.tag.Attr("@click", web.POST().EventFunc(eventFuncId).Go())
	return b
}

func (b *VXChartBuilder) AttrIf(key, value interface{}, add bool) (r *VXChartBuilder) {
	b.tag.AttrIf(key, value, add)
	return b
}
