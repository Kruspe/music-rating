package model

import "fmt"

type MissingParameterError struct {
	ParameterName string
}

func (e MissingParameterError) Error() string {
	return fmt.Sprintf("missing parameter '%s'", e.ParameterName)
}

type InvalidFieldError[T any] struct {
	Condition     string
	FieldName     string
	ProvidedValue T
}

func (e InvalidFieldError[T]) Error() string {
	return fmt.Sprintf("'%v' is invalid for field '%s': %s", e.ProvidedValue, e.FieldName, e.Condition)
}
