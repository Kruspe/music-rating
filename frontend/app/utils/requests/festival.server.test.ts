import { http, HttpResponse } from "msw";
import { expect } from "vitest";

import { getFestivalArtists } from "~/utils/requests/festival.server";

import { testFestivalName } from "../../../test/mock-data/festival";
import { errorMessage } from "../../../test/mock-data/generic";
import mockServer, { testApi } from "../../../test/mocks";

describe("getFestivalArtists", () => {
  test("returns artists for the festival", async () => {
    const festivalArtists = await getFestivalArtists(
      new Request("http://app.com"),
      testFestivalName,
    );

    expect(festivalArtists.ok).toBeTruthy();
    expect(festivalArtists.data?.artists).toHaveLength(3);
  });

  test("returns error", async () => {
    mockServer.use(
      http.get(`${testApi}/festivals/${testFestivalName}`, () => {
        return new HttpResponse(errorMessage, { status: 500 });
      }),
    );
    const response = await getFestivalArtists(
      new Request("http://app.com"),
      testFestivalName,
    );

    expect(response.ok).toBeFalsy();
    expect(response.error).toBe(errorMessage);
  });
});
