package presets

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/inflection"

	"github.com/iancoleman/strcase"
)

type ModelBuilder struct {
	p            *Builder
	model        interface{}
	modelType    reflect.Type
	inGroup      bool
	menuIcon     string
	uriName      string
	label        string
	fieldLabels  []string
	placeholders []string
	listing      *ListingBuilder
	editing      *EditingBuilder
	detailing    *DetailingBuilder
}

func NewModelBuilder(p *Builder, model interface{}) (r *ModelBuilder) {
	r = &ModelBuilder{p: p, model: model}
	r.modelType = reflect.TypeOf(model)
	if r.modelType.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("model %#+v must be pointer", model))
	}
	modelstr := r.modelType.String()
	modelName := modelstr[strings.LastIndex(modelstr, ".")+1:]
	r.label = strcase.ToCamel(inflection.Plural(modelName))
	r.uriName = strcase.ToKebab(modelName)
	r.newListing()
	r.newDetailing()
	r.newEditing()
	r.inspectModel()
	return
}

func (b *ModelBuilder) newModel() (r interface{}) {
	return reflect.New(b.modelType.Elem()).Interface()
}

func (b *ModelBuilder) newModelArray() (r interface{}) {
	return reflect.New(reflect.SliceOf(b.modelType)).Interface()
}

func (b *ModelBuilder) inspectModel() {
	v := reflect.ValueOf(b.model)

	for v.Elem().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v = v.Elem()

	t := v.Type()

	var sc []string
	var stringType = reflect.TypeOf("")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		//fmt.Println(f.Name, f.Type)
		ft := b.p.fieldTypeByType(f.Type)
		if !b.p.fieldNameExcluded(LISTING, f.Name) {
			b.listing.Field(f.Name).ComponentFunc(ft.listingCompFunc)
		}
		if !b.p.fieldNameExcluded(DETAILING, f.Name) {
			b.detailing.Field(f.Name).ComponentFunc(ft.detailingCompFunc)
		}
		if !b.p.fieldNameExcluded(EDITING, f.Name) {
			b.editing.Field(f.Name).ComponentFunc(ft.editingCompFunc)
		}
		if f.Type == stringType {
			sc = append(sc, strcase.ToSnake(f.Name))
		}
	}
	b.listing.searchColumns = sc
}

func (b *ModelBuilder) newListing() (r *ListingBuilder) {
	b.listing = &ListingBuilder{mb: b}
	if b.p.dataOperator != nil {
		b.listing.Searcher(b.p.dataOperator.Search)
	}
	return
}

func (b *ModelBuilder) newEditing() (r *EditingBuilder) {
	b.editing = &EditingBuilder{mb: b}
	if b.p.dataOperator != nil {
		b.editing.Fetcher(b.p.dataOperator.Fetch)
		b.editing.Saver(b.p.dataOperator.Save)
	}
	return
}

func (b *ModelBuilder) newDetailing() (r *DetailingBuilder) {
	b.detailing = &DetailingBuilder{mb: b}
	if b.p.dataOperator != nil {
		b.detailing.Fetcher(b.p.dataOperator.Fetch)
	}
	return
}

func (b *ModelBuilder) listingHref() string {
	muri := inflection.Plural(b.uriName)
	return fmt.Sprintf("%s/%s", b.p.prefix, muri)
}

func (b *ModelBuilder) detailingHref(id string) string {
	muri := inflection.Plural(b.uriName)
	return fmt.Sprintf("%s/%s/%s", b.p.prefix, muri, id)
}

func (b *ModelBuilder) URIName(v string) (r *ModelBuilder) {
	b.uriName = v
	return b
}

func (b *ModelBuilder) MenuGroup(v string) (r *ModelBuilder) {
	b.p.MenuGroup(v).AppendModels(b)
	b.inGroup = true
	return b
}

func (b *ModelBuilder) MenuIcon(v string) (r *ModelBuilder) {
	b.menuIcon = v
	return b
}

func (b *ModelBuilder) Label(v string) (r *ModelBuilder) {
	b.label = v
	return b
}

func (b *ModelBuilder) Labels(vs ...string) (r *ModelBuilder) {
	b.fieldLabels = append(b.fieldLabels, vs...)
	return b
}

func (b *ModelBuilder) Placeholders(vs ...string) (r *ModelBuilder) {
	b.placeholders = append(b.placeholders, vs...)
	return b
}

func (b *ModelBuilder) getComponentFuncField(field *FieldBuilder) (r *Field) {
	r = &Field{
		Name:  field.name,
		Label: b.getLabel(field),
	}
	return
}

func (b *ModelBuilder) getLabel(field *FieldBuilder) (r string) {
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
