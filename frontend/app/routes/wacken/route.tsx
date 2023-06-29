import type { ActionArgs, LoaderArgs } from "@remix-run/node";
import { authenticator } from "~/utils/auth.server";
import { useLoaderData } from "@remix-run/react";
import Grid from "@mui/material/Unstable_Grid2";
import RatingCard from "~/routes/wacken/RatingCard";
import { Typography } from "@mui/material";
import { get } from "~/utils/request.server";
import type { ArtistRatingData } from "~/utils/types.server";

export async function loader({ request }: LoaderArgs) {
  return get<ArtistRatingData[]>(request, "/festivals/wacken");
}

export async function action({ request }: ActionArgs) {
  const user = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  const form = await request.formData();
  const artistName = form.get("artist_name") as string;
  const festival = form.get("festival") as string;
  const rating = form.get("rating") as string;
  const year = form.get("year") as string;
  const comment = form.get("comment") as string;

  await fetch(`${process.env.API_ENDPOINT}/ratings`, {
    method: "POST",
    body: JSON.stringify({
      artist_name: artistName,
      festival_name: festival,
      rating: parseFloat(rating),
      year: parseInt(year, 10),
      comment: comment,
    }),
    headers: {
      authorization: `Bearer ${user.token}`,
    },
  });
  return null;
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
