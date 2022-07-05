package perm

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
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

type Context = ladon.Context
type Conditions = ladon.Conditions
type SubjectsFunc func(r *http.Request) []string
type ContextFunc func(r *http.Request, objs []interface{}) Context

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
	policies      []*PolicyBuilder
	ladon         *ladon.Ladon
	subjectsFunc  SubjectsFunc
	contextFunc   ContextFunc
	dbPolicyModel DBPolicy
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
	for _, p := range b.policies {
		if p.policy.ID == "" {
			p.policy.ID = fmt.Sprintf("%x", md5.Sum(p.Json()))
		}
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

func (b *Builder) EnableDBPolicy(db *gorm.DB, dbPolicyModel DBPolicy, loadDuration time.Duration) (r *Builder) {
	if dbPolicyModel == nil {
		b.dbPolicyModel = DefaultDBPolicy{}
	} else {
		b.dbPolicyModel = dbPolicyModel
	}

	go b.loopLoadDBPolicies(db, loadDuration)
	return b
}

func (b *Builder) UpdatePolicies(toUpdate ...*PolicyBuilder) {
	for _, p := range toUpdate {
		b.ladon.Manager.Update(p.policy)
	}
}

func (b *Builder) DeletePolicies(toDelete ...*PolicyBuilder) {
	for _, p := range toDelete {
		b.ladon.Manager.Delete(p.GetID())
	}
}

func (b *Builder) LoadDBPoliciesToMemory(db *gorm.DB, startFrom *time.Time) {
	toUpdate, toDelete := b.dbPolicyModel.LoadDBPolicies(db, startFrom)
	b.DeletePolicies(toDelete...)
	b.UpdatePolicies(toUpdate...)
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
	allp, _ := b.ladon.Manager.GetAll(100, 0)
	for _, p := range allp {
		fmt.Printf("%+v \n", p)
	}
}
