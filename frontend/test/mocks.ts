import { http, HttpResponse } from "msw";
import { setupServer } from "msw/node";
import { testFestivalArtistsData } from "./mock-data/festival";
import { testArtistRatingsData } from "./mock-data/artist";

export const testApi = "http://localhost/api";
const handlers = [
  http.get(`${testApi}/festivals/:name`, () => {
    return HttpResponse.json(testFestivalArtistsData);
  }),
  http.get(`${testApi}/ratings`, () => {
    return HttpResponse.json(testArtistRatingsData);
  }),
];

const mockServer = setupServer(...handlers);
export default mockServer;
