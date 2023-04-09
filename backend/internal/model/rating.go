package model

import "fmt"

type Rating struct {
	ArtistName   string
	Comment      string
	FestivalName string
	Rating       int
	Year         int
}

type RatingUpdate struct {
	Comment      string
	FestivalName string
	Rating       int
	Year         int
}

type UpdateNonExistingRatingError struct {
	ArtistName string
}

func (e UpdateNonExistingRatingError) Error() string {
	return fmt.Sprintf("trying to update non existing rating for '%s'", e.ArtistName)
}

func (e UpdateNonExistingRatingError) Is(err error) bool {
	_, ok := err.(*UpdateNonExistingRatingError)
	return ok
}
