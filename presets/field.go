package presets

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/goplaid/web"
	"github.com/goplaid/x/i18n"
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

func (b *FieldsBuilder) appendNewFieldWithDefault(name string) (r *FieldBuilder) {
	r = &FieldBuilder{}

	if b.model == nil {
		panic("model must be provided")
	}

	fType := reflectutils.GetType(b.model, name)
	if fType == nil {
		fType = reflect.TypeOf("")
	}

	if b.defaults == nil {
		panic("field defaults must be provided")
	}

	ft := b.defaults.fieldTypeByTypeOrCreate(fType)
	r.name = name
	r.ComponentFunc(ft.compFunc).
		SetterFunc(ft.setterFunc)
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
}

func NewFieldsBuilder() *FieldsBuilder {
	return &FieldsBuilder{}
}

func (b *FieldsBuilder) Defaults(v *FieldDefaults) (r *FieldsBuilder) {
	b.defaults = v
	return b
}

func (b *FieldsBuilder) Unmarshal(toObj interface{}, info *ModelInfo, ctx *web.EventContext) (vErr web.ValidationErrors) {
	t := reflect.TypeOf(toObj)
	if t.Kind() != reflect.Ptr {
		panic("toObj must be pointer")
	}

	var fromObj = reflect.New(t.Elem()).Interface()
	// don't panic for fields that set in SetterFunc
	_ = ctx.UnmarshalForm(fromObj)
	// testingutils.PrintlnJson("Unmarshal fromObj", fromObj)

	return b.setObjectFields(fromObj, toObj, &FieldContext{
		ModelInfo: info,
	}, ctx)
}

func (b *FieldsBuilder) setObjectFields(fromObj interface{}, toObj interface{}, parent *FieldContext, ctx *web.EventContext) (vErr web.ValidationErrors) {

	for _, f := range b.fields {
		info := parent.ModelInfo
		if info != nil {
			if info.Verifier().Do(PermUpdate).ObjectOn(toObj).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil {
				continue
			}
		}

		if f.listItemBuilder != nil {
			childFromObjs := reflectutils.MustGet(fromObj, f.name)
			// fmt.Printf("childFromObjs %#+v, %+v\n", childFromObjs, reflect.TypeOf(childFromObjs))
			if childFromObjs == nil || reflect.TypeOf(childFromObjs).Kind() != reflect.Slice {
				continue
			}
			var i = 0
			funk.ForEach(childFromObjs, func(childFromObj interface{}) {
				if childFromObj == nil {
					return
				}
				sliceFieldName := fmt.Sprintf("%s[%d]", f.name, i)
				keyPath := sliceFieldName
				if parent != nil && parent.FormKey != "" {
					keyPath = fmt.Sprintf("%s.%s", parent.FormKey, sliceFieldName)
				}
				pf := &FieldContext{
					ModelInfo: info,
					FormKey:   keyPath,
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
				f.listItemBuilder.setObjectFields(childFromObj, childToObj, pf, ctx)
				i++
			})

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

	r = b.appendNewFieldWithDefault(name)
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

func (b *FieldsBuilder) getLabel(field NameLabel) (r string) {
	if len(field.label) > 0 {
		return field.label
	}

	for i := 0; i < len(b.fieldLabels)-1; i = i + 2 {
		if b.fieldLabels[i] == field.name {
			return b.fieldLabels[i+1]
		}
	}

	return field.name
}

func (b *FieldsBuilder) getField(name string) (r *FieldBuilder) {
	for _, f := range b.fields {
		if f.name == name {
			return f
		}
	}
	return
}

func (b *FieldsBuilder) Only(names ...string) (r *FieldsBuilder) {
	if len(names) == 0 {
		return b
	}

	r = b.Clone()

	for _, n := range names {
		f := b.getField(n)
		if f == nil {
			r.appendNewFieldWithDefault(n)
		} else {
			r.fields = append(r.fields, f.Clone())
		}
	}

	return
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
	return b.toComponentWithFormValueKey(info, obj, "", ctx)
}

func (b *FieldsBuilder) toComponentWithFormValueKey(info *ModelInfo, obj interface{}, parentFormValueKey string, ctx *web.EventContext) h.HTMLComponent {

	var comps []h.HTMLComponent

	vErr, _ := ctx.Flash.(*web.ValidationErrors)
	if vErr == nil {
		vErr = &web.ValidationErrors{}
	}

	for _, f := range b.fields {
		if f.compFunc == nil {
			continue
		}

		label := b.getLabel(f.NameLabel)
		if info != nil {
			if info.Verifier().Do(PermUpdate).ObjectOn(obj).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil {
				continue
			}
			label = i18n.PT(ctx.R, ModelsI18nModuleKey, info.Label(), b.getLabel(f.NameLabel))
		}

		contextKeyPath := f.name
		if parentFormValueKey != "" {
			contextKeyPath = fmt.Sprintf("%s.%s", parentFormValueKey, f.name)
		}

		comps = append(comps, f.compFunc(obj, &FieldContext{
			ModelInfo:       info,
			Name:            f.name,
			FormKey:         contextKeyPath,
			Label:           label,
			Errors:          vErr.GetFieldErrors(f.name),
			ListItemBuilder: f.listItemBuilder,
			Context:         f.context,
		}, ctx))
	}

	return h.Components(comps...)
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
	var i = 0

	funk.ForEach(slice, func(obj interface{}) {
		formKey := fmt.Sprintf("%s[%d]", parentKeyPath, i)
		comps := b.toComponentWithFormValueKey(info, obj, formKey, ctx)
		r = append(r, rowFunc(obj, formKey, comps, ctx))
		i++
	})

	return h.Components(r...)
}
