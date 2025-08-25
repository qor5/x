package jsonx

import (
	"bytes"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

func Equal(a, b any) bool {
	aJSON, aErr := Marshal(a)
	bJSON, bErr := Marshal(b)
	if aErr != nil || bErr != nil {
		return false
	}
	switch aJSON[0] {
	case '[':
		return bJSON[0] == '[' && jsonpatch.Equal(aJSON, bJSON)
	case '{':
		return bJSON[0] == '{' && jsonpatch.Equal(aJSON, bJSON)
	default:
		return bytes.Equal(aJSON, bJSON)
	}
}
