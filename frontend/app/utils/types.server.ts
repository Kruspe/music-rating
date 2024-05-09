import { FestivalArtistData } from "~/utils/.server/requests/festival";
import { ArtistRatingData } from "~/utils/.server/requests/rating";

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
  festivalName?: string;
  rating: number;
  year?: number;
  comment?: string;
}

export function toArtistRating(d: ArtistRatingData): ArtistRating {
  return {
    artistName: d.artist_name,
    festivalName: d.festival_name,
    rating: d.rating,
    year: d.year,
    comment: d.comment,
  };
}
