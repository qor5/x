package perm

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
	"github.com/sunfmin/reflectutils"
)

const (
	Allowed          = ladon.AllowAccess
	Denied           = ladon.DenyAccess
	Anything         = "*"
	Anybody          = "*"
	Anonymous        = "anonymous"
	PermissionDenied = "Permission Denied"
)

type Context = ladon.Context
type Conditions = ladon.Conditions
type SubjectsFunc func(r *http.Request) []string
type ContextFunc func(r *http.Request, objs []interface{}) Context

type permRNer interface {
	PermRN() []string
}

func ToPermRN(v interface{}) []string {
	if pr, ok := v.(permRNer); ok {
		return pr.PermRN()
	}
	typeName := fmt.Sprint(reflect.TypeOf(v))
	ss := strings.Split(typeName, ".")
	typeName = ss[len(ss)-1]
	typeName = strings.NewReplacer("*", "", ".", "-").Replace(typeName)
	typeName = strcase.ToSnake(inflection.Plural(typeName))
	id, err := reflectutils.Get(v, "ID")
	if err == nil && len(fmt.Sprint(id)) > 0 && fmt.Sprint(id) != "0" {
		return []string{typeName, fmt.Sprint(id)}
	}
	return []string{typeName}
}

type Builder struct {
	policies     []*PolicyBuilder
	ladon        *ladon.Ladon
	subjectsFunc SubjectsFunc
	contextFunc  ContextFunc
}

func New() *Builder {
	return &Builder{
		ladon: &ladon.Ladon{
			Manager: memory.NewMemoryManager(),
			Matcher: &PathMatcher{},
		},
	}
}

func (b *Builder) Policies(ps ...*PolicyBuilder) (r *Builder) {
	b.policies = ps
	for i, p := range b.policies {
		p.policy.ID = fmt.Sprint(i)
		err := b.ladon.Manager.Create(p.policy)
		if err != nil {
			panic(err)
		}
	}
	return b
}

func (b *Builder) SubjectsFunc(v SubjectsFunc) (r *Builder) {
	b.subjectsFunc = v
	return b
}

func (b *Builder) ContextFunc(v ContextFunc) (r *Builder) {
	b.contextFunc = v
	return b
}

type PolicyBuilder struct {
	policy *ladon.DefaultPolicy
	module string
}

func NewPolicy() *PolicyBuilder {
	return &PolicyBuilder{
		policy: &ladon.DefaultPolicy{},
	}
}

func (b *PolicyBuilder) Module(module string) (r *PolicyBuilder) {
	b.module = module
	return b
}

func (b *PolicyBuilder) They(subjects ...string) (r *PolicyBuilder) {
	b.policy.Subjects = subjects
	return b
}

func (b *PolicyBuilder) Are(effect string) (r *PolicyBuilder) {
	b.policy.Effect = effect
	return b
}

func (b *PolicyBuilder) ToDo(actions ...string) (r *PolicyBuilder) {
	b.policy.Actions = actions
	return b
}

func (b *PolicyBuilder) On(resources ...string) (r *PolicyBuilder) {
	if b.module == "" {
		b.policy.Resources = append(b.policy.Resources, resources...)
		return
	}

	var newRes []string
	for _, res := range resources {
		newRes = append(newRes, strings.Join([]string{b.module, res}, ":"))
	}
	b.policy.Resources = append(b.policy.Resources, newRes...)
	return b
}

func (b *PolicyBuilder) Given(conditions Conditions) (r *PolicyBuilder) {
	b.policy.Conditions = conditions
	return b
}
