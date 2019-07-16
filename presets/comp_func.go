package presets

import (
	"fmt"

	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type CompFunc func(obj interface{}, fieldName string, ctx *ui.EventContext) h.HTMLComponent

type fieldTypeCompFunc struct {
	fieldType string
	compFunc  CompFunc
}

type CompFuncRegistry struct {
	fieldTypeComponentFuncs []*fieldTypeCompFunc
}

func (b *CompFuncRegistry) ComponentFuncByType(fieldType string) (f CompFunc) {
	for _, fc := range b.fieldTypeComponentFuncs {
		if fc.fieldType == fieldType {
			return fc.compFunc
		}
	}
	return nil
}

func (b *CompFuncRegistry) RegisterComponentFunc(fieldTypes []string, f CompFunc, replaceIfExists bool) {
	for _, fc := range b.fieldTypeComponentFuncs {
		for _, ft := range fieldTypes {
			if fc.fieldType == ft {
				if replaceIfExists {
					fc.compFunc = f
					return
				} else {
					panic(fmt.Sprintf("%s already registered", ft))
				}
			}
		}
	}
	for _, ft := range fieldTypes {
		b.fieldTypeComponentFuncs = append(b.fieldTypeComponentFuncs, &fieldTypeCompFunc{
			fieldType: ft,
			compFunc:  f,
		})
	}
	return
}
