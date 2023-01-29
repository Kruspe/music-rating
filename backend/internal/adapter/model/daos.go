package model

type RatingDao struct {
	ArtistName   string `json:"artist_name"`
	Comment      string `json:"comment"`
	FestivalName string `json:"festival_name"`
	Rating       int    `json:"rating"`
	Year         int    `json:"year"`
}

type ArtistDao struct {
	ArtistName string `json:"artist_name"`
	ImageUrl   string `json:"image_url"`
}
