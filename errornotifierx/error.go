package errornotifierx

import "github.com/pkg/errors"

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type trackedError struct {
	errorMsg string
	tracer   stackTracer
}

// newTrackedError creates a tracked error with deepest stack tracer
func newTrackedError(err error) error {
	if err == nil {
		return nil
	}

	tracerErr := findDeepestStackTracer(err)
	if tracerErr == nil {
		return err //nolint:errhandle
	}

	return &trackedError{ //nolint:errhandle
		errorMsg: err.Error(),
		tracer:   tracerErr,
	}
}

func (e *trackedError) Error() string {
	return e.errorMsg
}

func (e *trackedError) StackTrace() errors.StackTrace {
	return e.tracer.StackTrace()
}

func findDeepestStackTracer(err error) stackTracer {
	if err == nil {
		return nil
	}

	var deepest stackTracer
	current := err

	for current != nil {
		if tracer, ok := current.(stackTracer); ok {
			deepest = tracer
		}

		if unwrapper, ok := current.(interface{ Unwrap() error }); ok {
			current = unwrapper.Unwrap()
		} else {
			break
		}
	}

	return deepest
}
