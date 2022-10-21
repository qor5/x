package presets

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/goplaid/web"
	"github.com/goplaid/x/i18n"
	v "github.com/goplaid/x/vuetify"
	"github.com/sunfmin/reflectutils"
	h "github.com/theplant/htmlgo"
	"github.com/thoas/go-funk"
)

type FieldContext struct {
	Name            string
	FormKey         string
	Label           string
	Errors          []string
	ModelInfo       *ModelInfo
	ListItemBuilder *FieldsBuilder
	Context         context.Context
	Disabled        bool
}

func (fc *FieldContext) StringValue(obj interface{}) (r string) {
	val := fc.Value(obj)
	switch vt := val.(type) {
	case []rune:
		return string(vt)
	case []byte:
		return string(vt)
	}
	return fmt.Sprint(val)
}

func (fc *FieldContext) Value(obj interface{}) (r interface{}) {
	fieldName := fc.Name
	return reflectutils.MustGet(obj, fieldName)
}

func (fc *FieldContext) ContextValue(key interface{}) (r interface{}) {
	if fc.Context == nil {
		return
	}
	return fc.Context.Value(key)
}

type FieldBuilder struct {
	NameLabel
	compFunc        FieldComponentFunc
	setterFunc      FieldSetterFunc
	context         context.Context
	listItemBuilder *FieldsBuilder
}

func (b *FieldsBuilder) appendNewFieldWithName(name string) (r *FieldBuilder) {
	r = &FieldBuilder{}

	if b.model == nil {
		panic("model must be provided")
	}

	fType := reflectutils.GetType(b.model, name)
	if fType == nil {
		fType = reflect.TypeOf("")
	}

	// if b.defaults == nil {
	// 	panic("field defaults must be provided")
	// }

	// ft := b.defaults.fieldTypeByTypeOrCreate(fType)
	r.name = name
	// r.ComponentFunc(ft.compFunc).
	// 	SetterFunc(ft.setterFunc)
	b.fields = append(b.fields, r)
	return
}

func emptyComponentFunc(obj interface{}, field *FieldContext, ctx *web.EventContext) (r h.HTMLComponent) {
	log.Printf("No ComponentFunc for field %v\n", field.Name)
	return
}

func (b *FieldBuilder) Label(v string) (r *FieldBuilder) {
	b.label = v
	return b
}

func (b *FieldBuilder) Clone() (r *FieldBuilder) {
	r = &FieldBuilder{}
	r.name = b.name
	r.label = b.label
	r.compFunc = b.compFunc
	r.setterFunc = b.setterFunc
	return r
}

func (b *FieldBuilder) ComponentFunc(v FieldComponentFunc) (r *FieldBuilder) {
	if v == nil {
		panic("value required")
	}
	b.compFunc = v
	return b
}

func (b *FieldBuilder) SetterFunc(v FieldSetterFunc) (r *FieldBuilder) {
	b.setterFunc = v
	return b
}

func (b *FieldBuilder) WithContextValue(key interface{}, val interface{}) (r *FieldBuilder) {
	if b.context == nil {
		b.context = context.Background()
	}
	b.context = context.WithValue(b.context, key, val)
	return b
}

type NameLabel struct {
	name  string
	label string
}

type FieldsBuilder struct {
	model       interface{}
	defaults    *FieldDefaults
	fieldLabels []string
	fields      []*FieldBuilder
	// string / []string / *FieldsSection
	fieldsLayout []interface{}
}

type FieldsSection struct {
	Title string
	Rows  [][]string
}

func NewFieldsBuilder() *FieldsBuilder {
	return &FieldsBuilder{}
}

func (b *FieldsBuilder) Defaults(v *FieldDefaults) (r *FieldsBuilder) {
	b.defaults = v
	return b
}

func (b *FieldsBuilder) Unmarshal(toObj interface{}, info *ModelInfo, removeDeletedAndSort bool, ctx *web.EventContext) (vErr web.ValidationErrors) {
	t := reflect.TypeOf(toObj)
	if t.Kind() != reflect.Ptr {
		panic("toObj must be pointer")
	}

	var fromObj = reflect.New(t.Elem()).Interface()
	// don't panic for fields that set in SetterFunc
	_ = ctx.UnmarshalForm(fromObj)
	// testingutils.PrintlnJson("Unmarshal fromObj", fromObj)

	modifiedIndexes := ContextModifiedIndexesBuilder(ctx).FromHidden(ctx.R)

	return b.setObjectFields(fromObj, toObj, &FieldContext{
		ModelInfo: info,
	}, removeDeletedAndSort, modifiedIndexes, ctx)
}

