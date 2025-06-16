import {
  Form,
  isRouteErrorResponse,
  Link,
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useLocation,
  useRouteError,
} from "react-router";
import {
  AppBar,
  Button,
  ClickAwayListener,
  createTheme,
  CssBaseline,
  List,
  ListItemButton,
  Paper,
  Popper,
  Tab,
  Tabs,
  ThemeProvider,
  Toolbar,
  Typography,
} from "@mui/material";
import { useRef, useState } from "react";
import type { Route } from "./+types/root";
import { authenticator } from "~/utils/auth.server";
import { sessionStorage } from "~/utils/session.server";

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

export async function loader({ request }: Route.LoaderArgs) {
  const session = await sessionStorage.getSession(
    request.headers.get("cookie"),
  );
  const user = session.get("user");

  return { isAuthenticated: !!user };
}

export async function action({ request }: Route.ActionArgs) {
  await authenticator.authenticate("auth0", request);
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

export default function App({ loaderData }: Route.ComponentProps) {
  const loggedIn = loaderData && loaderData.isAuthenticated;

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
        <Toolbar
          disableGutters
          sx={{
            justifyContent: "space-between",
            paddingLeft: { xs: 0, md: "16px" },
            paddingRight: { xs: 0, md: "16px" },
          }}
        >
          <Tabs value={activeTab} variant="scrollable" allowScrollButtonsMobile>
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
