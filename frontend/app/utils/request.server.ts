import type { TypedResponse } from "@remix-run/node";
import { json } from "@remix-run/node";
import { authenticator } from "~/utils/auth.server";

export const get = async <T>(
  request: Request,
  path: string
): Promise<TypedResponse<T>> => {
  const { token } = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });
  const response = await fetch(`${process.env.API_ENDPOINT}${path}`, {
    headers: {
      authorization: `Bearer ${token}`,
    },
  });
  return json(await response.json());
};
