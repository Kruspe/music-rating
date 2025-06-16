import { Grid, Typography } from "@mui/material";
import { sessionStorage } from "~/utils/session.server";
import { redirect } from "react-router";
import type { Route } from "./+types/home";

export async function loader({ request }: Route.LoaderArgs) {
  const session = await sessionStorage.getSession(
    request.headers.get("cookie"),
  );
  const user = session.get("user");
  console.log(user);
  if (user) {
    return redirect("/ratings");
  }
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