func (b *FieldsBuilder) setObjectFields(fromObj interface{}, toObj interface{}, parent *FieldContext, removeDeletedAndSort bool, modifiedIndexes *ModifiedIndexesBuilder, ctx *web.EventContext) (vErr web.ValidationErrors) {

	for _, f := range b.fields {
		info := parent.ModelInfo
		if info != nil {
			if info.Verifier().Do(PermCreate).ObjectOn(toObj).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil && info.Verifier().Do(PermUpdate).ObjectOn(toObj).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil {
				continue
			}
		}

		if f.listItemBuilder != nil {
			formKey := f.name
			if parent != nil && parent.FormKey != "" {
				formKey = fmt.Sprintf("%s.%s", parent.FormKey, f.name)
			}

			b.setWithChildFromObjs(fromObj, formKey, f, info, modifiedIndexes, toObj, removeDeletedAndSort, ctx)

			b.setToObjNilOrDelete(toObj, formKey, f, modifiedIndexes, removeDeletedAndSort)

			continue
		}

		if f.setterFunc == nil {
			val, err1 := reflectutils.Get(fromObj, f.name)
			if err1 != nil {
				continue
			}
			_ = reflectutils.Set(toObj, f.name, val)
			// fmt.Printf("fromObj %#+v, f.name %#+v, toObj %#+v, val %#+v\n", fromObj, f.name, toObj, val)
			continue
		}

		keyPath := f.name
		if parent != nil && parent.FormKey != "" {
			keyPath = fmt.Sprintf("%s.%s", parent.FormKey, f.name)
		}

		err1 := f.setterFunc(toObj, &FieldContext{
			ModelInfo: info,
			FormKey:   keyPath,
			Name:      f.name,
			Label:     b.getLabel(f.NameLabel),
		}, ctx)
		if err1 != nil {
			vErr.FieldError(f.name, err1.Error())
		}
	}
	return
}

func (b *FieldsBuilder) setToObjNilOrDelete(toObj interface{}, formKey string, f *FieldBuilder, modifiedIndexes *ModifiedIndexesBuilder, removeDeletedAndSort bool) {
	childToObjs := reflectutils.MustGet(toObj, f.name)
	if childToObjs == nil {
		return
	}
	j := 0
	if !removeDeletedAndSort {
		funk.ForEach(childToObjs, func(childToObj interface{}) {
			defer func() { j++ }()
			sliceFieldName := fmt.Sprintf("%s[%d]", f.name, j)
			if modifiedIndexes.DeletedContains(formKey, j) {
				err := reflectutils.Set(toObj, sliceFieldName, nil)
				if err != nil {
					panic(err)
				}
			}
		})
		return
	}

	t := reflectutils.GetType(toObj, f.name)
	newSlice := reflect.MakeSlice(t, 0, 0)
	modifiedIndexes.SortedForEach(childToObjs, formKey, func(obj interface{}, i int) {
		// remove deleted
		if modifiedIndexes.DeletedContains(formKey, i) {
			return
		}
		newSlice = reflect.Append(newSlice, reflect.ValueOf(obj))
	})

	err := reflectutils.Set(toObj, f.name, newSlice.Interface())
	if err != nil {
		panic(err)
	}

	return
}

