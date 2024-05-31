package perm

import (
	"path/filepath"

	"github.com/ory/ladon"
)

type PathMatcher struct{}

func (m *PathMatcher) Matches(p ladon.Policy, haystack []string, needle string) (bool, error) {
	for _, h := range haystack {
		m, err := filepath.Match(h, needle)
		if err != nil {
			return false, err
		}
		if m {
			return true, nil
		}
	}
	return false, nil
}
