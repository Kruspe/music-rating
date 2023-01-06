package model

const (
	RatingType = "RATING"
)

type DbKey struct {
	PK string `dynamodbav:"PK"`
	SK string `dynamodbav:"SK"`
}

type RatingRecord struct {
	DbKey
	Type         string `dynamodbav:"type"`
	ArtistName   string `dynamodbav:"artist_name"`
	Comment      string `dynamodbav:"comment"`
	FestivalName string `dynamodbav:"festival_name"`
	Rating       int    `dynamodbav:"rating"`
	UserId       string `dynamodbav:"user_id"`
	Year         int    `dynamodbav:"year"`
}

type ArtistRecord struct {
	Artist string `json:"artist"`
	Image  string `json:"image"`
}