func (b *FieldsBuilder) setWithChildFromObjs(
	fromObj interface{},
	formKey string,
	f *FieldBuilder,
	info *ModelInfo,
	modifiedIndexes *ModifiedIndexesBuilder,
	toObj interface{},
	removeDeletedAndSort bool,
	ctx *web.EventContext) {

	childFromObjs := reflectutils.MustGet(fromObj, f.name)
	if childFromObjs == nil || reflect.TypeOf(childFromObjs).Kind() != reflect.Slice {
		return
	}

	var i = 0
	funk.ForEach(childFromObjs, func(childFromObj interface{}) {
		defer func() { i++ }()
		if childFromObj == nil {
			return
		}
		// if is deleted, do nothing, later, it will be set to nil
		if modifiedIndexes.DeletedContains(formKey, i) {
			return
		}

		sliceFieldName := fmt.Sprintf("%s[%d]", f.name, i)

		pf := &FieldContext{
			ModelInfo: info,
			FormKey:   fmt.Sprintf("%s[%d]", formKey, i),
		}

		childToObj := reflectutils.MustGet(toObj, sliceFieldName)
		if childToObj == nil {
			arrayElementType := reflectutils.GetType(toObj, sliceFieldName)

			if arrayElementType.Kind() == reflect.Ptr {
				arrayElementType = arrayElementType.Elem()
			} else {
				panic(fmt.Sprintf("%s must be a pointer", sliceFieldName))
			}

			err := reflectutils.Set(toObj, sliceFieldName, reflect.New(arrayElementType).Interface())
			if err != nil {
				panic(err)
			}
			childToObj = reflectutils.MustGet(toObj, sliceFieldName)
		}

		// fmt.Printf("childFromObj %#+v\n", childFromObj)
		// fmt.Printf("childToObj %#+v\n", childToObj)
		// fmt.Printf("fieldContext %#+v\n", pf)
		f.listItemBuilder.setObjectFields(childFromObj, childToObj, pf, removeDeletedAndSort, modifiedIndexes, ctx)
	})

}

func (b *FieldsBuilder) Clone() (r *FieldsBuilder) {
	r = &FieldsBuilder{
		model:       b.model,
		defaults:    b.defaults,
		fieldLabels: b.fieldLabels,
	}
	return
}

func (b *FieldsBuilder) Model(v interface{}) (r *FieldsBuilder) {
	b.model = v
	return b
}

func (b *FieldsBuilder) Field(name string) (r *FieldBuilder) {
	r = b.getField(name)
	if r != nil {
		return
	}

	r = b.appendNewFieldWithName(name)
	return
}

func (b *FieldsBuilder) ListField(name string, listItemBuilder *FieldsBuilder) (r *FieldBuilder) {
	r = b.Field(name)
	if listItemBuilder.defaults == nil {
		listItemBuilder.Defaults(b.defaults)
	}
	r.listItemBuilder = listItemBuilder
	return
}

func (b *FieldsBuilder) Labels(vs ...string) (r *FieldsBuilder) {
	b.fieldLabels = append(b.fieldLabels, vs...)
	return b
}

// humanizeString humanize separates string based on capitalizd letters
// e.g. "OrderItem" -> "Order Item, CNNName to CNN Name"
func humanizeString(str string) string {
	var human []rune
	input := []rune(str)
	for i, l := range input {
		if i > 0 && unicode.IsUpper(l) {
			if (!unicode.IsUpper(input[i-1]) && input[i-1] != ' ') || (i+1 < len(input) && !unicode.IsUpper(input[i+1]) && input[i+1] != ' ' && input[i-1] != ' ') {
				human = append(human, rune(' '))
			}
		}
		human = append(human, l)
	}
	return strings.Title(string(human))
}

func (b *FieldsBuilder) getLabel(field NameLabel) (r string) {
	if len(field.label) > 0 {
		return field.label
	}

	for i := 0; i < len(b.fieldLabels)-1; i = i + 2 {
		if b.fieldLabels[i] == field.name {
			return b.fieldLabels[i+1]
		}
	}

	return humanizeString(field.name)
}

func (b *FieldsBuilder) getFieldOrDefault(name string) (r *FieldBuilder) {
	r = b.getField(name)
	if r.compFunc == nil {
		if b.defaults == nil {
			panic("field defaults must be provided")
		}

		fType := reflectutils.GetType(b.model, name)
		if fType == nil {
			fType = reflect.TypeOf("")
		}

		ft := b.defaults.fieldTypeByTypeOrCreate(fType)
		r.ComponentFunc(ft.compFunc).
			SetterFunc(ft.setterFunc)
	}
	return
}

func (b *FieldsBuilder) getField(name string) (r *FieldBuilder) {
	for _, f := range b.fields {
		if f.name == name {
			return f
		}
	}
	return
}

