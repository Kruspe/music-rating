import { http, HttpResponse } from "msw";
import { setupServer } from "msw/node";
import { testFestivalArtistsData } from "./mock-data/festival";

export const testApi = "http://localhost/api";
const handlers = [
  http.get(`${testApi}/festivals/:name`, () => {
    return HttpResponse.json(testFestivalArtistsData);
  }),
];

const mockServer = setupServer(...handlers);
export default mockServer;
