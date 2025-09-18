package statusx

import (
	"github.com/samber/lo"
)

func ExtractDetail[T any](details []any) T {
	pb, ok := lo.Find(details, func(d any) bool {
		_, ok := d.(T)
		return ok
	})
	if !ok {
		var nop T
		return nop
	}
	return pb.(T)
}
