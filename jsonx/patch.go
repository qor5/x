package jsonx

import (
	"github.com/pkg/errors"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

func Patch(patch []byte, dest any) error {
	restore, err := Marshal(dest)
	if err != nil {
		return err
	}

	// In the patch, null values indicate field addition or overriding existing fields,
	// which doesn't conform to RFC7386 standard. However, it's more suitable for structs
	// since structs typically don't have the concept of field deletion.
	//
	// MergeMergePatches is a modified version of RFC7386, which is more suitable for structs.
	merged, err := jsonpatch.MergeMergePatches(restore, patch)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := Unmarshal(merged, dest); err != nil {
		return err
	}

	return nil
}

func Copy(dst any, src any) error {
	b, err := Marshal(src)
	if err != nil {
		return err
	}
	if err := Patch(b, dst); err != nil {
		return err
	}
	return nil
}
