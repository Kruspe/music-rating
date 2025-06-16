import { sessionStorage } from "~/utils/session.server";
import type { Route } from "./+types/logout";
import { redirect } from "react-router";

export const action = async ({ request }: Route.ActionArgs) => {
  const session = await sessionStorage.getSession(
    request.headers.get("cookie"),
  );
  return redirect("/", {
    headers: { "Set-Cookie": await sessionStorage.destroySession(session) },
  });
};
