import {
  DataGrid,
  type GridColDef,
  type GridRenderCellParams,
  GridToolbar,
  type GridValidRowModel,
  useGridApiContext,
} from "@mui/x-data-grid";
import type { ArtistRating } from "~/utils/types.server";
import { Rating } from "@mui/material";
import { useSubmit } from "react-router";

const getColumns = (updatable?: boolean): GridColDef<ArtistRating>[] => [
  {
    field: "artistName",
    headerName: "Artist",
    flex: 2,
  },
  {
    field: "year",
    headerName: "Year",
    editable: updatable,
    flex: 1,
  },
  {
    field: "festivalName",
    headerName: "Festival",
    editable: updatable,
    flex: 1.5,
  },
  {
    field: "rating",
    headerName: "Rating",
    renderCell: renderRating,
    renderEditCell: renderEditRating,
    width: 180,
    editable: updatable,
    flex: 1,
  },
  {
    field: "comment",
    headerName: "Comment",
    flex: 4,
    sortable: false,
    editable: updatable,
  },
];

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

interface RatingTableProps {
  data: ArtistRating[];
  updatable?: boolean;
}

export function RatingTable({ data, updatable }: RatingTableProps) {
  const submit = useSubmit();

  return (
    <DataGrid
      columns={getColumns(updatable)}
      rows={data}
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
  );
}
