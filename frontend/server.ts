import { createRequestHandler } from "@remix-run/architect";
import * as build from "./build";

export const handler = createRequestHandler({
  build,
  getLoadContext(event) {
    // use lambda event to generate a context for loaders
    return {};
  },
});
