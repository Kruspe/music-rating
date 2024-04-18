import { Typography, Unstable_Grid2 as Grid } from "@mui/material";
import type { DataFunctionArgs } from "@remix-run/node";
import { authenticator } from "~/utils/auth.server";

export async function loader({ request }: DataFunctionArgs) {
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
