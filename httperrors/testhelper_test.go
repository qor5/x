package httperrors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertFieldViolations(t *testing.T) {
	t.Run("passes with matching violations", func(t *testing.T) {
		fv1 := NewFieldViolation("email", "REQUIRED", "email is required")
		fv2 := NewFieldViolation("name", "TOO_SHORT", "name is too short")
		err := ValidationError(fv1, fv2).Err()

		// Should not fail
		AssertFieldViolations(t, err,
			NewFieldViolation("email", "REQUIRED", "email is required"),
			NewFieldViolation("name", "TOO_SHORT", "name is too short"),
		)
	})

	t.Run("passes with subset of violations", func(t *testing.T) {
		fv1 := NewFieldViolation("email", "REQUIRED", "email is required")
		fv2 := NewFieldViolation("name", "TOO_SHORT", "name is too short")
		err := ValidationError(fv1, fv2).Err()

		// Checking only one violation should still pass
		AssertFieldViolations(t, err,
			NewFieldViolation("email", "REQUIRED", "email is required"),
		)
	})

	t.Run("fails on wrong status code", func(t *testing.T) {
		mockT := &testing.T{}
		err := Error(500, "INTERNAL", "internal error")
		AssertFieldViolations(mockT, err)
		assert.True(t, mockT.Failed())
	})
}
