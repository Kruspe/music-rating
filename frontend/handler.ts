import { createRequestHandler } from "@remix-run/architect";
import * as build from "./build/server";

export const handler = createRequestHandler({
  //@ts-expect-error build can be undefined if the project was not build before
  build,
});
