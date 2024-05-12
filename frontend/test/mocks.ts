import { http, HttpResponse } from "msw";
import { setupServer } from "msw/node";
import { testFestivalArtistsData } from "./mock-data/festival";
import { testArtistRatingsData } from "./mock-data/rating";

export const testApi = "http://localhost/api";
const handlers = [
  http.get(`${testApi}/festivals/:name`, () => {
    return HttpResponse.json(testFestivalArtistsData);
  }),
  http.get(`${testApi}/ratings`, () => {
    return HttpResponse.json(testArtistRatingsData);
  }),
  http.post(`${testApi}/ratings`, () => {
    return HttpResponse.json(undefined, { status: 201 });
  }),
  http.get(`${testApi}/ratings/:festivalName`, () => {
    return HttpResponse.json(testArtistRatingsData);
  }),
  http.put(`${testApi}/ratings/:artistName`, () => {
    return HttpResponse.json(undefined, { status: 201 });
  }),
];

const mockServer = setupServer(...handlers);
export default mockServer;
