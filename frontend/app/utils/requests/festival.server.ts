import {
  createAuthHeader,
  FetchResponse,
  hasError,
} from "~/utils/requests/util.server";
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
  festivalName: string,
): Promise<FetchResponse<FestivalArtists>> {
  const header = await createAuthHeader(request);
  const response = await fetch(
    `${process.env.API_ENDPOINT}/festivals/${festivalName}`,
    {
      headers: header,
    },
  );
  if (hasError(response)) {
    return { error: await response.text(), ok: false };
  }
  const festivalArtists: FestivalArtistsData = await response.json();
  return {
    data: {
      artists: festivalArtists.artists.map((a) => ({
        artistName: a.artist_name,
        imageUrl: a.image_url,
      })),
    },
    ok: true,
  };
}
