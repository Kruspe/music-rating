import { useMutation, useQuery } from '@tanstack/react-query';
import { DataGrid, GridToolbar, useGridApiContext } from '@mui/x-data-grid';
import { Rating, Typography } from '@mui/material';
import { useAuth0 } from '@auth0/auth0-react';
import PropTypes from 'prop-types';
import { useState } from 'react';

function useUpdateRating() {
  const { getAccessTokenSilently } = useAuth0();

  return useMutation({
    mutationFn: async (data) => {
      fetch(`${process.env.REACT_APP_API_ENDPOINT}/ratings/${data.artist_name}`, {
        method: 'PATCH',
        body: JSON.stringify(data),
        headers: {
          authorization: `Bearer ${await getAccessTokenSilently()}`,
        },
      });
    },
  });
}

function renderRating(rating) {
  return <Rating readOnly value={rating.value} />;
}

function EditRating({ field, id, value }) {
  const apiRef = useGridApiContext();
  const [rating, setRating] = useState(value);
  return (
    <Rating
      precision={1}
      value={rating}
      onChange={(event) => {
        const newRating = parseInt(event.target.value, 10);
        apiRef.current.setEditCellValue({
          id,
          field,
          value: newRating,
        });
        setRating(newRating);
      }}
    />
  );
}
EditRating.propTypes = {
  field: PropTypes.string.isRequired,
  id: PropTypes.string.isRequired,
  value: PropTypes.number.isRequired,
};

function renderUpdateRating(params) {
  return <EditRating id={params.id} field={params.field} value={params.value} />;
}

const columns = [
  { field: 'artist_name', headerName: 'Artist', flex: 2 },
  {
    field: 'year', headerName: 'Year', editable: true, flex: 1,
  },
  {
    field: 'festival_name', headerName: 'Festival', editable: true, flex: 1.5,
  },
  {
    field: 'rating',
    headerName: 'Rating',
    renderCell: renderRating,
    renderEditCell: renderUpdateRating,
    width: 180,
    editable: true,
    flex: 1,
  },
  {
    field: 'comment', headerName: 'Comment', flex: 4, sortable: false, editable: true,
  },
];

export default function Ratings() {
  const updateRating = useUpdateRating();
  const { getAccessTokenSilently, user } = useAuth0();

  const { data: ratings, isFetched } = useQuery(['ratings', user.sub], async () => {
    const result = await fetch(`${process.env.REACT_APP_API_ENDPOINT}/ratings`, {
      headers: {
        authorization: `Bearer ${(await getAccessTokenSilently())}`,
      },
    });
    return result.json();
  });

  if (isFetched) {
    return ratings.length > 0
      ? (
        <DataGrid
          columns={columns}
          rows={ratings}
          getRowId={(row) => row.artist_name}
          autoHeight
          hideFooterSelectedRowCount
          disableColumnFilter
          disableColumnSelector
          processRowUpdate={(row) => {
            updateRating.mutate({
              artist_name: row.artist_name,
              year: parseInt(row.year, 10),
              festival_name: row.festival_name,
              rating: parseInt(row.rating, 10),
              comment: row.comment,
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
      )
      : <Typography variant="h2">You have not rated any bands so far</Typography>;
  }
  return <div />;
}
