package festival_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kruspe/music-rating/internal/handler"
	"github.com/kruspe/music-rating/internal/handler/errors"
	. "github.com/kruspe/music-rating/internal/handler/test"
	"github.com/kruspe/music-rating/internal/model"
	. "github.com/kruspe/music-rating/internal/model/helper"
	"github.com/kruspe/music-rating/internal/persistence"
	. "github.com/kruspe/music-rating/internal/persistence/helper"
	"github.com/kruspe/music-rating/internal/usecase"
	"github.com/stretchr/testify/suite"
)

type unratedArtistResponse struct {
	ArtistName string `json:"artist_name"`
	ImageUrl   string `json:"image_url"`
}

type festivalHandlerSuite struct {
	suite.Suite
	repos *persistence.Repositories
	ph    *PersistenceHelper
}

func TestFestivalHandlerSuite(t *testing.T) {
	suite.Run(t, &festivalHandlerSuite{})
}

func (s *festivalHandlerSuite) BeforeTest(_, _ string) {
	s.ph = NewPersistenceHelper()
	s.repos = persistence.NewRepositories(s.ph.Dynamo, s.ph.TableName)
}

func (s *festivalHandlerSuite) Test_GetArtistsForFestival_Returns200AndAllArtists() {
	err := s.repos.RatingRepo.Save(context.Background(), AnUserId, AnArtistRating("Bloodbath"))
	s.Require().NoError(err)
	err = s.repos.RatingRepo.Save(context.Background(), AnUserId, AnArtistRating("Hypocrisy"))
	s.Require().NoError(err)

	bloodbath := AnArtistWithName("Bloodbath")
	hypocrisy := AnArtistWithName("Hypocrisy")
	festivalStorage := persistence.NewFestivalStorage(s.ph.MockFestivals(map[string][]model.Artist{
		AFestivalName: {
			bloodbath,
			hypocrisy,
		},
	}))
	useCases := usecase.NewUseCases(s.repos, festivalStorage)

	request := NewAuthenticatedRequest(http.MethodGet, fmt.Sprintf("/festivals/%s", AFestivalName), nil)
	recorder := httptest.NewRecorder()
	mux := http.NewServeMux()
	handler.Register(mux, &handler.Config{
		RatingRepo:      s.repos.RatingRepo,
		FestivalUseCase: useCases.FestivalUseCase,
	})
	mux.ServeHTTP(recorder, request)

	s.Equal(http.StatusOK, recorder.Result().StatusCode)

	var r []unratedArtistResponse
	err = json.NewDecoder(recorder.Body).Decode(&r)
	s.Require().NoError(err)
	s.Len(r, 2)
	s.Equal(bloodbath.Name, r[0].ArtistName)
	s.Equal(bloodbath.ImageUrl, r[0].ImageUrl)
	s.Equal(hypocrisy.Name, r[1].ArtistName)
	s.Equal(hypocrisy.ImageUrl, r[1].ImageUrl)
}

func (s *festivalHandlerSuite) Test_GetArtistsForFestival_Returns200AndAllUnratedArtists_WhenFiltering() {
	err := s.repos.RatingRepo.Save(context.Background(), AnUserId, AnArtistRating("Bloodbath"))
	s.Require().NoError(err)
	err = s.repos.RatingRepo.Save(context.Background(), AnUserId, AnArtistRating("Hypocrisy"))
	s.Require().NoError(err)

	unratedArtist := AnArtistWithName("Benediction")
	festivalStorage := persistence.NewFestivalStorage(s.ph.MockFestivals(map[string][]model.Artist{
		AFestivalName: {
			AnArtistWithName("Bloodbath"),
			AnArtistWithName("Hypocrisy"),
			unratedArtist,
		},
	}))
	useCases := usecase.NewUseCases(s.repos, festivalStorage)

	request := NewAuthenticatedRequest(http.MethodGet, fmt.Sprintf("/festivals/%s?filter=unrated", AFestivalName), nil)
	recorder := httptest.NewRecorder()
	mux := http.NewServeMux()
	handler.Register(mux, &handler.Config{
		RatingRepo:      s.repos.RatingRepo,
		FestivalUseCase: useCases.FestivalUseCase,
	})
	mux.ServeHTTP(recorder, request)

	s.Equal(http.StatusOK, recorder.Result().StatusCode)

	var r []unratedArtistResponse
	err = json.NewDecoder(recorder.Body).Decode(&r)
	s.Require().NoError(err)
	s.Len(r, 1)
	s.Equal(unratedArtist.Name, r[0].ArtistName)
	s.Equal(unratedArtist.ImageUrl, r[0].ImageUrl)
}

func (s *festivalHandlerSuite) Test_GetArtistsForFestival_Returns200AndEmptyList_WhenFilteringAndAllArtistsAreRated() {
	err := s.repos.RatingRepo.Save(context.Background(), AnUserId, AnArtistRating("Bloodbath"))
	s.Require().NoError(err)
	err = s.repos.RatingRepo.Save(context.Background(), AnUserId, AnArtistRating("Hypocrisy"))
	s.Require().NoError(err)
	err = s.repos.RatingRepo.Save(context.Background(), AnUserId, AnArtistRating("Benediction"))
	s.Require().NoError(err)

	festivalStorage := persistence.NewFestivalStorage(s.ph.MockFestivals(map[string][]model.Artist{
		AFestivalName: {
			AnArtistWithName("Bloodbath"),
			AnArtistWithName("Hypocrisy"),
			AnArtistWithName("Benediction"),
		},
	}))
	useCases := usecase.NewUseCases(s.repos, festivalStorage)

	request := NewAuthenticatedRequest(http.MethodGet, fmt.Sprintf("/festivals/%s?filter=unrated", AFestivalName), nil)
	recorder := httptest.NewRecorder()
	mux := http.NewServeMux()
	handler.Register(mux, &handler.Config{
		RatingRepo:      s.repos.RatingRepo,
		FestivalUseCase: useCases.FestivalUseCase,
	})
	mux.ServeHTTP(recorder, request)

	s.Equal(http.StatusOK, recorder.Result().StatusCode)
	var r []unratedArtistResponse
	err = json.NewDecoder(recorder.Body).Decode(&r)
	s.Require().NoError(err)
	s.Empty(r, 0)
}

func (s *festivalHandlerSuite) Test_GetArtistsForFestival_Returns404_WhenFestivalIsNotSupported() {
	festivalStorage := persistence.NewFestivalStorage(s.ph.MockFestivals(map[string][]model.Artist{
		AFestivalName: {
			AnArtistWithName("Bloodbath"),
			AnArtistWithName("Hypocrisy"),
		},
	}))
	useCases := usecase.NewUseCases(s.repos, festivalStorage)

	request := NewAuthenticatedRequest(http.MethodGet, fmt.Sprintf("/festivals/%s", AnotherFestivalName), nil)
	recorder := httptest.NewRecorder()
	mux := http.NewServeMux()
	handler.Register(mux, &handler.Config{
		RatingRepo:      s.repos.RatingRepo,
		FestivalUseCase: useCases.FestivalUseCase,
	})
	mux.ServeHTTP(recorder, request)

	s.Equal(http.StatusNotFound, recorder.Result().StatusCode)
	var r errors.ErrorResponse
	err := json.NewDecoder(recorder.Body).Decode(&r)
	s.Require().NoError(err)
	s.Equal(model.FestivalNotSupportedError{FestivalName: AnotherFestivalName}.Error(), r.Error)
}
