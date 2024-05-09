import mockServer, { testApi } from "../../../../test/mocks";
import { http, HttpResponse } from "msw";
import { getRatings, saveRating } from "~/utils/.server/requests/rating";
import { testArtistRatingsData } from "../../../../test/mock-data/artist";
import { toArtistRating } from "~/utils/types.server";

describe("saveRating", () => {
  test("saves rating", async () => {
    let endpointCalled = false;
    mockServer.use(
      http.post(`${testApi}/ratings`, () => {
        endpointCalled = true;
        return HttpResponse.json(undefined, { status: 201 });
      }),
    );
    const response = await saveRating(new Request("http://app.com"), {
      artist_name: "Bloodbath",
      festival_name: "Wacken",
      rating: 5,
      year: 2023,
      comment: "Swedish Death Metal",
    });

    expect(endpointCalled).toBeTruthy();
    expect(response).toEqual({
      ok: true,
    });
  });

  test("returns error on error", async () => {
    mockServer.use(
      http.post(`${testApi}/ratings`, () => {
        return HttpResponse.json(
          { error: "Something went wrong" },
          { status: 500 },
        );
      }),
    );
    const response = await saveRating(new Request("http://app.com"), {
      artist_name: "Bloodbath",
      festival_name: "Wacken",
      rating: 5,
      year: 2023,
      comment: "Swedish Death Metal",
    });

    expect(response).toEqual({
      ok: false,
      error: "Something went wrong",
    });
  });
});

describe("getRatings", () => {
  test("saves rating", async () => {
    const response = await getRatings(new Request("http://app.com"));

    expect(response).toEqual({
      ok: true,
      data: testArtistRatingsData.map((r) => toArtistRating(r)),
    });
  });

  test("returns error on error", async () => {
    mockServer.use(
      http.get(`${testApi}/ratings`, () => {
        return HttpResponse.json(
          { error: "Something went wrong" },
          { status: 500 },
        );
      }),
    );
    const response = await getRatings(new Request("http://app.com"));

    expect(response).toEqual({
      ok: false,
      error: "Something went wrong",
    });
  });
});
