package perm_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ory/ladon"
	"github.com/qor5/x/v3/perm"
	"github.com/sunfmin/reflectutils"
)

type Post struct {
	ID    uint
	Owner string
}

type MediaLibrary struct {
	ID       uint
	Category string
}

// RN -> Resource Name
func (m *MediaLibrary) PermissionRN() []string {
	return []string{"media_libraries", fmt.Sprint(m.ID), m.Category}
}

func getPost() *Post {
	return &Post{ID: 12, Owner: "user_123"}
}

func getMediaLibrary() *MediaLibrary {
	return &MediaLibrary{ID: 33, Category: "images"}
}

const (
	Create = "Create"
	Upload = "Upload"
)

func TestPermission(t *testing.T) {
	perm.Verbose = true

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var p *perm.Builder
			if !c.nilBuilder {
				p = perm.New().Policies(c.policies...).
					SubjectsFunc(sf(c.subjects...)).
					ContextFunc(c.contextFunc)
			}

			verifier := perm.NewVerifier("presets", p)

			hello := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				post := getPost()
				ml := getMediaLibrary()
				if verifier.Do(Upload).ObjectOn(post).On("heroImage").ObjectOn(ml).WithReq(r).IsAllowed() == nil {
					_, _ = fmt.Fprintln(w, "upload")
				}

				if verifier.Do(Create).ObjectOn(&Post{}).WithReq(r).IsAllowed() == nil {
					_, _ = fmt.Fprintln(w, "create")
				}
			})

			w := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", "/", nil)
			hello.ServeHTTP(w, r)

			if len(c.dontWantPermission) > 0 {
				if strings.Contains(w.Body.String(), c.dontWantPermission) {
					t.Errorf("%s should not have permission for %s, but was %s",
						c.subjects, c.dontWantPermission, w.Body.String())
				}
			}

			if len(c.wantPermission) > 0 {
				if !strings.Contains(w.Body.String(), c.wantPermission) {
					t.Errorf("%s should have permission for %s, but was %s",
						c.subjects, c.wantPermission, w.Body.String())
				}
			}
		})
	}
}

func sf(roles ...string) perm.SubjectsFunc {
	return func(r *http.Request) []string {
		return roles
	}
}

func ownerFunc(r *http.Request, objs []interface{}) perm.Context {
	if len(objs) > 0 {
		v, _ := reflectutils.Get(objs[0], "Owner")
		if v != nil {
			return perm.Context{
				"owner": v,
			}
		}
	}

	return nil
}

var cases = []struct {
	policies           []*perm.PolicyBuilder
	name               string
	subjects           []string
	dontWantPermission string
	wantPermission     string
	nilBuilder         bool
	contextFunc        perm.ContextFunc
}{
	{
		name: "anonymous should not have permission for upload on posts",
		policies: []*perm.PolicyBuilder{
			perm.PolicyFor("developer").WhoAre(perm.Allowed).ToDo(Upload).On("*:posts:*"),
		},
		subjects:           nil,
		dontWantPermission: "upload",
	},

	{
		name: "developer should have permission for upload on posts",
		policies: []*perm.PolicyBuilder{
			perm.PolicyFor("developer").WhoAre(perm.Allowed).ToDo(Upload).On("*:posts:*"),
		},
		subjects:       []string{"developer"},
		wantPermission: "upload",
	},

	{
		name: "developer should not have permission for upload on posts",
		policies: []*perm.PolicyBuilder{
			perm.PolicyFor("developer").WhoAre(perm.Allowed).ToDo(Upload).On("*:users:*"),
		},
		subjects:           []string{"developer"},
		dontWantPermission: "upload",
	},

	{
		name: "developer should have permission for upload on any posts media_libraries 33",
		policies: []*perm.PolicyBuilder{
			perm.PolicyFor("developer").WhoAre(perm.Allowed).ToDo(Upload).On("*media_libraries:33*"),
		},
		subjects:       []string{"developer"},
		wantPermission: "upload",
	},

	{
		name: "developer should have permission for upload on any posts media_libraries images category",
		policies: []*perm.PolicyBuilder{
			perm.PolicyFor("developer").WhoAre(perm.Allowed).ToDo(Upload).On("*media_libraries:*:images:"),
		},
		subjects:       []string{"developer"},
		wantPermission: "upload",
	},

	{
		name: "developer cant do anything",
		policies: []*perm.PolicyBuilder{
			perm.PolicyFor("developer").WhoAre(perm.Denied).ToDo(perm.Anything).On(perm.Anything),
		},
		subjects:           []string{"developer"},
		dontWantPermission: "upload",
	},

	{
		name: "developer can do anything",
		policies: []*perm.PolicyBuilder{
			perm.PolicyFor("developer").WhoAre(perm.Allowed).ToDo(perm.Anything).On(perm.Anything),
		},
		subjects:       []string{"developer"},
		wantPermission: "upload",
	},

	{
		name: "any body can do anything if they are owner",
		policies: []*perm.PolicyBuilder{
			perm.PolicyFor(perm.Anybody).
				WhoAre(perm.Allowed).ToDo(perm.Anything).On(perm.Anything).Given(
				perm.Conditions{
					"owner": &ladon.EqualsSubjectCondition{},
				},
			),
		},
		subjects:       []string{"developer", "user_123"},
		wantPermission: "upload",
		contextFunc:    ownerFunc,
	},

	{
		name: "any body cant do anything if they are not owner",
		policies: []*perm.PolicyBuilder{
			perm.PolicyFor(perm.Anybody).
				WhoAre(perm.Allowed).ToDo(perm.Anything).On(perm.Anything).Given(
				perm.Conditions{
					"owner": &ladon.EqualsSubjectCondition{},
				},
			),
		},
		subjects:           []string{"developer", "user_not_owner"},
		dontWantPermission: "upload",
		contextFunc:        ownerFunc,
	},

	{
		name:           "nil builder should allow to do everything",
		nilBuilder:     true,
		subjects:       []string{"developer"},
		wantPermission: "upload",
	},

	{
		name:               "empty policies should not allowed to do everything",
		policies:           []*perm.PolicyBuilder{},
		subjects:           []string{"developer"},
		dontWantPermission: "upload",
	},
}
