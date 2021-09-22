package perm

import (
	"strings"

	"github.com/ory/ladon"
)

type PolicyBuilder struct {
	policy *ladon.DefaultPolicy
	module string
}

func PolicyFor(subjects ...string) *PolicyBuilder {
	return &PolicyBuilder{
		policy: &ladon.DefaultPolicy{
			Subjects: subjects,
		},
	}
}

func (b *PolicyBuilder) Module(module string) (r *PolicyBuilder) {
	b.module = module
	return b
}

func (b *PolicyBuilder) WhoAre(effect string) (r *PolicyBuilder) {
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
		return b
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
