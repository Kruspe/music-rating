package model

import "fmt"

type Rating struct {
	ArtistName   string
	Comment      string
	FestivalName string
	Rating       float64
	Year         int
}

type Ratings struct {
	Keys   []string
	Values map[string]Rating
}

type UpdateNonExistingRatingError struct {
	ArtistName string
}

func (e UpdateNonExistingRatingError) Error() string {
	return fmt.Sprintf("trying to update non existing rating for '%s'", e.ArtistName)
}
