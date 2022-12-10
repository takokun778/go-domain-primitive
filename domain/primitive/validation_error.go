package primitive

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	msg string
	err error
}

func NewValidationError(msg string, err error) ValidationError {
	return ValidationError{
		msg: msg,
		err: err,
	}
}

func (e ValidationError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.err)
	}

	return e.msg
}

func AsValidationError(err error) bool {
	var target ValidationError

	return errors.As(err, &target)
}
