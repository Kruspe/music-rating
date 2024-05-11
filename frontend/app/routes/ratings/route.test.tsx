import * as ratingRequests from "~/utils/.server/requests/rating";
import { action, loader } from "~/routes/ratings/route";
import {
  testArtistName,
  testArtistRatingsData,
} from "../../../test/mock-data/artist";
import { toArtistRating } from "~/utils/types.server";
import { RatingRequest } from "~/utils/.server/requests/rating";
import { testFestivalName } from "../../../test/mock-data/festival";
import mockServer, { testApi } from "../../../test/mocks";
import { http, HttpResponse } from "msw";

describe("loader", () => {
  test("loads ratings", async () => {
    const getRatingsSpy = vi.spyOn(ratingRequests, "getRatings");
    const response = await loader({
      request: new Request("http://app.com"),
      params: {},
      context: {},
    });
    const responseData = await response.json();

    expect(getRatingsSpy).toHaveBeenCalledTimes(1);
    expect(getRatingsSpy).toHaveBeenCalledWith(expect.anything());
    expect(responseData).toEqual({
      ok: true,
      data: testArtistRatingsData.map((r) => toArtistRating(r)),
    });
  });

  test("throws error if festival is not supported", async () => {
    const errorMessage = "Error loading ratings";
    mockServer.use(
      http.get(`${testApi}/ratings`, () => {
        return HttpResponse.json({ error: errorMessage }, { status: 500 });
      }),
    );
    let errorData;
    try {
      await loader({
        request: new Request("http://app.com"),
        params: {},
        context: {},
      });
    } catch (error) {
      errorData = await (error as Response).json();
    }
    expect(errorData).toEqual(errorMessage);
  });
});

describe("action", () => {
  test("saves rating", async () => {
    const newRatingRequest: RatingRequest = {
      artist_name: testArtistName,
      festival_name: testFestivalName,
      rating: 5,
      year: 2015,
      comment: "Old school swedish death metal",
    };
    const formData = new FormData();
    formData.append("artist_name", newRatingRequest.artist_name);
    formData.append("festival_name", newRatingRequest.festival_name);
    formData.append("rating", newRatingRequest.rating.toString());
    formData.append("year", newRatingRequest.year.toString());
    formData.append("comment", newRatingRequest.comment);

    const saveRatingSpy = vi.spyOn(ratingRequests, "saveRating");
    const response = await action({
      request: new Request("http://app.com", {
        method: "POST",
        body: formData,
      }),
      params: {},
      context: {},
    });

    expect(saveRatingSpy).toHaveBeenCalledTimes(1);
    expect(saveRatingSpy).toHaveBeenCalledWith(expect.anything(), {
      artist_name: newRatingRequest.artist_name,
      festival_name: newRatingRequest.festival_name,
      rating: newRatingRequest.rating,
      year: newRatingRequest.year,
      comment: newRatingRequest.comment,
    });
    expect(response).toEqual({ ok: true });
  });
});
