import { getUnratedFestivalArtists } from "~/utils/.server/requests/festival";
import { testFestivalArtistsData } from "../../../../test/mock-data/festival";
import { toFestivalArtist } from "~/utils/types.server";
import mockServer, { testApi } from "../../../../test/mocks";
import { http, HttpResponse } from "msw";
import { expect } from "vitest";
import { testFestivalName } from "../../../../test/mock-data/rating";

describe("getUnratedFestivalArtists", () => {
  test("returns unrated artists", async () => {
    const response = await getUnratedFestivalArtists(
      new Request("http://app.com"),
      testFestivalName,
    );

    expect(response).toEqual({
      ok: true,
      data: testFestivalArtistsData.map((a) => toFestivalArtist(a)),
    });
  });

  test("returns error when an error occurs", async () => {
    mockServer.use(
      http.get(`${testApi}/festivals/:name`, () => {
        return HttpResponse.json(
          { error: "Something went wrong" },
          { status: 500 },
        );
      }),
    );
    const response = await getUnratedFestivalArtists(
      new Request("http://app.com"),
      testFestivalName,
    );

    expect(response).toEqual({
      ok: false,
      error: "Something went wrong",
    });
  });
});
