import type { Route } from "./+types/auth.callback";
import { authenticator } from "~/utils/auth.server";
import { redirect } from "react-router";
import { sessionStorage } from "~/utils/session.server";

export async function loader({ request }: Route.LoaderArgs) {
  const user = await authenticator.authenticate("auth0", request);
  const session = await sessionStorage.getSession(
    request.headers.get("cookie"),
  );
  session.set("user", user);
  return redirect("/", {
    headers: {
      "Set-Cookie": await sessionStorage.commitSession(session),
    },
  });
}
