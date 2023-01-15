import Grid from '@mui/material/Unstable_Grid2';
import { Typography } from '@mui/material';
import { Navigate } from 'react-router-dom';
import { useAuth0 } from '@auth0/auth0-react';

export default function Home() {
  const { isAuthenticated, isLoading } = useAuth0();

  if (isLoading) {
    return <div />;
  }

  return isAuthenticated ? <Navigate to="/wacken" /> : (
    <Grid container justifyContent="center" alignItems="center" justifyItems="start" minHeight="100vh">
      <Grid flexWrap="wrap" justifyContent="center">
        <Typography variant="h1">MusicRating</Typography>
        <Typography variant="h6" align="center">You need to log in in order to rate your music</Typography>
      </Grid>
    </Grid>
  );
}
