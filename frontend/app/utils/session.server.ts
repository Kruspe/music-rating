import { createCookieSessionStorage } from "@remix-run/node";

const { SESSION_SECRET } = process.env;
if (!SESSION_SECRET) {
  console.error("SESSION_SECRET missing from environment variables");
  throw new Error("SESSION_SECRET missing from environment variables");
}
export const sessionStorage = createCookieSessionStorage({
  cookie: {
    name: "MR_session",
    sameSite: "lax",
    path: "/",
    httpOnly: true,
    secrets: [SESSION_SECRET],
    secure: process.env.NODE_ENV === "production",
  },
});
