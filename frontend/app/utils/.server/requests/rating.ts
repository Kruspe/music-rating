import {
  createAuthHeader,
  ErrorResponseData,
  FetchResponse,
  hasError,
} from "~/utils/.server/requests/util";
import { ArtistRating, toArtistRating } from "~/utils/types.server";

export interface ArtistRatingData {
  artist_name: string;
  festival_name?: string;
  rating: number;
  year?: number;
  comment?: string;
}

export interface RatingRequest {
  artist_name: string;
  festival_name?: string;
  rating: number;
  year?: number;
  comment?: string;
}

export async function saveRating(
  request: Request,
  rating: RatingRequest,
): Promise<FetchResponse> {
  const headers = await createAuthHeader(request);
  const response = await fetch(`${process.env.API_ENDPOINT}/ratings`, {
    headers: headers,
    method: "POST",
    body: JSON.stringify(rating),
  });
  if (hasError(response)) {
    const errorData: ErrorResponseData = await response.json();
    return { ok: false, error: errorData.error };
  }

  return { ok: true };
}

export async function updateRating(
  request: Request,
  rating: RatingRequest,
): Promise<FetchResponse> {
  const headers = await createAuthHeader(request);
  const response = await fetch(
    `${process.env.API_ENDPOINT}/ratings/${rating.artist_name}`,
    {
      headers: headers,
      method: "PUT",
      body: JSON.stringify(rating),
    },
  );
  if (hasError(response)) {
    const errorData: ErrorResponseData = await response.json();
    return { ok: false, error: errorData.error };
  }

  return { ok: true };
}

export async function getRatings(
  request: Request,
): Promise<FetchResponse<ArtistRating[]>> {
  const headers = await createAuthHeader(request);
  const response = await fetch(`${process.env.API_ENDPOINT}/ratings`, {
    headers: headers,
  });
  if (hasError(response)) {
    const errorData: ErrorResponseData = await response.json();
    return { ok: false, error: errorData.error };
  }

  const responseData: ArtistRatingData[] = await response.json();
  return { ok: true, data: responseData.map((r) => toArtistRating(r)) };
}

export async function getFestivalRatings(
  request: Request,
  festivalName: string,
): Promise<FetchResponse<ArtistRating[]>> {
  const headers = await createAuthHeader(request);
  const response = await fetch(
    `${process.env.API_ENDPOINT}/ratings/${festivalName}`,
    {
      headers: headers,
    },
  );
  if (hasError(response)) {
    const errorData: ErrorResponseData = await response.json();
    return { ok: false, error: errorData.error };
  }

  const responseData: ArtistRatingData[] = await response.json();
  return { ok: true, data: responseData.map((r) => toArtistRating(r)) };
}
