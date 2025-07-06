import { redirect } from "react-router";
import { sessionStorage } from "~/utils/session.server";

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
  const user = session.get("user");
  if (!user) {
    throw redirect("/");
  }

  headers.set("authorization", `Bearer ${user.token}`);
  return headers;
}

export function hasError(response: Response) {
  return response.status >= 400;
}
