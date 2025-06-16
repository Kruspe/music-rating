import { Authenticator } from "remix-auth";
import { Auth0Strategy } from "remix-auth-auth0";

type User = {
  id: string;
  token: string;
};

export const authenticator = new Authenticator<User>();

const { API_ENDPOINT, CLIENT_ID, CLIENT_SECRET, DOMAIN_NAME } = process.env;
if (!API_ENDPOINT || !CLIENT_ID || !CLIENT_SECRET || !DOMAIN_NAME) {
  console.error("Environment variables not set correctly!!");
  throw new Error("Environment variables not set correctly!!");
}
const auth0Strategy = new Auth0Strategy(
  {
    redirectURI: `${DOMAIN_NAME}/auth/callback`,
    clientId: CLIENT_ID,
    clientSecret: CLIENT_SECRET,
    domain: "musicrating.eu.auth0.com",
    scopes: ["openid"],
    audience: process.env.NODE_ENV === "production" ? API_ENDPOINT : undefined,
  },
  async ({ tokens }) => {
    const userId = JSON.parse(
      Buffer.from(tokens.idToken().split(".")[1], "base64").toString(),
    ).sub;
    return { id: userId, token: tokens.accessToken() };
  },
);

authenticator.use(auth0Strategy);
