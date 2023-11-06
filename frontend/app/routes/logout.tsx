import type { DataFunctionArgs } from "@remix-run/node";

import { authenticator } from "~/utils/auth.server";

export const action = async ({ request }: DataFunctionArgs) => {
  await authenticator.logout(request, { redirectTo: "/" });
};
