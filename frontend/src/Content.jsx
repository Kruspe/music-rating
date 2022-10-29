import { useAuth0 } from '@auth0/auth0-react';
import { QueryClientProvider, useQueryClient } from '@tanstack/react-query';
import Grid from '@mui/material/Unstable_Grid2';
import {
  Button, Card, CardContent, Typography,
} from '@mui/material';

export default function Content() {
  const queryClient = useQueryClient();
  const { isAuthenticated, loginWithRedirect } = useAuth0();

  return (
    isAuthenticated ? (
      <QueryClientProvider client={queryClient}>
        <div>Hello World</div>
      </QueryClientProvider>
    ) : (
      <Grid container justifyContent="center" alignItems="center" minHeight="100vh">
        <Card>
          <CardContent>
            <Typography>You need to log in in order to rate your music</Typography>
            <Button onClick={loginWithRedirect}>Log in</Button>
          </CardContent>
        </Card>
      </Grid>
    ));
}
