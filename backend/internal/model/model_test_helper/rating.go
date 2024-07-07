package model_test_helper

import (
	"github.com/kruspe/music-rating/internal/model"
	"log"
)

const (
	AnArtistName        = "artist-name"
	AComment            = "comment"
	AnotherComment      = "another-comment"
	AFestivalName       = "festival-name"
	AnotherFestivalName = "another-festival-name"
	ARating             = model.Rating(5)
	AnotherRating       = model.Rating(1)
	AYear               = 2020
	AnotherYear         = 2015
)

func AnArtistRating(name string) model.ArtistRating {
	festivalName := AFestivalName
	year := AYear
	comment := AComment
	rating, err := model.NewArtistRating(name, ARating.Float64(), &festivalName, &year, &comment)
	if err != nil {
		log.Panicln("could not create test ArtistRating", err)
	}
	return *rating
}

func AnArtistRatingWithRating(name string, rating float64) model.ArtistRating {
	festivalName := AFestivalName
	year := AYear
	comment := AComment
	artistRating, err := model.NewArtistRating(name, rating, &festivalName, &year, &comment)
	if err != nil {
		log.Panicln("could not create test ArtistRating", err)
	}
	return *artistRating
}
