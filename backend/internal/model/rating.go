package model

type Rating struct {
	ArtistName   string
	Comment      string
	FestivalName string
	Rating       int
	Year         int
}

type RatingUpdate struct {
	Comment      string
	FestivalName string
	Rating       int
	Year         int
}
