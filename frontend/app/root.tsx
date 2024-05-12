import {
  Form,
  isRouteErrorResponse,
  Link,
  Links,
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
  ClickAwayListener,
  createTheme,
  CssBaseline,
  List,
  ListItemButton,
  Paper,
  Tab,
  Tabs,
  ThemeProvider,
  Toolbar,
  Typography,
} from "@mui/material";
import { useRef, useState } from "react";
import { Popper } from "@mui/base";

const darkTheme = createTheme({
  palette: {
    mode: "dark",
  },
});

interface AdditionalPath {
  path: string;
  displayName: string;
}

interface Route {
  displayName: string;
  id: string;
  path: string;
  additionalPaths?: AdditionalPath[];
}

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

const routes: Route[] = [
  {
    displayName: "My Ratings",
    id: "ratings",
    path: "/ratings",
  },
  {
    displayName: "Wacken",
    id: "wacken",
    path: "/festivals/wacken",
    additionalPaths: [
      {
        path: "/ratings/wacken",
        displayName: "Overview",
      },
    ],
  },
  {
    displayName: "Dong",
    id: "dong",
    path: "/festivals/dong",
    additionalPaths: [
      {
        path: "/ratings/dong",
        displayName: "Overview",
      },
    ],
  },
  {
    displayName: "RUDE",
    id: "rude",
    path: "/festivals/rude",
    additionalPaths: [
      {
        path: "/ratings/rude",
        displayName: "Overview",
      },
    ],
  },
];

function FestivalTab({ route }: { route: Route }) {
  const [showOptions, setShowOptions] = useState(false);
  const anchorEl = useRef<HTMLButtonElement>(null);
  return (
    <>
      <Tab
        ref={anchorEl}
        component={Button}
        onClick={() =>
          setShowOptions((prevState) => {
            return !prevState;
          })
        }
        label={route.displayName}
      />
      <Popper open={showOptions} anchorEl={anchorEl.current}>
        <Paper>
          <ClickAwayListener onClickAway={() => setShowOptions(false)}>
            <List>
              <ListItemButton
                component={Link}
                to={route.path}
                onClick={() => setShowOptions(false)}
              >
                Rate
              </ListItemButton>
              {route.additionalPaths?.map((a) => (
                <ListItemButton
                  key={a.path}
                  component={Link}
                  to={a.path}
                  onClick={() => setShowOptions(false)}
                >
                  {a.displayName}
                </ListItemButton>
              ))}
            </List>
          </ClickAwayListener>
        </Paper>
      </Popper>
    </>
  );
}

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

  const { pathname } = useLocation();
  let activeTab: number | false = false;
  if (pathname.includes("wacken")) {
    activeTab = 1;
  } else if (pathname.includes("dong")) {
    activeTab = 2;
  } else if (pathname.includes("rude")) {
    activeTab = 3;
  } else if (pathname.includes("ratings")) {
    activeTab = 0;
  }

  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <AppBar position="static">
        <Toolbar>
          <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
            <Tabs value={activeTab}>
              {routes.map((route) => {
                if (route.path.includes("festivals")) {
                  return <FestivalTab key={route.id} route={route} />;
                } else {
                  return (
                    <Tab
                      key={route.id}
                      component={Link}
                      label={route.displayName}
                      to={route.path}
                    />
                  );
                }
              })}
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
