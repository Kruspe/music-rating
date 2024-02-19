import { Authenticator } from "remix-auth";
import { Auth0Strategy } from "remix-auth-auth0";

import { sessionStorage } from "~/utils/session.server";

import { testApi } from "../../test/mocks";

type User = {
  id: string;
  token: string;
};

export const authenticator = new Authenticator<User>(sessionStorage);

let apiEndpoint = testApi;
let clientId = "client-id";
let clientSecret = "client-secret";
let domainName = "domain-name";

if (process.env.NODE_ENV === "production") {
  const { API_ENDPOINT, CLIENT_ID, CLIENT_SECRET, DOMAIN_NAME } = process.env;
  if (!API_ENDPOINT || !CLIENT_ID || !CLIENT_SECRET || !DOMAIN_NAME) {
    console.error("Environment variables not set correctly!!");
    throw new Error("Environment variables not set correctly!!");
  }
  apiEndpoint = API_ENDPOINT;
  clientId = CLIENT_ID;
  clientSecret = CLIENT_SECRET;
  domainName = DOMAIN_NAME;
}
const auth0Strategy = new Auth0Strategy(
  {
    callbackURL: `${domainName}/auth/callback`,
    clientID: clientId,
    clientSecret: clientSecret,
    domain: "musicrating.eu.auth0.com",
    audience: process.env.NODE_ENV === "production" ? apiEndpoint : undefined,
  },
  async ({ accessToken, refreshToken, extraParams, profile }) => {
    return { id: profile._json!.sub!, token: accessToken };
  },
);

authenticator.use(auth0Strategy);
