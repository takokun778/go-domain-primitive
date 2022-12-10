package sample

import (
	"go-domain-primitive/domain/primitive"
)

// Name for call New() Reconstruct() Validate().
var Name name //nolint:gochecknoglobals

// name domain primitive struct.
type name struct {
	value string
}

// New domain primitibe construct function.
func (name) New(v string) (primitive.Plimitive[string], error) {
	if err := Name.Validate(v); err != nil {
		return name{}, err
	}

	return name{value: v}, nil
}

// Reconstruct reconstruct domain primitive function.
func (name) Reconstruct(v string) primitive.Plimitive[string] {
	return name{value: v}
}

// Validate validation function.
func (name) Validate(v string) error {
	if len(v) == 0 {
		return primitive.NewValidationError("name is empty", nil)
	}

	return nil
}

// Value get value.
func (n name) Value() string {
	return n.value
}

// String get string mask primitive information.
func (n name) String() string {
	return "*****"
}