func (b *FieldsBuilder) Only(vs ...interface{}) (r *FieldsBuilder) {
	if len(vs) == 0 {
		return b
	}

	r = b.Clone()

	r.fieldsLayout = vs
	for _, iv := range vs {
		switch t := iv.(type) {
		case string:
			r.appendFieldAfterClone(b, t)
		case []string:
			for _, n := range t {
				r.appendFieldAfterClone(b, n)
			}
		case *FieldsSection:
			for _, row := range t.Rows {
				for _, n := range row {
					r.appendFieldAfterClone(b, n)
				}
			}
		default:
			panic("unknown fields layout, must be string/[]string/*FieldsSection")
		}
	}

	return
}

func (b *FieldsBuilder) appendFieldAfterClone(ob *FieldsBuilder, name string) {
	f := ob.getField(name)
	if f == nil {
		b.appendNewFieldWithName(name)
	} else {
		b.fields = append(b.fields, f.Clone())
	}
}

func (b *FieldsBuilder) Except(patterns ...string) (r *FieldsBuilder) {
	if len(patterns) == 0 {
		return
	}

	r = &FieldsBuilder{fieldLabels: b.fieldLabels}
	for _, f := range b.fields {
		if hasMatched(patterns, f.name) {
			continue
		}
		r.fields = append(r.fields, f.Clone())
	}
	return
}

func (b *FieldsBuilder) String() (r string) {
	var names []string
	for _, f := range b.fields {
		names = append(names, f.name)
	}
	return fmt.Sprint(names)
}

func (b *FieldsBuilder) ToComponent(info *ModelInfo, obj interface{}, ctx *web.EventContext) h.HTMLComponent {
	modifiedIndexes := ContextModifiedIndexesBuilder(ctx)
	return b.toComponentWithFormValueKey(info, obj, "", modifiedIndexes, ctx)
}

func (b *FieldsBuilder) toComponentWithFormValueKey(info *ModelInfo, obj interface{}, parentFormValueKey string, modifiedIndexes *ModifiedIndexesBuilder, ctx *web.EventContext) h.HTMLComponent {

	var comps []h.HTMLComponent
	if parentFormValueKey == "" {
		comps = append(comps, modifiedIndexes.ToFormHidden())
	}

	vErr, _ := ctx.Flash.(*web.ValidationErrors)
	if vErr == nil {
		vErr = &web.ValidationErrors{}
	}

	id, err := reflectutils.Get(obj, "ID")
	edit := false
	if err == nil && len(fmt.Sprint(id)) > 0 && fmt.Sprint(id) != "0" {
		edit = true
	}

	layout := b.fieldsLayout
	if layout == nil {
		layout = make([]interface{}, 0, len(b.fields))
		for _, f := range b.fields {
			layout = append(layout, f.name)
		}
	}
	for _, iv := range layout {
		var comp h.HTMLComponent
		switch t := iv.(type) {
		case string:
			comp = b.fieldToComponentWithFormValueKey(info, obj, parentFormValueKey, ctx, t, id, edit, vErr)
		case []string:
			colsComp := make([]h.HTMLComponent, 0, len(t))
			for _, n := range t {
				fComp := b.fieldToComponentWithFormValueKey(info, obj, parentFormValueKey, ctx, n, id, edit, vErr)
				if fComp == nil {
					continue
				}
				colsComp = append(colsComp, v.VCol(fComp).Class("pr-4"))
			}
			if len(colsComp) > 0 {
				comp = v.VRow(colsComp...).NoGutters(true)
			}
		case *FieldsSection:
			rowsComp := make([]h.HTMLComponent, 0, len(t.Rows))
			for _, row := range t.Rows {
				colsComp := make([]h.HTMLComponent, 0, len(row))
				for _, n := range row {
					fComp := b.fieldToComponentWithFormValueKey(info, obj, parentFormValueKey, ctx, n, id, edit, vErr)
					if fComp == nil {
						continue
					}
					colsComp = append(colsComp, v.VCol(fComp).Class("pr-4"))
				}
				if len(colsComp) > 0 {
					rowsComp = append(rowsComp, v.VRow(colsComp...).NoGutters(true))
				}
			}
			if len(rowsComp) > 0 {
				var titleComp h.HTMLComponent
				if t.Title != "" {
					titleComp = h.Label(t.Title).Class("v-label theme--light text-caption")
				}
				comp = h.Div(
					titleComp,
					v.VCard(rowsComp...).Elevation(1).Class("mx-1 mt-1 mb-4 px-4 pb-0 pt-4"),
				)
			}
		default:
			panic("unknown fields layout, must be string/[]string/*FieldsSection")
		}
		if comp == nil {
			continue
		}
		comps = append(comps, comp)
	}

	return h.Components(comps...)
}

