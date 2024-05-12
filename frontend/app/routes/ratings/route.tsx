import type {
  GridColDef,
  GridRenderCellParams,
  GridValidRowModel,
} from "@mui/x-data-grid";
import { DataGrid, GridToolbar, useGridApiContext } from "@mui/x-data-grid";
import {
  Button,
  Dialog,
  DialogContent,
  Fab,
  Rating,
  TextField,
  Typography,
  Unstable_Grid2 as Grid,
} from "@mui/material";
import { ActionFunctionArgs, json, LoaderFunctionArgs } from "@remix-run/node";
import { Form, useLoaderData, useSubmit } from "@remix-run/react";
import { Add } from "@mui/icons-material";
import { useState } from "react";
import {
  getRatings,
  saveRating,
  updateRating,
} from "~/utils/.server/requests/rating";
import { ArtistRating } from "~/utils/types.server";

function renderRating({ value }: GridRenderCellParams) {
  return <Rating readOnly defaultValue={value} precision={0.5} />;
}

function EditRatingCell({
  id,
  field,
  value,
}: GridRenderCellParams<GridValidRowModel, number>) {
  const apiRef = useGridApiContext();
  return (
    <Rating
      precision={0.5}
      name="rating"
      value={value}
      onChange={(event, newValue) =>
        apiRef.current.setEditCellValue({ id, field, value: newValue })
      }
    />
  );
}

function renderEditRating(params: GridRenderCellParams) {
  return <EditRatingCell {...params} />;
}

const columns: GridColDef<ArtistRating>[] = [
  {
    field: "artistName",
    headerName: "Artist",
    flex: 2,
  },
  {
    field: "year",
    headerName: "Year",
    editable: true,
    flex: 1,
  },
  {
    field: "festivalName",
    headerName: "Festival",
    editable: true,
    flex: 1.5,
  },
  {
    field: "rating",
    headerName: "Rating",
    renderCell: renderRating,
    renderEditCell: renderEditRating,
    width: 180,
    editable: true,
    flex: 1,
  },
  {
    field: "comment",
    headerName: "Comment",
    flex: 4,
    sortable: false,
    editable: true,
  },
];

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
  const submit = useSubmit();

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
                <Grid xs={12}>
                  <TextField fullWidth name="artist_name" label="Artist" />
                </Grid>
                <Grid xs={12}>
                  <TextField
                    fullWidth
                    name="festival_name"
                    label="Festival/Concert"
                  />
                </Grid>
                <Grid xs={12}>
                  <TextField fullWidth name="year" label="Year" />
                </Grid>
                <Grid xs={12}>
                  <Rating precision={0.5} name="rating" />
                </Grid>
                <Grid xs={12}>
                  <TextField fullWidth name="comment" label="Comment" />
                </Grid>
                <Grid>
                  <Button variant="outlined" type="submit">
                    Rate
                  </Button>
                </Grid>
              </Grid>
            </Form>
          </DialogContent>
        </Dialog>
      )}
      {loaderData.data!.length > 0 ? (
        <DataGrid
          columns={columns}
          rows={loaderData.data!}
          getRowId={(row) => row.artistName}
          autoHeight
          hideFooterSelectedRowCount
          disableColumnFilter
          disableColumnSelector
          processRowUpdate={(row) => {
            const formData = new FormData();
            formData.set("_action", "UPDATE_RATING");
            formData.append("artist_name", row.artistName);
            if (row.festivalName) {
              formData.append("festival_name", row.festivalName);
            }
            formData.append("rating", row.rating.toString());
            if (row.year) {
              formData.append("year", row.year.toString());
            }
            if (row.comment) {
              formData.append("comment", row.comment);
            }

            submit(formData, {
              method: "PUT",
              action: `/ratings`,
            });
            return row;
          }}
          slots={{ toolbar: GridToolbar }}
          slotProps={{
            toolbar: {
              showQuickFilter: true,
            },
          }}
        />
      ) : (
        <Typography variant="h2">
          You have not rated any bands so far
        </Typography>
      )}
    </>
  );
}
