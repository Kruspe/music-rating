package model

import "fmt"

type FestivalNotSupportedError struct {
	FestivalName string
}

func (e FestivalNotSupportedError) Error() string {
	return fmt.Sprintf("Festival '%s' is not supported yet.", e.FestivalName)
}
