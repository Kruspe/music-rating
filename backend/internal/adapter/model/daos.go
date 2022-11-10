package model

type RatingDao struct {
	ArtistName   string `json:"artist_name"`
	Comment      string `json:"comment,omitempty"`
	FestivalName string `json:"festival_name"`
	Rating       *int   `json:"rating"`
	Year         *int   `json:"year"`
}
