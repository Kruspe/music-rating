package model_test_helper

import (
	"backend/internal/adapter/model"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
)

const (
	TestToken  = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJNZXRhbExvdmVyNjY2IiwiaWF0IjoxNTE2MjM5MDIyfQ.JZ3R_3it-1K9ttH5NA80fpIsBhnW6DNsIzwB2zEFRmo7hgE-HhW3jJbArXNS0fw2Pcj-xrU-DMF8KoLr8_EJh2XdTDjaRqz859p0RJ1gPLovsQ8N1HeqeQXKi2mwDJe2rZhWILHdWZ1zmduCY5fF8jUYyBIqLRh1B44L_CBlgeEejKoJfw7V3WoZhxdLeW8SlS2PQ7kN0XIyzm-_TPq1j5QnNHRnXRIh8V7o9rBtdM7PVGEFTpzb1jC6bZ3W-aHuZEWk5e1kRTV8IOXiLf-xtPQ42Hn4j2F27mDg0h2PsgVWmNjr2eqc9y0izps-rmoXHnzmBzvbtGS2yytEFw_WAA"
	TestUserId = "MetalLover666"
)

const (
	BloodbathImageUrl = "https://bloodbath.com"
	HypocrisyImageUrl = "https://hypocrisy.com"
)

var BloodbathRating = model.Rating{
	ArtistName:   "Bloodbath",
	Comment:      "old school swedish death metal",
	FestivalName: "Wacken",
	Rating:       5,
	Year:         2015,
}
var HypocrisyRating = model.Rating{
	ArtistName:   "Hypocrisy",
	FestivalName: "Wacken",
	Rating:       2022,
	Year:         5,
}

var BloodbathRatingDao = model.RatingDao{
	ArtistName:   BloodbathRating.ArtistName,
	Comment:      BloodbathRating.Comment,
	FestivalName: BloodbathRating.FestivalName,
	Rating:       aws.Int(BloodbathRating.Rating),
	Year:         aws.Int(BloodbathRating.Year),
}

var BloodbathRatingRecord = model.RatingRecord{
	DbKey: model.DbKey{
		PK: fmt.Sprintf("USER#%s", TestUserId),
		SK: fmt.Sprintf("ARTIST#%s", BloodbathRating.ArtistName),
	},
	Type:         model.RatingType,
	ArtistName:   BloodbathRating.ArtistName,
	Comment:      BloodbathRating.Comment,
	FestivalName: BloodbathRating.FestivalName,
	Rating:       BloodbathRating.Rating,
	UserId:       TestUserId,
	Year:         BloodbathRating.Year,
}
var HypocrisyRatingRecord = model.RatingRecord{
	DbKey: model.DbKey{
		PK: fmt.Sprintf("USER#%s", TestUserId),
		SK: fmt.Sprintf("ARTIST#%s", HypocrisyRating.ArtistName),
	},
	Type:         model.RatingType,
	ArtistName:   HypocrisyRating.ArtistName,
	FestivalName: HypocrisyRating.FestivalName,
	Rating:       HypocrisyRating.Rating,
	UserId:       TestUserId,
	Year:         HypocrisyRating.Year,
}

var ArtistsRecord = []model.ArtistRecord{
	{
		Artist: BloodbathRating.ArtistName,
		Image:  BloodbathImageUrl,
	},
	{
		Artist: HypocrisyRating.ArtistName,
		Image:  HypocrisyImageUrl,
	},
	{
		Artist: UnratedArtist.ArtistName,
		Image:  UnratedArtist.ImageUrl,
	},
}

var Festival = model.Festival{
	Artists: []model.Artist{
		{
			ArtistName: BloodbathRating.ArtistName,
			ImageUrl:   BloodbathImageUrl,
		},
		{
			ArtistName: HypocrisyRating.ArtistName,
			ImageUrl:   HypocrisyImageUrl,
		},
		{
			ArtistName: UnratedArtist.ArtistName,
			ImageUrl:   UnratedArtist.ImageUrl,
		},
	},
}

var UnratedArtist = model.Artist{
	ArtistName: "Benediction",
	ImageUrl:   "https://unrated-artist-image.url",
}

var UnratedArtistDao = model.ArtistDao(UnratedArtist)