func (b *FieldsBuilder) fieldToComponentWithFormValueKey(info *ModelInfo, obj interface{}, parentFormValueKey string, ctx *web.EventContext, name string, id interface{}, edit bool, vErr *web.ValidationErrors) h.HTMLComponent {
	f := b.getFieldOrDefault(name)
	// if f.compFunc == nil {
	// 	return nil
	// }
	if info != nil && info.Verifier().Do(PermGet).ObjectOn(obj).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil {
		return nil
	}

	label := b.getLabel(f.NameLabel)
	if info != nil {
		label = i18n.PT(ctx.R, ModelsI18nModuleKey, info.Label(), b.getLabel(f.NameLabel))
	}

	contextKeyPath := f.name
	if parentFormValueKey != "" {
		contextKeyPath = fmt.Sprintf("%s.%s", parentFormValueKey, f.name)
	}

	disabled := false
	if info != nil {
		if edit {
			disabled = info.Verifier().Do(PermUpdate).ObjectOn(obj).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil
		} else {
			disabled = info.Verifier().Do(PermCreate).ObjectOn(obj).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil
		}
	}
	return f.compFunc(obj, &FieldContext{
		ModelInfo:       info,
		Name:            f.name,
		FormKey:         contextKeyPath,
		Label:           label,
		Errors:          vErr.GetFieldErrors(f.name),
		ListItemBuilder: f.listItemBuilder,
		Context:         f.context,
		Disabled:        disabled,
	}, ctx)
}

type RowFunc func(obj interface{}, formKey string, content h.HTMLComponent, ctx *web.EventContext) h.HTMLComponent

func defaultRowFunc(obj interface{}, formKey string, content h.HTMLComponent, ctx *web.EventContext) h.HTMLComponent {
	return content
}

func (b *FieldsBuilder) ToComponentForEach(field *FieldContext, slice interface{}, ctx *web.EventContext, rowFunc RowFunc) h.HTMLComponent {
	var info *ModelInfo
	var parentKeyPath = ""
	if field != nil {
		info = field.ModelInfo
		parentKeyPath = field.FormKey
	}
	if rowFunc == nil {
		rowFunc = defaultRowFunc
	}
	var r []h.HTMLComponent
	modifiedIndexes := ContextModifiedIndexesBuilder(ctx)

	modifiedIndexes.SortedForEach(slice, parentKeyPath, func(obj interface{}, i int) {
		if modifiedIndexes.DeletedContains(parentKeyPath, i) {
			return
		}

		formKey := fmt.Sprintf("%s[%d]", parentKeyPath, i)
		comps := b.toComponentWithFormValueKey(info, obj, formKey, modifiedIndexes, ctx)
		r = append(r, rowFunc(obj, formKey, comps, ctx))
	})

	return h.Components(r...)
}

type ModifiedIndexesBuilder struct {
	deletedValues map[string][]string
	sortedValues  map[string][]string
}

type deletedIndexBuilderKeyType int

const theDeleteIndexBuilderKey deletedIndexBuilderKeyType = iota

const deletedHiddenNamePrefix = "__Deleted"
const sortedHiddenNamePrefix = "__Sorted"

func ContextModifiedIndexesBuilder(ctx *web.EventContext) (r *ModifiedIndexesBuilder) {
	r, ok := ctx.R.Context().Value(theDeleteIndexBuilderKey).(*ModifiedIndexesBuilder)
	if !ok {
		r = &ModifiedIndexesBuilder{}
		ctx.R = ctx.R.WithContext(context.WithValue(ctx.R.Context(), theDeleteIndexBuilderKey, r))
	}
	return r
}

