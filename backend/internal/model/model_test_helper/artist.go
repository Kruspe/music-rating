package model_test_helper

import (
	"fmt"
	"github.com/kruspe/music-rating/internal/model"
)

func AnArtistWithName(name string) model.Artist {
	return model.Artist{
		Name:     name,
		ImageUrl: fmt.Sprintf("https://%s.com", name),
	}
}
