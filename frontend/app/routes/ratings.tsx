import type { GridRenderCellParams, GridValidRowModel } from "@mui/x-data-grid";
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
import {
  ActionFunctionArgs,
  LoaderFunctionArgs,
  redirect,
} from "@remix-run/node";
import { get } from "~/utils/request.server";
import { Form, useLoaderData, useSubmit } from "@remix-run/react";
import { authenticator } from "~/utils/auth.server";
import type { RatingData } from "~/utils/types.server";
import { Add } from "@mui/icons-material";
import { useState } from "react";

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

const columns = [
  { field: "artist_name", headerName: "Artist", flex: 2 },
  {
    field: "year",
    headerName: "Year",
    editable: true,
    flex: 1,
  },
  {
    field: "festival_name",
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

export async function loader({ request }: LoaderFunctionArgs) {
  return get<RatingData[]>(request, "/ratings");
}

export async function action({ request }: ActionFunctionArgs) {
  const { token } = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  const form = await request.formData();
  const artistName = form.get("artist_name") as string;
  const festival = form.get("festival_name") as string;
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
      authorization: `Bearer ${token}`,
    },
  });

  const intent = form.get("intent") as string;
  if (intent === "wacken") {
    return redirect("/wacken");
  }
  return null;
}

export default function RatingsRoute() {
  const ratings = useLoaderData<typeof loader>();
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
      {ratings.length > 0 ? (
        <DataGrid
          columns={columns}
          rows={ratings}
          getRowId={(row) => row.artist_name}
          autoHeight
          hideFooterSelectedRowCount
          disableColumnFilter
          disableColumnSelector
          processRowUpdate={(row) => {
            const formData = new FormData();
            formData.append("artist_name", row.artist_name);
            formData.append("festival_name", row.festival_name);
            formData.append("rating", row.rating);
            formData.append("year", row.year);
            formData.append("comment", row.comment);

            submit(formData, {
              method: "PUT",
              action: `/ratings/${row.artist_name}`,
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
