package model_test_helper

import (
	"github.com/kruspe/music-rating/internal/model"
)

const (
	AComment            = "comment"
	AnotherComment      = "another-comment"
	AFestivalName       = "festival-name"
	AnotherFestivalName = "another-festival-name"
	ARating             = float64(5)
	AnotherRating       = float64(1)
	AYear               = 2020
	AnotherYear         = 2015
)

func ARatingForArtist(name string) model.Rating {
	return model.Rating{
		ArtistName:   name,
		Comment:      AComment,
		FestivalName: AFestivalName,
		Rating:       ARating,
		Year:         AYear,
	}
}
