import { authenticator } from "~/utils/auth.server";

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
  const { token } = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });
  headers.set("authorization", `Bearer ${token}`);
  return headers;
}

export function hasError(response: Response) {
  return response.status >= 400;
}
