import './App.css';
import { useAuth0 } from '@auth0/auth0-react';
import {
  Button, Card, CardContent, Typography,
} from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import Wacken from './wacken';

function App() {
  const { isAuthenticated, loginWithRedirect } = useAuth0();
  return (
    isAuthenticated ? <Wacken /> : (
      <Grid container justifyContent="center" alignItems="center" minHeight="100vh">
        <Card>
          <CardContent>
            <Typography>You need to log in in order to rate your music</Typography>
            <Button onClick={loginWithRedirect}>Log in</Button>
          </CardContent>
        </Card>
      </Grid>
    )
  );
}

export default App;
