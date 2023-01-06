package model

type Rating struct {
	ArtistName   string
	Comment      string
	FestivalName string
	Rating       int
	Year         int
}

type Festival struct {
	Artists []Artist
}

type Artist struct {
	ArtistName string
	ImageUrl   string
}
