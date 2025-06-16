import {
  Button,
  Dialog,
  DialogContent,
  Fab,
  Rating,
  TextField,
  Typography,
  Grid,
} from "@mui/material";
import { data, Form } from "react-router";
import { Add } from "@mui/icons-material";
import { useState } from "react";
import {
  getRatings,
  type RatingRequest,
  saveRating,
  updateRating,
} from "~/utils/.server/requests/rating";
import { RatingTable } from "~/components/rating-table";
import type { Route } from "./+types/home";

export async function loader({ request }: Route.LoaderArgs) {
  const response = await getRatings(request);
  if (!response.ok) {
    throw data(response.error);
  }
  return data(response);
}

export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  const rating: RatingRequest = {
    artist_name: formData.get("artist_name") as string,
    rating: parseFloat(formData.get("rating") as string),
  };
  if (formData.get("festival_name")) {
    rating.festival_name = formData.get("festival_name") as string;
  }
  if (formData.get("year")) {
    rating.year = parseInt(formData.get("year") as string, 10);
  }
  if (formData.get("comment")) {
    rating.comment = formData.get("comment") as string;
  }
  let response;
  switch (formData.get("_action")) {
    case "SAVE_RATING": {
      response = saveRating(request, rating);
      break;
    }
    case "UPDATE_RATING": {
      response = updateRating(request, rating);
      break;
    }
  }
  return response;
}

export default function RatingsRoute({ loaderData }: Route.ComponentProps) {
  const [showAdd, setShowAdd] = useState(false);

  return (
    <>
      <Fab color="primary" aria-label="add">
        <Add onClick={() => setShowAdd((prevState) => !prevState)} />
      </Fab>
      {showAdd && (
        <Dialog open={showAdd} onClose={() => setShowAdd(false)}>
          <DialogContent>
            <Form method="post" onSubmit={() => setShowAdd(false)}>
              <Grid container spacing={2}>
                <Grid>
                  <Typography variant="h6">Add rating</Typography>
                </Grid>
                <Grid size={{ xs: 12 }}>
                  <TextField fullWidth name="artist_name" label="Artist" />
                </Grid>
                <Grid size={{ xs: 12 }}>
                  <TextField
                    fullWidth
                    name="festival_name"
                    label="Festival/Concert"
                  />
                </Grid>
                <Grid size={{ xs: 12 }}>
                  <TextField fullWidth name="year" label="Year" />
                </Grid>
                <Grid size={{ xs: 12 }}>
                  <Rating precision={0.5} name="rating" />
                </Grid>
                <Grid size={{ xs: 12 }}>
                  <TextField fullWidth name="comment" label="Comment" />
                </Grid>
                <Grid>
                  <Button
                    variant="outlined"
                    type="submit"
                    name="_action"
                    value="SAVE_RATING"
                  >
                    Rate
                  </Button>
                </Grid>
              </Grid>
            </Form>
          </DialogContent>
        </Dialog>
      )}
      {loaderData.data!.length > 0 ? (
        <RatingTable data={loaderData.data!} updatable />
      ) : (
        <Typography variant="h2">
          You have not rated any bands so far
        </Typography>
      )}
    </>
  );
}
