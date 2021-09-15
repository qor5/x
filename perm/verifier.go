package perm

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ory/ladon"
)

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
	verbose bool
}

func Module(v string, b *Builder) (r *Verifier) {
	r = &Verifier{
		module: v,
	}

	if b == nil {
		return r
	}

	r.builder = b
	return
}

func (b *Verifier) Verbose(v bool) (r *Verifier) {
	b.verbose = v
	return b
}

func (b *Verifier) Do(v string) (r *Verifier) {
	if b.builder == nil {
		return b
	}

	r = &Verifier{
		module:  b.module,
		builder: b.builder,
		verbose: b.verbose,
	}

	r.vr = &verReq{
		resourcesParts: []string{b.module},
		req: &ladon.Request{
			Action: v,
		},
	}
	return
}

func (b *Verifier) On(vs ...string) (r *Verifier) {
	if b.builder == nil {
		return b
	}

	b.vr.resourcesParts = append(b.vr.resourcesParts, vs...)
	return b
}

func (b *Verifier) OnObject(v interface{}) (r *Verifier) {
	if b.builder == nil {
		return b
	}

	b.vr.objs = append(b.vr.objs, v)
	b.vr.resourcesParts = append(b.vr.resourcesParts, ToPermRN(v)...)
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

	b.vr.req.Resource = strings.Join(b.vr.resourcesParts, ":")

	if len(b.vr.subjects) == 0 && b.builder.subjectFunc != nil {
		b.vr.subjects = b.builder.subjectFunc(b.vr.r)
	}

	if len(b.vr.subjects) == 0 {
		b.vr.subjects = []string{"anonymous"}
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
		if b.verbose {
			fmt.Printf("permission req: %#+v\n", b.vr.req)
		}
		err = b.builder.ladon.IsAllowed(b.vr.req)
		if err == nil {
			return nil
		}
	}

	return err
}
