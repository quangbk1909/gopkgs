package errs

import (
	"errors"
	"fmt"

	"gitlab.id.vin/platform/gopkgs/internal/stacks"
)

type Debugger interface {
	StackTrace() string
	Caller() string
}

func Cause(err error) error {
	inside := errors.Unwrap(err)
	for inside != nil {
		err = inside
		inside = errors.Unwrap(err)
	}
	return err
}

func Wrap(err error, msg string) error {
	var d Debugger
	if !errors.As(err, &d) {
		s, c := stacks.DebugInformation(1)
		err = &errorStack{
			error:      err,
			stackTrace: s,
			caller:     c,
		}
	}
	return fmt.Errorf("%s: %w", msg, err)
}

func WithStack(err error) error {
	s, c := stacks.DebugInformation(1)
	return &errorStack{
		error:      err,
		stackTrace: s,
		caller:     c,
	}
}

type errorStack struct {
	error
	stackTrace string
	caller     string
}

func (e *errorStack) StackTrace() string {
	return e.stackTrace
}

func (e *errorStack) Caller() string {
	return e.caller
}

func (e *errorStack) Unwrap() error {
	return e.error
}
