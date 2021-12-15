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

const SelfFieldName = "."

type FieldContext struct {
	Name      string
	KeyPath   string
	Label     string
	Errors    []string
	ModelInfo *ModelInfo
	Context   context.Context
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
	if fieldName == SelfFieldName {
		return obj
	}
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
	compFunc   FieldComponentFunc
	setterFunc FieldSetterFunc
	context    context.Context
}

func NewField(name string) (r *FieldBuilder) {
	r = &FieldBuilder{}
	r.name = name
	r.compFunc = emptyComponentFunc
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

type FieldBuilders struct {
	modelType   reflect.Type
	model       interface{}
	defaults    *FieldDefaults
	fieldLabels []string
	fields      []*FieldBuilder
}

func NewFieldBuilders() *FieldBuilders {
	return &FieldBuilders{}
}

func (b *FieldBuilders) Model(v interface{}) (r *FieldBuilders) {
	b.modelType = reflect.TypeOf(v)
	b.model = v
	return b
}

func (b *FieldBuilders) Defaults(v *FieldDefaults) (r *FieldBuilders) {
	b.defaults = v
	return b
}

func (b *FieldBuilders) NewModel() (r interface{}) {
	return reflect.New(b.modelType.Elem()).Interface()
}

func (b *FieldBuilders) NewModelSlice() (r interface{}) {
	return reflect.New(reflect.SliceOf(b.modelType)).Interface()
}

func (b *FieldBuilders) SetModelSliceFields(toObj interface{}, info *ModelInfo, ctx *web.EventContext) (vErr web.ValidationErrors) {

	return
}
func (b *FieldBuilders) SetModelFields(toObj interface{}, info *ModelInfo, ctx *web.EventContext) (vErr web.ValidationErrors) {
	var fromObj = b.NewModel()
	// don't panic for fields that set in SetterFunc
	_ = ctx.UnmarshalForm(fromObj)

	for _, f := range b.fields {
		if info != nil {
			if info.Verifier().Do(PermUpdate).ObjectOn(toObj).SnakeOn(f.name).WithReq(ctx.R).IsAllowed() != nil {
				continue
			}
		}

		if f.setterFunc == nil {
			val, err1 := reflectutils.Get(fromObj, f.name)
			if err1 != nil {
				continue
			}
			_ = reflectutils.Set(toObj, f.name, val)
			continue
		}

		err1 := f.setterFunc(toObj, &FieldContext{
			ModelInfo: info,
			Name:      f.name,
			Label:     b.getLabel(f.NameLabel),
		}, ctx)
		if err1 != nil {
			vErr.FieldError(f.name, err1.Error())
		}
	}
	return
}

func (b *FieldBuilders) Clone() (r *FieldBuilders) {
	r = &FieldBuilders{
		model:       b.model,
		defaults:    b.defaults,
		fieldLabels: b.fieldLabels,
	}
	return
}

func (b *FieldBuilders) Field(name string) (r *FieldBuilder) {
	r = b.GetField(name)
	if r != nil {
		return
	}

	r = NewField(name)
	b.fields = append(b.fields, r)
	return
}

func (b *FieldBuilders) Labels(vs ...string) (r *FieldBuilders) {
	b.fieldLabels = append(b.fieldLabels, vs...)
	return b
}

func (b *FieldBuilders) getLabel(field NameLabel) (r string) {
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

func (b *FieldBuilders) GetField(name string) (r *FieldBuilder) {
	for _, f := range b.fields {
		if f.name == name {
			return f
		}
	}
	return
}

func (b *FieldBuilders) Only(names ...string) (r *FieldBuilders) {
	if len(names) == 0 {
		return b
	}

	r = b.Clone()

	for _, n := range names {
		f := b.GetField(n)
		if f == nil {
			fType := reflectutils.GetType(b.model, n)
			if fType == nil {
				fType = reflect.TypeOf("")
			}

			ft := b.defaults.fieldTypeByTypeOrCreate(fType)
			r.Field(n).
				ComponentFunc(ft.compFunc).
				SetterFunc(ft.setterFunc)
		} else {
			r.fields = append(r.fields, f.Clone())
		}
	}

	return
}

func (b *FieldBuilders) Except(patterns ...string) (r *FieldBuilders) {
	if len(patterns) == 0 {
		return
	}

	r = &FieldBuilders{fieldLabels: b.fieldLabels}
	for _, f := range b.fields {
		if hasMatched(patterns, f.name) {
			continue
		}
		r.fields = append(r.fields, f.Clone())
	}
	return
}

func (b *FieldBuilders) String() (r string) {
	var names []string
	for _, f := range b.fields {
		names = append(names, f.name)
	}
	return fmt.Sprint(names)
}

func (b *FieldBuilders) ToComponent(info *ModelInfo, obj interface{}, ctx *web.EventContext) h.HTMLComponent {
	return b.ToComponentWithKeyPath(info, obj, "", ctx)
}

func (b *FieldBuilders) ToComponentWithKeyPath(info *ModelInfo, obj interface{}, keyPath string, ctx *web.EventContext) h.HTMLComponent {

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

		contextKeyPath := keyPath
		if f.name != SelfFieldName {
			contextKeyPath = fmt.Sprintf("%s.%s", keyPath, f.name)
		}

		comps = append(comps, f.compFunc(obj, &FieldContext{
			ModelInfo: info,
			Name:      f.name,
			KeyPath:   contextKeyPath,
			Label:     label,
			Errors:    vErr.GetFieldErrors(f.name),
			Context:   f.context,
		}, ctx))
	}

	return h.Components(comps...)
}

type RowFunc func(obj interface{}, content h.HTMLComponent, ctx *web.EventContext) h.HTMLComponent

func defaultRowFunc(obj interface{}, content h.HTMLComponent, ctx *web.EventContext) h.HTMLComponent {
	return content
}

func (b *FieldBuilders) ToComponentForEach(field *FieldContext, slice interface{}, ctx *web.EventContext, rowFunc RowFunc) h.HTMLComponent {
	var info *ModelInfo
	var parentKeyPath = ""
	if field != nil {
		info = field.ModelInfo
		parentKeyPath = field.KeyPath
	}
	if rowFunc == nil {
		rowFunc = defaultRowFunc
	}
	var r []h.HTMLComponent
	var i = 0

	funk.ForEach(slice, func(obj interface{}) {
		comps := b.ToComponentWithKeyPath(info, obj, fmt.Sprintf("%s[%d]", parentKeyPath, i), ctx)
		r = append(r, rowFunc(obj, comps, ctx))
		i++
	})

	return h.Components(r...)
}
