import {
  Form,
  Link,
  LiveReload,
  Outlet,
  Scripts,
  ScrollRestoration,
  useLoaderData,
  useLocation,
} from "@remix-run/react";
import {
  AppBar,
  Box,
  Button,
  createTheme,
  CssBaseline,
  Tab,
  Tabs,
  ThemeProvider,
  Toolbar,
} from "@mui/material";
import type { ActionArgs, LoaderArgs } from "@remix-run/node";
import { authenticator } from "~/utils/auth.server";
import type { ReactNode } from "react";
import { matchPath } from "@remix-run/router";

const darkTheme = createTheme({
  palette: {
    mode: "dark",
  },
});

interface DocumentProps {
  children: ReactNode;
}

const Document = ({ children }: DocumentProps) => {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width,initial-scale=1" />
        <title>MusicRating</title>
      </head>
      <body>
        {children}
        <ScrollRestoration />
        <Scripts />
        <LiveReload />
      </body>
    </html>
  );
};

export async function loader({ request }: LoaderArgs) {
  return authenticator.isAuthenticated(request);
}

export async function action({ request }: ActionArgs) {
  return authenticator.authenticate("auth0", request);
}

function useRouteMatch(patterns: readonly string[]) {
  const { pathname } = useLocation();

  for (let i = 0; i < patterns.length; i++) {
    const possibleMatch = matchPath(patterns[i], pathname);
    if (possibleMatch !== null) {
      return possibleMatch;
    }
  }
  return null;
}

const routes = ["/ratings", "/wacken"];

export default function App() {
  const data = useLoaderData<typeof loader>();
  const loggedIn = data && data.id;

  const currentTab = useRouteMatch(routes)?.pattern.path;

  return (
    <Document>
      <ThemeProvider theme={darkTheme}>
        <CssBaseline />
        <AppBar position="static">
          <Toolbar>
            <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
              <Tabs value={currentTab}>
                <Tab
                  component={Link}
                  value="/ratings"
                  label="My Ratings"
                  to="/ratings"
                />
                <Tab
                  component={Link}
                  value="/wacken"
                  label="Wacken"
                  to="/wacken"
                />
              </Tabs>
            </Box>
            <Box sx={{ flexGrow: 1 }} />
            <Form action={loggedIn ? "/logout" : ""}>
              <Button variant="contained" type="submit" formMethod="post">
                {loggedIn ? "Log out" : "Log in"}
              </Button>
            </Form>
          </Toolbar>
        </AppBar>
        <Outlet />
      </ThemeProvider>
    </Document>
  );
}
