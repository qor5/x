package presets

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/goplaid/web"
	"github.com/goplaid/x/perm"
	"github.com/goplaid/x/presets/actions"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

type ModelBuilder struct {
	p             *Builder
	model         interface{}
	primaryField  string
	modelType     reflect.Type
	inGroup       bool
	menuGroupName string
	notInMenu     bool
	menuIcon      string
	uriName       string
	label         string
	fieldLabels   []string
	placeholders  []string
	listing       *ListingBuilder
	detailing     *DetailingBuilder
	editing       *EditingBuilder
	creating      *EditingBuilder
	writeFields   *FieldBuilders
	hasDetailing  bool
}

func NewModelBuilder(p *Builder, model interface{}) (r *ModelBuilder) {
	r = &ModelBuilder{p: p, model: model, primaryField: "ID"}
	r.modelType = reflect.TypeOf(model)
	if r.modelType.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("model %#+v must be pointer", model))
	}
	modelstr := r.modelType.String()
	modelName := modelstr[strings.LastIndex(modelstr, ".")+1:]
	r.label = strcase.ToCamel(inflection.Plural(modelName))
	r.uriName = inflection.Plural(strcase.ToKebab(modelName))

	r.newListing()
	r.newDetailing()
	r.newEditing()

	return
}

func (b *ModelBuilder) ensureEventFuncs(hub web.EventFuncHub) {
	hub.RegisterEventFunc(actions.DrawerNew, b.editing.formDrawerNew)
	hub.RegisterEventFunc(actions.DrawerEdit, b.editing.formDrawerEdit)
	hub.RegisterEventFunc(actions.DeleteConfirmation, b.listing.deleteConfirmation)
	hub.RegisterEventFunc(actions.Update, b.editing.defaultUpdate)
	hub.RegisterEventFunc(actions.DoDelete, b.editing.doDelete)
	hub.RegisterEventFunc(actions.DoBulkAction, b.listing.doBulkAction)
	hub.RegisterEventFunc(actions.DrawerAction, b.detailing.formDrawerAction)
	hub.RegisterEventFunc(actions.DoAction, b.detailing.doAction)
}

func (b *ModelBuilder) newModel() (r interface{}) {
	return reflect.New(b.modelType.Elem()).Interface()
}

func (b *ModelBuilder) newModelArray() (r interface{}) {
	return reflect.New(reflect.SliceOf(b.modelType)).Interface()
}

func (b *ModelBuilder) newListing() (r *ListingBuilder) {
	b.listing = &ListingBuilder{mb: b, FieldBuilders: *b.p.listFieldDefaults.InspectFields(b.model)}
	if b.p.dataOperator != nil {
		b.listing.Searcher(b.p.dataOperator.Search)
	}
	return
}

func (b *ModelBuilder) newEditing() (r *EditingBuilder) {
	b.writeFields, b.listing.searchColumns = b.p.writeFieldDefaults.inspectFieldsAndCollectName(b.model, reflect.TypeOf(""))
	b.editing = &EditingBuilder{mb: b, FieldBuilders: *b.writeFields}
	if b.p.dataOperator != nil {
		b.editing.FetchFunc(b.p.dataOperator.Fetch)
		b.editing.SaveFunc(b.p.dataOperator.Save)
		b.editing.DeleteFunc(b.p.dataOperator.Delete)
	}
	return
}

func (b *ModelBuilder) newDetailing() (r *DetailingBuilder) {
	b.detailing = &DetailingBuilder{mb: b, FieldBuilders: *b.p.detailFieldDefaults.InspectFields(b.model)}
	if b.p.dataOperator != nil {
		b.detailing.Fetcher(b.p.dataOperator.Fetch)
	}
	return
}

func (b *ModelBuilder) Info() (r *ModelInfo) {
	mi := ModelInfo(*b)
	return &mi
}

type ModelInfo ModelBuilder

func (b *ModelInfo) ListingHref() string {
	return fmt.Sprintf("%s/%s", b.p.prefix, b.uriName)
}

func (b *ModelInfo) EditingHref(id string) string {
	return fmt.Sprintf("%s/%s/%s/edit", b.p.prefix, b.uriName, id)
}

func (b *ModelInfo) DetailingHref(id string) string {
	return fmt.Sprintf("%s/%s/%s", b.p.prefix, b.uriName, id)
}

func (b *ModelInfo) HasDetailing() bool {
	return b.hasDetailing
}

func (b *ModelInfo) PresetsPrefix() string {
	return b.p.prefix
}

func (b *ModelInfo) URIName() string {
	return b.uriName
}

func (b *ModelInfo) Verifier() *perm.Verifier {
	return b.p.verifier.Spawn().
		SnakeOn(b.menuGroupName).
		SnakeOn(b.uriName)
}

func (b *ModelBuilder) URIName(v string) (r *ModelBuilder) {
	b.uriName = v
	return b
}

func (b *ModelBuilder) PrimaryField(v string) (r *ModelBuilder) {
	b.primaryField = v
	return b
}

func (b *ModelBuilder) MenuGroup(v string) (r *ModelBuilder) {
	b.p.MenuGroup(v).AppendModels(b)
	b.menuGroupName = v
	b.inGroup = true
	return b
}

func (b *ModelBuilder) InMenu(v bool) (r *ModelBuilder) {
	b.notInMenu = !v
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

func (b *ModelBuilder) getComponentFuncField(field *FieldBuilder) (r *FieldContext) {
	r = &FieldContext{
		ModelInfo: b.Info(),
		Name:      field.name,
		Label:     b.getLabel(field.NameLabel),
	}
	return
}

func (b *ModelBuilder) getLabel(field NameLabel) (r string) {
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
