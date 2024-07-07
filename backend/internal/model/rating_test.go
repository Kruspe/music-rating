package model_test

import (
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/model_test_helper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
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
	rating, err := model.NewArtistRating(AnArtistName, ARating, &festivalName, &year, &comment)
	require.NoError(s.T(), err)
	require.Equal(s.T(), AnArtistName, rating.ArtistName)
	require.Equal(s.T(), ARating, rating.Rating)
	require.Equal(s.T(), festivalName, *rating.FestivalName)
	require.Equal(s.T(), year, *rating.Year)
	require.Equal(s.T(), comment, *rating.Comment)
}

func (s *ratingSuite) Test_NewArtistRating_Errors_WhenRatingIsInvalid() {
	festivalName := AFestivalName
	year := AYear
	comment := AComment
	_, err := model.NewArtistRating(AnArtistName, -0.5, &festivalName, &year, &comment)
	require.Error(s.T(), err)
	require.IsType(s.T(), &model.InvalidFieldError[float64]{}, err)
}

func (s *ratingSuite) Test_NewArtistRating_Errors_WhenArtistNameIsEmpty() {
	festivalName := AFestivalName
	year := AYear
	comment := AComment
	_, err := model.NewArtistRating("", 5, &festivalName, &year, &comment)
	require.Error(s.T(), err)
	require.IsType(s.T(), &model.InvalidFieldError[string]{}, err)
}

func (s *ratingSuite) Test_NewRating() {
	rating, err := model.NewRating(float64(5))
	require.NoError(s.T(), err)
	require.Equal(s.T(), float64(5), rating.Float64())

	rating, err = model.NewRating(float64(0))
	require.NoError(s.T(), err)
	require.Equal(s.T(), float64(0), rating.Float64())
}

func (s *ratingSuite) Test_NewRating_Errors_WhenRatingIsInvalid() {
	_, err := model.NewRating(-0.5)
	require.Error(s.T(), err)
	require.IsType(s.T(), &model.InvalidFieldError[float64]{}, err)

	_, err = model.NewRating(5.1)
	require.Error(s.T(), err)
	require.IsType(s.T(), &model.InvalidFieldError[float64]{}, err)
}
