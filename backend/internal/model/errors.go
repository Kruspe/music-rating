package model

import "fmt"

type MissingParameterError struct {
	ParameterName string
}

func (e MissingParameterError) Error() string {
	return fmt.Sprintf("missing parameter '%s'", e.ParameterName)
}
