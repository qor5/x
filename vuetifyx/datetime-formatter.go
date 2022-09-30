package vuetifyx

import (
	"context"

	h "github.com/theplant/htmlgo"
)

type VXDateTimeFormatterBuilder struct {
	tag *h.HTMLTagBuilder
}

func VXDateTimeFormatter() (r *VXDateTimeFormatterBuilder) {
	r = &VXDateTimeFormatterBuilder{
		tag: h.Tag("vx-datetimeformatter"),
	}
	return
}

func (b *VXDateTimeFormatterBuilder) Value(unixTimestamp int64) *VXDateTimeFormatterBuilder {
	if unixTimestamp <= 0 {
		return b
	}
	b.tag.Attr(":value", h.JSONString(unixTimestamp))
	return b
}

// https://date-fns.org/v2.29.3/docs/format
func (b *VXDateTimeFormatterBuilder) Format(v string) *VXDateTimeFormatterBuilder {
	b.tag.Attr(":format", h.JSONString(v))
	return b
}

func (b *VXDateTimeFormatterBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
