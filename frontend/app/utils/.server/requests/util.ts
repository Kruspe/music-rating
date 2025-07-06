import { redirect } from "react-router";
import { sessionStorage } from "~/utils/session.server";
import { auth0Strategy, type User } from "~/utils/auth.server";

export interface FetchResponse<T = void> {
  data?: T;
  error?: string;
  ok: boolean;
}

export interface ErrorResponseData {
  error: string;
}

export async function createAuthHeader(request: Request) {
  const headers = new Headers();
  const session = await sessionStorage.getSession(
    request.headers.get("cookie"),
  );
  const user: User = session.get("user");
  if (!user) {
    throw redirect("/");
  }
  if (user.expiry && user.refreshToken && Date.now() >= user.expiry) {
    const tokens = await auth0Strategy.refreshToken(user.refreshToken);
    session.set("user", {
      ...user,
      token: tokens.accessToken,
      refreshToken: tokens.refreshToken,
      expiry: tokens.accessTokenExpiresAt(),
    });
    headers.set("set-cookie", await sessionStorage.commitSession(session));
  }

  headers.set("authorization", `Bearer ${user.token}`);
  return headers;
}

export function hasError(response: Response) {
  return response.status >= 400;
}
