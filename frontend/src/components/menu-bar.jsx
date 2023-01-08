import {
  AppBar, Box, Button, Toolbar,
} from '@mui/material';
import { useAuth0 } from '@auth0/auth0-react';

export default function MenuBar() {
  const { isAuthenticated, loginWithRedirect, logout } = useAuth0();

  return (
    <AppBar position="static">
      <Toolbar>
        <Box sx={{ flexGrow: 1 }} />
        {isAuthenticated
          ? (
            <Button variant="contained" onClick={logout}>Log out</Button>
          ) : (
            <Button variant="contained" onClick={loginWithRedirect}>Log in</Button>
          )}
      </Toolbar>
    </AppBar>
  );
}
