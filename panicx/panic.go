package panicx

import "github.com/pkg/errors"

func On(f func(perr error)) {
	var perr error
	if r := recover(); r != nil {
		defer panic(r)
		if e, ok := r.(error); ok {
			perr = e
		} else {
			perr = errors.Errorf("panic: %v", r)
		}
	}
	f(perr)
}
