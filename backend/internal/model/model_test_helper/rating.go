package model_test_helper

import (
	"fmt"
	"github.com/kruspe/music-rating/internal/model"
)

func ARatingForArtist(name string) model.Rating {
	return model.Rating{
		ArtistName:   name,
		Comment:      fmt.Sprintf("Comment for %s", name),
		FestivalName: "Wacken",
		Rating:       5,
		Year:         2020,
	}
}
