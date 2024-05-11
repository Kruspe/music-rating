import {
  Form,
  isRouteErrorResponse,
  Link,
  Links,
  matchPath,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useLoaderData,
  useLocation,
  useRouteError,
} from "@remix-run/react";
import { ActionFunctionArgs, LoaderFunctionArgs } from "@remix-run/node";
import { authenticator } from "~/utils/auth.server";
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
  Typography,
} from "@mui/material";

const darkTheme = createTheme({
  palette: {
    mode: "dark",
  },
});

export function ErrorBoundary() {
  const error = useRouteError();

  if (isRouteErrorResponse(error)) {
    return <Typography variant="h3">{error.data}</Typography>;
  }
}

export async function loader({ request }: LoaderFunctionArgs) {
  return authenticator.isAuthenticated(request);
}

export async function action({ request }: ActionFunctionArgs) {
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

const routes = ["/ratings", "/festivals/wacken", "/festivals/dong"];

export function Layout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <title>MusicRating</title>
        <Meta />
        <Links />
      </head>
      <body>
        {children}
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  );
}

export default function App() {
  const data = useLoaderData<typeof loader>();
  const loggedIn = data && data.id;

  const currentTab = useRouteMatch(routes)?.pattern.path ?? false;

  return (
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
                value="/festivals/wacken"
                label="Wacken"
                to="/festivals/wacken"
              />
              <Tab
                component={Link}
                value="/festivals/dong"
                label="Dong"
                to="/festivals/dong"
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
  );
}
