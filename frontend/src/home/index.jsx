import Grid from '@mui/material/Unstable_Grid2';
import { Typography } from '@mui/material';

export default function Home() {
  return (
    <Grid container justifyContent="center" alignItems="center" justifyItems="start" minHeight="100vh">
      <Grid flexWrap="wrap" justifyContent="center">
        <Typography variant="h1">MusicRating</Typography>
        <Typography variant="h6" align="center">You need to log in in order to rate your music</Typography>
      </Grid>
    </Grid>
  );
}