func (b *ModifiedIndexesBuilder) AppendDeleted(sliceFormKey string, index int) (r *ModifiedIndexesBuilder) {
	if b.deletedValues == nil {
		b.deletedValues = make(map[string][]string)
	}
	b.deletedValues[sliceFormKey] = append(b.deletedValues[sliceFormKey], fmt.Sprint(index))
	return b
}

func (b *ModifiedIndexesBuilder) SetSorted(sliceFormKey string, indexes []string) (r *ModifiedIndexesBuilder) {
	if b.sortedValues == nil {
		b.sortedValues = make(map[string][]string)
	}
	b.sortedValues[sliceFormKey] = indexes
	return b
}

func (b *ModifiedIndexesBuilder) DeletedContains(sliceFormKey string, index int) (r bool) {

	if b.deletedValues == nil {
		return false
	}
	if b.deletedValues[sliceFormKey] == nil {
		return false
	}
	sIndex := fmt.Sprint(index)
	for _, v := range b.deletedValues[sliceFormKey] {
		if v == sIndex {
			return true
		}
	}
	return false
}

func (b *ModifiedIndexesBuilder) SortedForEach(slice interface{}, sliceFormKey string, f func(obj interface{}, i int)) {
	sortedIndexes, ok := b.sortedValues[sliceFormKey]
	if !ok {
		i := 0
		funk.ForEach(slice, func(obj interface{}) {
			defer func() { i++ }()
			f(obj, i)
		})
		return
	}

	sliceLen := reflect.ValueOf(slice).Len()
	for j1 := 0; j1 < sliceLen; j1++ {
		if funk.Contains(sortedIndexes, fmt.Sprint(j1)) {
			continue
		}
		sortedIndexes = append(sortedIndexes, fmt.Sprint(j1))
	}

	for _, j := range sortedIndexes {
		obj, err := reflectutils.Get(slice, fmt.Sprintf("[%s]", j))
		if obj == nil || err != nil {
			continue
		}
		j1, _ := strconv.Atoi(j)
		f(obj, j1)
	}

}

func deleteHiddenSliceFormKey(sliceFormKey string) string {
	return deletedHiddenNamePrefix + "." + sliceFormKey
}
func sortedHiddenSliceFormKey(sliceFormKey string) string {
	return sortedHiddenNamePrefix + "." + sliceFormKey
}

func (b *ModifiedIndexesBuilder) FromHidden(req *http.Request) (r *ModifiedIndexesBuilder) {
	if b.deletedValues == nil {
		b.deletedValues = make(map[string][]string)
	}
	if b.sortedValues == nil {
		b.sortedValues = make(map[string][]string)
	}
	for k, vs := range req.Form {
		if strings.HasPrefix(k, deletedHiddenNamePrefix) {
			b.deletedValues[k[len(deletedHiddenNamePrefix)+1:]] = strings.Split(vs[0], ",")
		}

		if strings.HasPrefix(k, sortedHiddenNamePrefix) {
			b.sortedValues[k[len(sortedHiddenNamePrefix)+1:]] = strings.Split(vs[0], ",")
		}
	}
	return b
}

func (b *ModifiedIndexesBuilder) ReversedmodifiedIndexes(sliceFormKey string) (r []int) {
	if b.deletedValues == nil {
		return
	}
	for _, v := range b.deletedValues[sliceFormKey] {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		r = append(r, i)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(r)))
	return
}

func (b *ModifiedIndexesBuilder) ToFormHidden() h.HTMLComponent {
	var hidden []h.HTMLComponent
	for sliceFormKey, values := range b.deletedValues {
		hidden = append(hidden, h.Input("").Type("hidden").
			Attr(web.VFieldName(deleteHiddenSliceFormKey(sliceFormKey))...).
			Value(strings.Join(values, ",")))

	}

	for sliceFormKey, values := range b.sortedValues {
		hidden = append(hidden, h.Input("").Type("hidden").
			Attr(web.VFieldName(sortedHiddenSliceFormKey(sliceFormKey))...).
			Value(strings.Join(values, ",")))
	}
	return h.Components(hidden...)
}
