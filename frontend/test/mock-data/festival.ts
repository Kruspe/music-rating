import { FestivalArtistData } from "~/utils/.server/requests/festival";
import { testArtistName } from "./rating";

export const testFestivalArtistsData: FestivalArtistData[] = [
  {
    artist_name: testArtistName,
    image_url: "http://bloodbath.com",
  },
  {
    artist_name: "Hypocrisy",
    image_url: "http://hypocrisy.com",
  },
  {
    artist_name: "Deserted Fear",
    image_url: "http://deserted_fear.com",
  },
];
