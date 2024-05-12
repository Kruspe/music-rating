import { ArtistRatingData } from "~/utils/.server/requests/rating";

export const testFestivalName = "Wacken";
export const testArtistName = "Bloodbath";

export const testArtistRatingData: ArtistRatingData = {
  artist_name: testArtistName,
  festival_name: testFestivalName,
  rating: 5,
  year: 2015,
  comment: "Old school swedish death metal",
};
export const testArtistRatingsData: ArtistRatingData[] = [
  testArtistRatingData,
  {
    artist_name: "Hypocrisy",
    festival_name: "Dong",
    rating: 5,
    year: 2023,
  },
  {
    artist_name: "Deserted Fear",
    rating: 4.5,
  },
];
