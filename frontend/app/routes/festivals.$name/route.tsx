import { useLoaderData } from "@remix-run/react";
import RatingCard from "~/routes/festivals.$name/RatingCard";
import { Typography, Unstable_Grid2 as Grid } from "@mui/material";
import type { FestivalArtist } from "~/utils/types.server";
import { json, LoaderFunctionArgs, TypedResponse } from "@remix-run/node";
import { getUnratedFestivalArtists } from "~/utils/.server/requests/festival";
import { FetchResponse } from "~/utils/.server/requests/util";

export async function loader({
  request,
  params,
}: LoaderFunctionArgs): Promise<
  TypedResponse<FetchResponse<FestivalArtist[]>>
> {
  const { name } = params;
  const response = await getUnratedFestivalArtists(request, name!);
  if (!response.ok) {
    throw json(response.error);
  }
  return json(response);
}

export default function FestivalRoute() {
  const loaderData = useLoaderData<typeof loader>();

  return loaderData.data!.length > 0 ? (
    <Grid container spacing={0.5}>
      {loaderData.data!.map((artist) => (
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
