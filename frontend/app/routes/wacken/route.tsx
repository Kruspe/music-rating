import type { DataFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import RatingCard from "~/routes/wacken/RatingCard";
import { Typography, Unstable_Grid2 as Grid } from "@mui/material";
import { get } from "~/utils/request.server";
import type { ArtistRatingData } from "~/utils/types.server";

export async function loader({ request }: DataFunctionArgs) {
  return get<ArtistRatingData[]>(request, "/festivals/wacken");
}

export default function WackenRoute() {
  const artists = useLoaderData<typeof loader>();

  return artists.length > 0 ? (
    <Grid container spacing={0.5}>
      {artists.map((artist) => (
        <Grid key={artist.artist_name}>
          <RatingCard
            artistName={artist.artist_name}
            imageUrl={artist.image_url}
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
