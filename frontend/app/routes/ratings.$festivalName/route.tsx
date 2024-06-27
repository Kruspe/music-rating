import { json, LoaderFunctionArgs, TypedResponse } from "@remix-run/node";
import { FetchResponse } from "~/utils/.server/requests/util";
import { ArtistRating } from "~/utils/types.server";
import { getFestivalRatings } from "~/utils/.server/requests/rating";
import {
  isRouteErrorResponse,
  useLoaderData,
  useRouteError,
} from "@remix-run/react";
import { Typography } from "@mui/material";
import { RatingTable } from "~/components/rating-table";

export function ErrorBoundary() {
  const error = useRouteError();

  if (isRouteErrorResponse(error)) {
    return <Typography variant="h3">{error.data}</Typography>;
  }
}

export async function loader({
  request,
  params,
}: LoaderFunctionArgs): Promise<TypedResponse<FetchResponse<ArtistRating[]>>> {
  const { festivalName } = params;
  const response = await getFestivalRatings(request, festivalName!);
  if (!response.ok) {
    throw json(response.error);
  }
  return json(response);
}

export default function FestivalRatingsRoute() {
  const loaderData = useLoaderData<typeof loader>();

  return <RatingTable data={loaderData.data!} updatable />;
}
