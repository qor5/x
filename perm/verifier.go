package perm

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/ory/ladon"
)

var Verbose = false

type verReq struct {
	subjects       []string
	objs           []interface{}
	r              *http.Request
	req            *ladon.Request
	resourcesParts []string
}

type Verifier struct {
	builder *Builder
	module  string
	vr      *verReq
}

func NewVerifier(module string, b *Builder) (r *Verifier) {
	r = &Verifier{
		module: module,
	}

	if b == nil {
		return r
	}

	r.builder = b
	return
}

func (b *Verifier) Spawn() (r *Verifier) {
	if b.builder == nil {
		return b
	}

	r = &Verifier{
		module:  b.module,
		builder: b.builder,
	}

	resourceParts := []string{b.module}
	if b.vr != nil {
		resourceParts = b.vr.resourcesParts
	}

	r.vr = &verReq{
		resourcesParts: append([]string{}, resourceParts...),
		req:            &ladon.Request{},
	}

	return
}

func (b *Verifier) Do(v string) (r *Verifier) {
	if b.builder == nil {
		return b
	}

	r = b.Spawn()
	r.vr.req.Action = v
	return
}

// SnakeDo convert string to snake form.
// e.g. "SnakeDo" -> "snake_do"
func (b *Verifier) SnakeDo(actions ...string) (r *Verifier) {
	fixed := []string{b.module}
	for _, a := range actions {
		fixed = append(fixed, strcase.ToSnake(a))
	}
	return b.Do(strings.Join(fixed, ":"))
}

func (b *Verifier) On(vs ...string) (r *Verifier) {
	if b.builder == nil {
		return b
	}

	b.vr.resourcesParts = append(b.vr.resourcesParts, vs...)
	return b
}

func (b *Verifier) SnakeOn(vs ...string) (r *Verifier) {
	if b.builder == nil {
		return b
	}

	var fixed []string
	for _, v := range vs {
		if v == "" {
			continue
		}
		fixed = append(fixed, strcase.ToSnake(v))
	}

	b.On(fixed...)
	return b
}

func (b *Verifier) ObjectOn(v interface{}) (r *Verifier) {
	if b.builder == nil {
		return b
	}

	b.vr.objs = append(b.vr.objs, v)
	b.vr.resourcesParts = append(b.vr.resourcesParts, ToPermissionRN(v)...)
	return b
}

func (b *Verifier) RemoveOn(length int) (r *Verifier) {
	if b.builder == nil {
		return b
	}
	if len(b.vr.resourcesParts) >= length {
		b.vr.resourcesParts = b.vr.resourcesParts[:len(b.vr.resourcesParts)-length]
	}
	return b
}

func (b *Verifier) WithReq(v *http.Request) (r *Verifier) {
	if b.builder == nil {
		return b
	}
	b.vr.r = v
	return b
}

func (b *Verifier) From(v string) (r *Verifier) {
	if b.builder == nil {
		return b
	}

	b.vr.subjects = append(b.vr.subjects, v)
	return b
}

func (b *Verifier) Given(v ladon.Context) (r *Verifier) {
	if b.builder == nil {
		return b
	}
	b.vr.req.Context = v
	return b
}

func (b *Verifier) IsAllowed() error {
	if b.builder == nil {
		return nil
	}

	b.vr.req.Resource = ":" + strings.Join(b.vr.resourcesParts, ":") + ":"

	if len(b.vr.subjects) == 0 && b.builder.subjectsFunc != nil {
		b.vr.subjects = b.builder.subjectsFunc(b.vr.r)
	}

	if len(b.vr.subjects) == 0 {
		b.vr.subjects = []string{Anonymous}
	}

	if b.builder.contextFunc != nil {
		newContext := b.builder.contextFunc(b.vr.r, b.vr.objs)
		if newContext != nil {
			for k, v := range b.vr.req.Context {
				newContext[k] = v
			}
			b.vr.req.Context = newContext
		}
	}

	var err error
	// any of the subjects have permission, then have permission
	for _, sub := range b.vr.subjects {
		b.vr.req.Subject = sub

		err = b.builder.ladon.IsAllowed(context.TODO(), b.vr.req)
		if Verbose {
			fmt.Printf("have permission: %+v, req: %#+v\n", err == nil, b.vr.req)
		}
		if err == nil {
			return nil
		}
	}

	return err
}
