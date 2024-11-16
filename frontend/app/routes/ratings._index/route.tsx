import {
  Button,
  Dialog,
  DialogContent,
  Fab,
  Rating,
  TextField,
  Typography,
  Grid2 as Grid,
} from "@mui/material";
import { ActionFunctionArgs, json, LoaderFunctionArgs } from "@remix-run/node";
import { Form, useLoaderData } from "@remix-run/react";
import { Add } from "@mui/icons-material";
import { useState } from "react";
import {
  getRatings,
  saveRating,
  updateRating,
} from "~/utils/.server/requests/rating";
import { RatingTable } from "~/components/rating-table";

export function ErrorBoundary() {}

export async function loader({ request }: LoaderFunctionArgs) {
  const data = await getRatings(request);
  if (!data.ok) {
    throw json(data.error);
  }
  return json(data);
}

export async function action({ request }: ActionFunctionArgs) {
  const formData = await request.formData();
  let response;
  switch (formData.get("_action")) {
    case "SAVE_RATING": {
      response = saveRating(request, {
        artist_name: formData.get("artist_name") as string,
        festival_name: formData.get("festival_name") as string,
        rating: parseFloat(formData.get("rating") as string),
        year: parseInt(formData.get("year") as string, 10),
        comment: formData.get("comment") as string,
      });
      break;
    }
    case "UPDATE_RATING": {
      response = updateRating(request, {
        artist_name: formData.get("artist_name") as string,
        festival_name: formData.get("festival_name") as string,
        rating: parseFloat(formData.get("rating") as string),
        year: parseInt(formData.get("year") as string, 10),
        comment: formData.get("comment") as string,
      });
      break;
    }
  }
  return response;
}

export default function RatingsRoute() {
  const loaderData = useLoaderData<typeof loader>();

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
