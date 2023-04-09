package model

import "fmt"

type MissingParameterError struct {
	ParameterName string
}

func (mpe MissingParameterError) Error() string {
	return fmt.Sprintf("missing parameter '%s'", mpe.ParameterName)
}

func (mpe MissingParameterError) Is(err error) bool {
	_, ok := err.(*MissingParameterError)
	return ok
}
