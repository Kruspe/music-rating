import { useQuery } from '@tanstack/react-query';
import {
  Table, TableBody, TableCell, TableHead, TableRow, Typography,
} from '@mui/material';
import { useAuth0 } from '@auth0/auth0-react';

export default function Ratings() {
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
        <Table sx={{ minWidth: 650 }} size="small" aria-label="a dense table">
          <TableHead>
            <TableRow>
              <TableCell>Artist</TableCell>
              <TableCell align="right">Year</TableCell>
              <TableCell align="right">Festival</TableCell>
              <TableCell align="right">Rating</TableCell>
              <TableCell align="right">Comment</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {ratings.map((rating) => (
              <TableRow
                key={rating.artist_name}
                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {rating.artist_name}
                </TableCell>
                <TableCell align="right">{rating.year}</TableCell>
                <TableCell align="right">{rating.festival_name}</TableCell>
                <TableCell align="right">{rating.rating}</TableCell>
                <TableCell align="right">{rating.comment}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )
      : <Typography variant="h2">You have not rated any bands so far</Typography>;
  }
  return <div />;
}
