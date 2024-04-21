import { useLoaderData } from "@remix-run/react";
import RatingCard from "~/routes/wacken/RatingCard";
import { Typography, Unstable_Grid2 as Grid } from "@mui/material";
import { get } from "~/utils/request.server";
import type { FestivalArtist } from "~/utils/types.server";
import { LoaderFunctionArgs } from "@remix-run/node";

export async function loader({ request }: LoaderFunctionArgs) {
  return get<FestivalArtist[]>(request, "/festivals/wacken");
}

export default function WackenRoute() {
  const artists = useLoaderData<typeof loader>();

  return artists.length > 0 ? (
    <Grid container spacing={0.5}>
      {artists.map((artist) => (
        <Grid key={artist.artistName}>
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
