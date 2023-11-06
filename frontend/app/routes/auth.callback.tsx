import type { DataFunctionArgs } from "@remix-run/node";

import { authenticator } from "~/utils/auth.server";

export let loader = ({ request }: DataFunctionArgs) => {
  return authenticator.authenticate("auth0", request, {
    successRedirect: "/ratings",
  });
};
