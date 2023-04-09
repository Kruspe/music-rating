package persistence_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kruspe/music-rating/internal/adapter/persistence"
	"github.com/kruspe/music-rating/internal/adapter/persistence/persistence_test_helper"
	"github.com/kruspe/music-rating/internal/model"
	"github.com/kruspe/music-rating/internal/model/model_test_helper"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io"
	"os"
	"testing"
)

type storageSuite struct {
	suite.Suite
}

func Test_StorageSuite(t *testing.T) {
	suite.Run(t, &storageSuite{})
}

func (s *storageSuite) SetupSuite() {
	err := os.Setenv("FESTIVAL_ARTIST_BUCKET_NAME", "test-bucket")
	require.NoError(s.T(), err)
}

func (s *storageSuite) Test_GetArtists() {
	artists := []model.Artist{
		model_test_helper.AnArtistWithName("Bloodbath"),
		model_test_helper.AnArtistWithName("Hypocrisy"),
		model_test_helper.AnArtistWithName("Benediction"),
	}
	s3Mock := func(t *testing.T) persistence.S3Client {
		return persistence_test_helper.MockS3Client{
			GetObjectMock: func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
				t.Helper()
				if params.Bucket == nil {
					t.Fatal("expect bucket not to be nil")
				}
				if e, a := "test-bucket", *params.Bucket; e != a {
					t.Fatalf("expect %v, got %v", e, a)
				}
				if params.Key == nil {
					t.Fatal("expect key not to be nil")
				}
				if e, a := "festival-name.json", *params.Key; e != a {
					t.Fatalf("expect %v, got %v", e, a)
				}

				s3Body, err := json.Marshal([]persistence.ArtistRecord{
					{
						Artist: artists[0].Name,
						Image:  artists[0].ImageUrl,
					},
					{
						Artist: artists[1].Name,
						Image:  artists[1].ImageUrl,
					},
					{
						Artist: artists[2].Name,
						Image:  artists[2].ImageUrl,
					},
				})
				require.NoError(t, err)
				return &s3.GetObjectOutput{
					Body: io.NopCloser(bytes.NewReader(s3Body)),
				}, nil
			},
		}
	}

	storage := persistence.NewFestivalStorage(s3Mock(s.T()))
	artists, err := storage.GetArtists(context.Background(), "festival-name")
	require.NoError(s.T(), err)
	require.Equal(s.T(), artists, artists)
}

func (s *storageSuite) Test_GetArtists_ReturnsS3Error() {
	s3Mock := func(t *testing.T) persistence.S3Client {
		return persistence_test_helper.MockS3Client{
			GetObjectMock: func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
				t.Helper()
				return nil, errors.New("an error occurred")
			},
		}
	}
	storage := persistence.NewFestivalStorage(s3Mock(s.T()))
	_, err := storage.GetArtists(context.Background(), "festival-name")
	require.ErrorContains(s.T(), err, "an error occurred")
}
