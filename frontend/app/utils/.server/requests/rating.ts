import {
  createAuthHeader,
  ErrorResponseData,
  FetchResponse,
  hasError,
} from "~/utils/.server/requests/util";

export interface RatingRequest {
  artist_name: string;
  festival_name: string;
  rating: number;
  year: number;
  comment: string;
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
