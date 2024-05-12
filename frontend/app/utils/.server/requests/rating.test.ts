import mockServer, { testApi } from "../../../../test/mocks";
import { http, HttpResponse } from "msw";
import {
  getFestivalRatings,
  getRatings,
  saveRating,
  updateRating,
} from "~/utils/.server/requests/rating";
import {
  testArtistRatingsData,
  testFestivalName,
} from "../../../../test/mock-data/rating";
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

describe("updateRating", () => {
  test("updates rating", async () => {
    let endpointCalled = false;
    mockServer.use(
      http.put(`${testApi}/ratings/:artistName`, () => {
        endpointCalled = true;
        return HttpResponse.json(undefined, { status: 201 });
      }),
    );
    const response = await updateRating(new Request("http://app.com"), {
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
      http.put(`${testApi}/ratings/:artistName`, () => {
        return HttpResponse.json(
          { error: "Something went wrong" },
          { status: 500 },
        );
      }),
    );
    const response = await updateRating(new Request("http://app.com"), {
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
  test("returns ratings", async () => {
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

describe("getFestivalRatings", () => {
  test("returns ratings for festival", async () => {
    const response = await getFestivalRatings(
      new Request("http://app.com"),
      testFestivalName,
    );

    expect(response).toEqual({
      ok: true,
      data: testArtistRatingsData.map((r) => toArtistRating(r)),
    });
  });

  test("returns error on error", async () => {
    mockServer.use(
      http.get(`${testApi}/ratings/:festivalName`, () => {
        return HttpResponse.json(
          { error: "Something went wrong" },
          { status: 500 },
        );
      }),
    );
    const response = await getFestivalRatings(
      new Request("http://app.com"),
      testFestivalName,
    );

    expect(response).toEqual({
      ok: false,
      error: "Something went wrong",
    });
  });
});
