import { http, HttpResponse } from "msw";
import { setupServer } from "msw/node";

export const testApi = "http://localhost/api";
const handlers = [
  http.get(`${testApi}/festivals/wacken`, () => {
    return HttpResponse.json({});
  }),
];

const mockServer = setupServer(...handlers);
export default mockServer;
