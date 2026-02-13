package model

import "fmt"

type ArtistRating struct {
	ArtistName   string
	Rating       Rating
	FestivalName *string
	Year         *int
	Comment      *string
}

func NewArtistRating(artistName string, rating float64, festivalName *string, year *int, comment *string) (*ArtistRating, error) {
	if artistName == "" {
		return nil, &InvalidFieldError[string]{
			Condition:     "must not be empty",
			FieldName:     "ArtistName",
			ProvidedValue: artistName,
		}
	}
	r, err := NewRating(rating)
	if err != nil {
		return nil, err
	}
	return &ArtistRating{
		ArtistName:   artistName,
		Rating:       *r,
		FestivalName: festivalName,
		Year:         year,
		Comment:      comment,
	}, nil
}

type Ratings struct {
	Keys   []string
	Values map[string]ArtistRating
}

type UpdateNonExistingRatingError struct {
	ArtistName string
}

func (e UpdateNonExistingRatingError) Error() string {
	return fmt.Sprintf("trying to update non existing rating for '%s'", e.ArtistName)
}

type Rating float64

func NewRating(rating float64) (*Rating, error) {
	if rating < 0 || rating > 5 {
		return nil, &InvalidFieldError[float64]{
			Condition:     "must be between 0 and 5",
			FieldName:     "Rating",
			ProvidedValue: rating,
		}
	}

	r := Rating(rating)
	return &r, nil
}

func (r Rating) Float64() float64 {
	return float64(r)
}

func (r Rating) String() string {
	return fmt.Sprintf("%.1f", r)
}
