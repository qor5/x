package presets

import (
	"fmt"
	"log"

	"github.com/sunfmin/bran/web"

	"github.com/sunfmin/reflectutils"

	v "github.com/sunfmin/bran/x/vuetify"
	h "github.com/theplant/htmlgo"
)

type FieldBuilder struct {
	NameLabel
	compFunc FieldComponentFunc
	//setterFunc SetterFunc
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
	//r.setterFunc = b.setterFunc
	return r
}

func (b *FieldBuilder) ComponentFunc(v FieldComponentFunc) (r *FieldBuilder) {
	if v == nil {
		panic("value required")
	}
	b.compFunc = v
	return b
}

//
//func (b *FieldBuilder) SetterFunc(v SetterFunc) (r *FieldBuilder) {
//	b.setterFunc = v
//	return b
//}

type NameLabel struct {
	name  string
	label string
}

type FieldBuilders struct {
	obj         interface{}
	defaults    *FieldDefaults
	fieldLabels []string
	fields      []*FieldBuilder
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

	r = &FieldBuilders{fieldLabels: b.fieldLabels}

	for _, n := range names {
		f := b.GetField(n)
		if f == nil {
			fType := reflectutils.GetType(b.obj, n)
			if fType == nil {
				continue
			}

			compFunc := b.defaults.fieldTypeByType(fType).compFunc
			if compFunc != nil {
				r.Field(n).ComponentFunc(compFunc)
				continue
			}
		}
		r.fields = append(r.fields, f.Clone())
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

func (b *FieldBuilders) MustSet(obj interface{}, newObj interface{}) {
	for _, f := range b.fields {
		err := reflectutils.Set(obj, f.name, reflectutils.MustGet(newObj, f.name))
		if err != nil {
			panic(err)
		}
	}
}

func (b *FieldBuilders) String() (r string) {
	var names []string
	for _, f := range b.fields {
		names = append(names, f.name)
	}
	return fmt.Sprint(names)
}

func (b *FieldBuilders) ToComponent(obj interface{}, verr *ValidationErrors, ctx *web.EventContext) h.HTMLComponent {

	var comps []h.HTMLComponent

	if verr == nil {
		verr = &ValidationErrors{}
	}

	gErr := verr.GetGlobalError()
	if len(gErr) > 0 {
		comps = append(
			comps,
			v.VAlert(h.Text(gErr)).
				Border("left").
				Type("error").
				Elevation(2).
				ColoredBorder(true),
		)
	}

	for _, f := range b.fields {
		if f.compFunc == nil {
			continue
		}

		comps = append(comps, f.compFunc(obj, &FieldContext{
			Name:   f.name,
			Label:  b.getLabel(f.NameLabel),
			Errors: verr.GetFieldErrors(f.name),
		}, ctx))
	}

	return h.Components(comps...)
}

type ValidationErrors struct {
	globalErrors []string
	fieldErrors  map[string][]string
}

func (b *ValidationErrors) FieldError(fieldName string, message string) (r *ValidationErrors) {
	if b.fieldErrors == nil {
		b.fieldErrors = make(map[string][]string)
	}
	b.fieldErrors[fieldName] = append(b.fieldErrors[fieldName], message)
	return b
}

func (b *ValidationErrors) GlobalError(message string) (r *ValidationErrors) {
	b.globalErrors = append(b.globalErrors, message)
	return b
}

func (b *ValidationErrors) GetFieldErrors(fieldName string) (r []string) {
	if b.fieldErrors == nil {
		return
	}

	r = b.fieldErrors[fieldName]
	return
}

func (b *ValidationErrors) GetGlobalError() (r string) {
	if len(b.globalErrors) == 0 {
		return
	}
	return b.globalErrors[0]
}

func (b *ValidationErrors) GetGlobalErrors() (r []string) {
	return b.globalErrors
}

func (b *ValidationErrors) HaveErrors() bool {
	if len(b.globalErrors) > 0 {
		return true
	}
	if len(b.fieldErrors) > 0 {
		return true
	}
	return false
}

func (b *ValidationErrors) Error() string {
	return fmt.Sprintf("validation error global: %+v, fields: %+v", b.globalErrors, b.fieldErrors)
}
