import { data, isRouteErrorResponse } from "react-router";
import { getFestivalRatings } from "~/utils/.server/requests/rating";
import { Typography } from "@mui/material";
import { RatingTable } from "~/components/rating-table";
import type { Route } from "./+types/festival";

export function ErrorBoundary({ error }: Route.ErrorBoundaryProps) {
  if (isRouteErrorResponse(error)) {
    return <Typography variant="h3">{error.data}</Typography>;
  }
}

export async function loader({ request, params }: Route.LoaderArgs) {
  const { festivalName } = params;
  const response = await getFestivalRatings(request, festivalName!);
  if (!response.ok) {
    throw data(response.error);
  }
  return data(response);
}

export default function FestivalRatingsRoute({
  loaderData,
}: Route.ComponentProps) {
  return <RatingTable data={loaderData.data!} updatable />;
}
