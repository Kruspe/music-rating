import {
  Button,
  Card, CardMedia, Rating, TextField, Typography,
} from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import PropTypes from 'prop-types';
import { useState } from 'react';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useAuth0 } from '@auth0/auth0-react';

export default function RatingCard({ artistName, imageUrl }) {
  const queryClient = useQueryClient();
  const { getAccessTokenSilently, user } = useAuth0();

  const [festival, setFestival] = useState('');
  const [year, setYear] = useState('');
  const [rating, setRating] = useState(0);
  const [comment, setComment] = useState('');

  const rateMutation = useMutation({
    mutationFn: async (newRating) => fetch(`${process.env.REACT_APP_API_ENDPOINT}/ratings`, {
      method: 'POST',
      body: JSON.stringify(newRating),
      headers: {
        authorization: `Bearer ${await getAccessTokenSilently()}`,
      },
    }),
    onSuccess: () => queryClient.invalidateQueries(['wacken', user.sub]),
  });

  return (
    <Card sx={{ width: 300 }}>
      <Grid container rowSpacing={1}>
        {imageUrl
          ? (
            <Grid xs={12} sx={{ height: 300 }}>
              <Grid xs={12} display="flex" justifyContent="center" alignItems="center" sx={{ maxHeight: 300 }}>
                <Typography variant="h5">{artistName}</Typography>
              </Grid>
              <Grid xs={12}>
                <CardMedia component="img" src={imageUrl} alt={`${artistName} image`} sx={{ height: 250 }} />
              </Grid>
            </Grid>
          )
          : (
            <Grid xs={12} sx={{ height: 300 }} display="flex" justifyContent="center" alignItems="center">
              <Typography variant="h4">{artistName}</Typography>
            </Grid>
          )}
        <Grid xs={12}>
          <form onSubmit={(event) => {
            event.preventDefault();
            rateMutation.mutate({
              artist_name: artistName,
              festival_name: festival,
              rating,
              year: parseInt(year, 10),
              comment,
            });
          }}
          >
            <Grid xs={12}>
              <TextField fullWidth label="Festival/Concert" value={festival} onChange={(event) => setFestival(event.target.value)} />
            </Grid>
            <Grid xs={12}>
              <TextField fullWidth label="Year" value={year} onChange={(event) => setYear(event.target.value)} />
            </Grid>
            <Grid xs={12}>
              <Rating
                value={rating}
                onChange={(event) => setRating(parseInt(event.target.value, 10))}
              />
            </Grid>
            <Grid xs={12}>
              <TextField fullWidth label="Comment" value={comment} onChange={(event) => setComment(event.target.value)} />
            </Grid>
            <Grid xs={12} display="flex" justifyContent="center" alignItems="center">
              <Button type="submit">Rate</Button>
            </Grid>
          </form>
        </Grid>
      </Grid>

    </Card>
  );
}

RatingCard.propTypes = {
  artistName: PropTypes.string.isRequired,
  imageUrl: PropTypes.string,
};

RatingCard.defaultProps = {
  imageUrl: '',
};
