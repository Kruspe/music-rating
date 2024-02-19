import { Authenticator } from "remix-auth";
import { Auth0Strategy } from "remix-auth-auth0";

import { sessionStorage } from "~/utils/session.server";

type User = {
  id: string;
  token: string;
};

export const authenticator = new Authenticator<User>(sessionStorage);

const { API_ENDPOINT, CLIENT_ID, CLIENT_SECRET, DOMAIN_NAME } = process.env;
if (!API_ENDPOINT || !CLIENT_ID || !CLIENT_SECRET || !DOMAIN_NAME) {
  console.error("Environment variables not set correctly!!");
  throw new Error("Environment variables not set correctly!!");
}
const auth0Strategy = new Auth0Strategy(
  {
    callbackURL: `${DOMAIN_NAME}/auth/callback`,
    clientID: CLIENT_ID,
    clientSecret: CLIENT_SECRET,
    domain: "musicrating.eu.auth0.com",
    audience: process.env.NODE_ENV === "production" ? API_ENDPOINT : undefined,
  },
  async ({ accessToken, refreshToken, extraParams, profile }) => {
    return { id: profile._json!.sub!, token: accessToken };
  },
);

authenticator.use(auth0Strategy);
