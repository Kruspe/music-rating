import {
  createAuthHeader,
  ErrorResponseData,
  FetchResponse,
  hasError,
} from "~/utils/.server/requests/util";
import { FestivalArtist, toFestivalArtist } from "~/utils/types.server";

export interface FestivalArtistData {
  artist_name: string;
  image_url: string;
}

export async function getUnratedFestivalArtists(
  request: Request,
  festivalName: string,
): Promise<FetchResponse<FestivalArtist[]>> {
  const headers = await createAuthHeader(request);
  const response = await fetch(
    `${process.env.API_ENDPOINT}/festivals/${festivalName}`,
    {
      headers: headers,
    },
  );
  if (hasError(response)) {
    const errorData: ErrorResponseData = await response.json();
    return { ok: false, error: errorData.error };
  }

  const responseData: FestivalArtistData[] = await response.json();
  return {
    ok: true,
    data: responseData.map((a) => toFestivalArtist(a)),
  };
}
