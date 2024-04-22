import mockServer, { testApi } from "../../../../test/mocks";
import { http, HttpResponse } from "msw";
import { saveRating } from "~/utils/.server/requests/rating";

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
