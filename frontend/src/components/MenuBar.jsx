import {
  AppBar, Box, Button, Tab, Tabs, Toolbar,
} from '@mui/material';
import { useAuth0 } from '@auth0/auth0-react';
import { useEffect, useState } from 'react';
import { Link, Outlet, useLocation } from 'react-router-dom';

export default function MenuBar() {
  const { isAuthenticated, loginWithRedirect, logout } = useAuth0();
  const [tabIndex, setTabIndex] = useState(0);
  const location = useLocation();

  useEffect(() => {
    if (location.pathname === '/wacken') {
      setTabIndex(1);
    }
  });

  return (
    <>
      <AppBar position="static">
        <Toolbar>
          <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
            <Tabs
              value={tabIndex}
              onChange={(event, newValue) => setTabIndex(newValue)}
            >
              <Tab component={Link} label="My Ratings" to="/ratings" />
              <Tab component={Link} label="Wacken" to="/wacken" />
            </Tabs>
          </Box>
          <Box sx={{ flexGrow: 1 }} />
          {isAuthenticated
            ? (
              <Button variant="contained" onClick={logout}>Log out</Button>
            ) : (
              <Button variant="contained" onClick={loginWithRedirect}>Log in</Button>
            )}
        </Toolbar>
      </AppBar>
      <Outlet />
    </>
  );
}
