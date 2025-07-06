import { Authenticator } from "remix-auth";
import { Auth0Strategy } from "remix-auth-auth0";

export type User = {
  id: string;
  token: string;
  // after some time these should be available everywhere
  refreshToken?: string;
  expiry?: number;
};

export const authenticator = new Authenticator<User>();

const { API_ENDPOINT, CLIENT_ID, CLIENT_SECRET, DOMAIN_NAME } = process.env;
if (!API_ENDPOINT || !CLIENT_ID || !CLIENT_SECRET || !DOMAIN_NAME) {
  console.error("Environment variables not set correctly!!");
  throw new Error("Environment variables not set correctly!!");
}
export const auth0Strategy = new Auth0Strategy(
  {
    redirectURI: `${DOMAIN_NAME}/auth/callback`,
    clientId: CLIENT_ID,
    clientSecret: CLIENT_SECRET,
    domain: "musicrating.eu.auth0.com",
    scopes: ["openid", "offline_access"],
    audience: API_ENDPOINT,
  },
  async ({ tokens }): Promise<User> => {
    const userId = JSON.parse(
      Buffer.from(tokens.idToken().split(".")[1], "base64").toString(),
    ).sub;
    return {
      id: userId,
      token: tokens.accessToken(),
      expiry: tokens.accessTokenExpiresAt().getTime(),
      refreshToken: tokens.refreshToken(),
    };
  },
);

authenticator.use(auth0Strategy);
