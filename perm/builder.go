package perm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
	"github.com/samber/lo"
	"github.com/sunfmin/reflectutils"
	"gorm.io/gorm"
)

const (
	Allowed   = ladon.AllowAccess
	Denied    = ladon.DenyAccess
	Anything  = "*"
	Anybody   = "*"
	Anonymous = "anonymous"
)

var PermissionDenied = errors.New("permission denied")

type (
	Context      = ladon.Context
	Conditions   = ladon.Conditions
	SubjectsFunc func(r *http.Request) []string
	ContextFunc  func(r *http.Request, objs []interface{}) Context
)

type DBPolicy interface {
	LoadDBPolicies(db *gorm.DB, startFrom *time.Time) ([]*PolicyBuilder, []*PolicyBuilder)
}

type permRNer interface {
	PermissionRN() []string
}

func ToPermissionRN(v interface{}) []string {
	if pr, ok := v.(permRNer); ok {
		return pr.PermissionRN()
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
	m            sync.Mutex
	policies     []*PolicyBuilder
	ladon        *ladon.Ladon
	subjectsFunc SubjectsFunc
	contextFunc  ContextFunc
	dbPolicy     *DBPolicyBuilder
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
	b.DeletePolicies(b.policies...)
	b.policies = make([]*PolicyBuilder, 0)
	b.CreatePolicies(ps...)
	return b
}

func (b *Builder) createPolicy(p *PolicyBuilder) {
	p.setIDIfEmpty()

	// just ignore if duplicate
	if lo.ContainsBy(b.policies, func(item *PolicyBuilder) bool {
		return item.GetID() == p.GetID()
	}) {
		return
	}

	err := b.ladon.Manager.Create(context.TODO(), p.policy)
	if err != nil {
		panic(err)
	}
	b.policies = append(b.policies, p)
}

func (b *Builder) updatePolicy(p *PolicyBuilder) {
	i := b.findPolicyIndex(p.GetID())
	if i < 0 {
		return
	}

	err := b.ladon.Manager.Update(context.TODO(), p.policy)
	if err != nil {
		panic(err)
	}
	b.policies[i] = p
}

func (b *Builder) findPolicyIndex(id string) int {
	for i, p := range b.policies {
		if p.GetID() == id {
			return i
		}
	}
	return -1
}

func (b *Builder) updateOrCreatePolicy(p *PolicyBuilder) {
	i := b.findPolicyIndex(p.GetID())
	if i < 0 {
		b.createPolicy(p)
	} else {
		b.updatePolicy(p)
	}
}

func (b *Builder) deletePolicy(p *PolicyBuilder) {
	for i, ep := range b.policies {
		if ep.GetID() == p.GetID() {
			err := b.ladon.Manager.Delete(context.TODO(), p.GetID())
			if err != nil {
				panic(err)
			}
			b.policies = append(b.policies[:i], b.policies[i+1:]...)
			break
		}
	}
}

func (b *Builder) SubjectsFunc(v SubjectsFunc) (r *Builder) {
	b.subjectsFunc = v
	return b
}

func (b *Builder) GetSubjectsFunc() SubjectsFunc {
	return b.subjectsFunc
}

func (b *Builder) ContextFunc(v ContextFunc) (r *Builder) {
	b.contextFunc = v
	return b
}

func (b *Builder) GetContextFunc() ContextFunc {
	return b.contextFunc
}

func (b *Builder) DBPolicy(dpb *DBPolicyBuilder) (r *Builder) {
	b.dbPolicy = dpb

	go b.loopLoadDBPolicies(dpb.db, dpb.loadFrequency)
	return b
}

func (b *Builder) CreatePolicies(ps ...*PolicyBuilder) {
	b.m.Lock()
	defer b.m.Unlock()
	for _, p := range ps {
		b.createPolicy(p)
	}
}

func (b *Builder) UpdatePolicies(toUpdate ...*PolicyBuilder) {
	b.m.Lock()
	defer b.m.Unlock()
	for _, p := range toUpdate {
		b.updatePolicy(p)
	}
}

func (b *Builder) UpdateOrCreatePolicies(toUpdate ...*PolicyBuilder) {
	b.m.Lock()
	defer b.m.Unlock()
	for _, p := range toUpdate {
		b.updateOrCreatePolicy(p)
	}
}

func (b *Builder) DeletePolicies(toDelete ...*PolicyBuilder) {
	b.m.Lock()
	defer b.m.Unlock()
	for _, p := range toDelete {
		b.deletePolicy(p)
	}
}

func (b *Builder) LoadDBPoliciesToMemory(db *gorm.DB, startFrom *time.Time) {
	toUpdateOrCreate, toDelete := b.dbPolicy.model.LoadDBPolicies(db, startFrom)
	b.DeletePolicies(toDelete...)
	b.UpdateOrCreatePolicies(toUpdateOrCreate...)
	if Verbose {
		b.printPolices()
	}
}

func (b *Builder) loopLoadDBPolicies(db *gorm.DB, duration time.Duration) {
	now := time.Now()
	b.LoadDBPoliciesToMemory(db, nil)

	for next := range time.Tick(duration) {
		b.LoadDBPoliciesToMemory(db, &now)
		now = next
	}
}

func (b *Builder) printPolices() {
	allp, _ := b.ladon.Manager.GetAll(context.TODO(), 100, 0)
	fmt.Printf("all permission policies: \n")
	for _, p := range allp {
		fmt.Printf("%+v \n", p)
	}
}
