import { FestivalArtistData } from "~/utils/.server/requests/festival";

export interface FestivalArtist {
  artistName: string;
  imageUrl: string;
}

export function toFestivalArtist(d: FestivalArtistData): FestivalArtist {
  return {
    artistName: d.artist_name,
    imageUrl: d.image_url,
  };
}

export interface ArtistRating {
  artistName: string;
  festivalName: string;
  rating: string;
  year: string;
  comment: string;
}
