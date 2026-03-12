package httperrors

import (
	"testing"
)

func TestAssertFieldViolations(t *testing.T) {
	t.Run("passes with matching violations", func(t *testing.T) {
		fv1 := NewFieldViolation("email", "REQUIRED", "email is required")
		fv2 := NewFieldViolation("name", "TOO_SHORT", "name is too short")
		err := BadRequest(fv1, fv2).Err()

		// Should not fail
		AssertFieldViolations(t, err,
			NewFieldViolation("email", "REQUIRED", "email is required"),
			NewFieldViolation("name", "TOO_SHORT", "name is too short"),
		)
	})

	t.Run("passes with subset of violations", func(t *testing.T) {
		fv1 := NewFieldViolation("email", "REQUIRED", "email is required")
		fv2 := NewFieldViolation("name", "TOO_SHORT", "name is too short")
		err := BadRequest(fv1, fv2).Err()

		// Checking only one violation should still pass
		AssertFieldViolations(t, err,
			NewFieldViolation("email", "REQUIRED", "email is required"),
		)
	})
}
