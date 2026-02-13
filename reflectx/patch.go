package reflectx

import (
	"reflect"

	"github.com/pkg/errors"
)

// Patch copies non-zero field values from patch to dest.
// Both patch and dest must be pointers to the same struct type.
// If patch is nil, dest is unchanged.
// If dest is nil, an error is returned.
// For struct fields, it recurses into them and patches each sub-field individually.
// For all other fields, a non-zero value in patch overwrites the
// corresponding field in dest; zero/nil values in patch are skipped.
func Patch(patch, dest any) error {
	pv := reflect.ValueOf(patch)
	dv := reflect.ValueOf(dest)

	if pv.Kind() != reflect.Ptr || dv.Kind() != reflect.Ptr {
		return errors.New("both patch and dest must be pointers")
	}
	if pv.IsNil() {
		return nil
	}
	if dv.IsNil() {
		return errors.New("dest must not be nil")
	}
	if pv.Type() != dv.Type() {
		return errors.New("patch and dest must be the same type")
	}

	patchFields(pv.Elem(), dv.Elem())
	return nil
}

func patchFields(pv, dv reflect.Value) {
	pt := pv.Type()
	for i := 0; i < pt.NumField(); i++ {
		pf := pv.Field(i)
		if pf.IsZero() {
			continue
		}
		if pf.Kind() == reflect.Struct && allFieldsExported(pt.Field(i).Type) {
			patchFields(pf, dv.Field(i))
			continue
		}
		dv.Field(i).Set(pf)
	}
}

func allFieldsExported(t reflect.Type) bool {
	for i := 0; i < t.NumField(); i++ {
		if !t.Field(i).IsExported() {
			return false
		}
	}
	return true
}
