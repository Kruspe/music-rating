import { createCookieSessionStorage } from "@remix-run/node";

let sessionSecret = "session-secret";
if (process.env.NODE_ENV === "production") {
  const { SESSION_SECRET } = process.env;
  if (!SESSION_SECRET) {
    console.error("SESSION_SECRET missing from environment variables");
    throw new Error("SESSION_SECRET missing from environment variables");
  }
  sessionSecret = SESSION_SECRET;
}

export const sessionStorage = createCookieSessionStorage({
  cookie: {
    name: "MR_session",
    sameSite: "lax",
    path: "/",
    httpOnly: true,
    secrets: [sessionSecret],
    secure: process.env.NODE_ENV === "production",
  },
});
