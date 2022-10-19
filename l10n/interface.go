package l10n

import "reflect"

type L10nInterface interface {
	SetLocale(locale string)
}

func IsLocalizable(obj interface{}) (IsLocalizable bool) {
	_, IsLocalizable = getStruct(reflect.TypeOf(obj)).(L10nInterface)
	return
}

func getStruct(t reflect.Type) interface{} {
	if t.Kind() == reflect.Struct {
		return reflect.New(t).Interface()
	}
	return getStruct(t.Elem())
}
