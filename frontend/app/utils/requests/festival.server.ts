import { createAuthHeader, FetchResponse } from "~/utils/requests/util.server";
import { FestivalArtists } from "~/utils/types.server";

export interface RatingData {
  artist_name: string;
  festival_name: string;
  rating: string;
  year: string;
  comment: string;
}

export interface FestivalArtistsData {
  artists: FestivalArtistData[];
}

export interface FestivalArtistData {
  artist_name: string;
  image_url: string;
}

export async function getFestivalArtists(
  request: Request,
): Promise<FetchResponse<FestivalArtists>> {
  // const headers = await createAuthHeader(request);
}
