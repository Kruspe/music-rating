import { data, isRouteErrorResponse } from "react-router";
import { Typography, Grid } from "@mui/material";
import RatingCard from "~/routes/festivals/RatingCard";
import { getUnratedFestivalArtists } from "~/utils/.server/requests/festival";
import type { Route } from "./+types/festival";

export function ErrorBoundary({ error }: Route.ErrorBoundaryProps) {
  if (isRouteErrorResponse(error)) {
    return <Typography variant="h3">{error.data}</Typography>;
  }
}

export async function loader({ request, params }: Route.LoaderArgs) {
  const { name } = params;
  const response = await getUnratedFestivalArtists(request, name!);
  if (!response.ok) {
    throw data(response.error);
  }
  return data(response);
}

export default function FestivalRoute({ loaderData }: Route.ComponentProps) {
  return loaderData.data!.length > 0 ? (
    <Grid container spacing={0.5}>
      {loaderData.data!.map((artist) => (
        <Grid key={artist.artistName} size={{ xs: 12, sm: 6, lg: 3, xl: 2 }}>
          <RatingCard
            artistName={artist.artistName}
            imageUrl={artist.imageUrl}
          />
        </Grid>
      ))}
    </Grid>
  ) : (
    <Typography variant="h2">
      You rated all bands that are announced at the moment
    </Typography>
  );
}
