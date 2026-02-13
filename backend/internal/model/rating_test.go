package model_test

import (
	"testing"

	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/helper"
	"github.com/stretchr/testify/suite"
)

type ratingSuite struct {
	suite.Suite
}

func Test_RatingSuite(t *testing.T) {
	suite.Run(t, &ratingSuite{})
}

func (s *ratingSuite) Test_NewArtistRating() {
	festivalName := AFestivalName
	year := AYear
	comment := AComment
	rating, err := model.NewArtistRating(AnArtistName, ARating.Float64(), &festivalName, &year, &comment)
	s.Require().NoError(err)
	s.Equal(AnArtistName, rating.ArtistName)
	s.InEpsilon(ARating.Float64(), rating.Rating.Float64(), 0.0001)
	s.Equal(festivalName, *rating.FestivalName)
	s.Equal(year, *rating.Year)
	s.Equal(comment, *rating.Comment)
}

func (s *ratingSuite) Test_NewArtistRating_Errors_WhenRatingIsInvalid() {
	festivalName := AFestivalName
	year := AYear
	comment := AComment
	_, err := model.NewArtistRating(AnArtistName, -0.5, &festivalName, &year, &comment)
	s.Require().Error(err)
	s.Require().ErrorAs(err, new(*model.InvalidFieldError[float64]))
}

func (s *ratingSuite) Test_NewArtistRating_Errors_WhenArtistNameIsEmpty() {
	festivalName := AFestivalName
	year := AYear
	comment := AComment
	_, err := model.NewArtistRating("", 5, &festivalName, &year, &comment)
	s.Require().Error(err)
	s.Require().ErrorAs(err, new(*model.InvalidFieldError[string]))
}

func (s *ratingSuite) Test_NewRating() {
	rating, err := model.NewRating(float64(5))
	s.Require().NoError(err)
	s.InEpsilon(float64(5), rating.Float64(), 0.0001)

	rating, err = model.NewRating(float64(0))
	s.Require().NoError(err)
	s.InDelta(float64(0), rating.Float64(), 0.0001)
}

func (s *ratingSuite) Test_NewRating_Errors_WhenRatingIsInvalid() {
	_, err := model.NewRating(-0.5)
	s.Require().Error(err)
	s.Require().ErrorAs(err, new(*model.InvalidFieldError[float64]))

	_, err = model.NewRating(5.1)
	s.Require().Error(err)
	s.Require().ErrorAs(err, new(*model.InvalidFieldError[float64]))
}
