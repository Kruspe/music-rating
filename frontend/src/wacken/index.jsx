import { useQuery } from '@tanstack/react-query';
import { Typography } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { useAuth0 } from '@auth0/auth0-react';
import RatingCard from '../rating/RatingCard';

export default function Wacken() {
  const { getAccessTokenSilently, user } = useAuth0();

  const { data: bands, isFetched } = useQuery(['wacken', user.sub], async () => {
    const result = await fetch(`${process.env.REACT_APP_API_ENDPOINT}/festivals/wacken`, {
      headers: {
        authorization: `Bearer ${(await getAccessTokenSilently())}`,
      },
    });
    return result.json();
  });

  return isFetched && bands.length > 0
    ? (
      <Grid container spacing={0.5}>
        {bands.map((band) => (
          <Grid key={band.artist_name}>
            <RatingCard artistName={band.artist_name} imageUrl={band.image_url} />
          </Grid>
        ))}
      </Grid>
    )
    : <Typography variant="h2">You rated all bands that are announced at the moment</Typography>;
}
