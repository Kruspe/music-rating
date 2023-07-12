import type { GridRenderCellParams } from "@mui/x-data-grid";
import { DataGrid, GridToolbar, useGridApiContext } from "@mui/x-data-grid";
import { Rating, Typography } from "@mui/material";
import type { ActionArgs, LoaderArgs } from "@remix-run/node";
import { get } from "~/utils/request.server";
import { useLoaderData, useSubmit } from "@remix-run/react";
import { authenticator } from "~/utils/auth.server";
import type { RatingData } from "~/utils/types.server";

function renderRating({ value }: GridRenderCellParams) {
  return <Rating readOnly defaultValue={value} precision={0.5} />;
}

function EditRatingCell({
  id,
  field,
  value,
}: GridRenderCellParams<any, number>) {
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

export async function loader({ request }: LoaderArgs) {
  return get<RatingData[]>(request, "/ratings");
}

export async function action({ request }: ActionArgs) {
  const { token } = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  const form = await request.formData();
  const artistName = form.get("artist_name") as string;
  const festivalName = form.get("festival_name") as string;
  const rating = form.get("rating") as string;
  const year = form.get("year") as string;
  const comment = form.get("comment") as string;

  await fetch(`${process.env.API_ENDPOINT}/ratings/${artistName}`, {
    method: "PUT",
    body: JSON.stringify({
      festival_name: festivalName,
      rating: parseFloat(rating),
      year: parseInt(year, 10),
      comment: comment,
    }),
    headers: {
      authorization: `Bearer ${token}`,
    },
  });
  return null;
}

export default function RatingsRoute() {
  const ratings = useLoaderData<typeof loader>();
  const submit = useSubmit();

  return ratings.length > 0 ? (
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

        submit(formData, { method: "POST" });
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
    <Typography variant="h2">You have not rated any bands so far</Typography>
  );
}
