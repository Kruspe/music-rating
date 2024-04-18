import { Typography, Unstable_Grid2 as Grid } from "@mui/material";
import { authenticator } from "~/utils/auth.server";
import { LoaderFunctionArgs } from "@remix-run/node";

export async function loader({ request }: LoaderFunctionArgs) {
  return authenticator.isAuthenticated(request, {
    successRedirect: "/ratings",
  });
}

export default function IndexRoute() {
  return (
    <Grid
      container
      justifyContent="center"
      alignItems="center"
      justifyItems="start"
      minHeight="100vh"
    >
      <Grid flexWrap="wrap" justifyContent="center">
        <Typography variant="h1">MusicRating</Typography>
        <Typography variant="h6" align="center">
          You need to log in in order to rate your music
        </Typography>
      </Grid>
    </Grid>
  );
}
