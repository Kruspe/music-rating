import * as ratingRequests from "~/utils/.server/requests/rating";
import { type RatingRequest } from "~/utils/.server/requests/rating";
import RatingsRoute, { action, loader } from "~/routes/ratings/home";
import {
  testArtistName,
  testArtistRatingData,
  testArtistRatingsData,
  testFestivalName,
} from "../../../test/mock-data/rating";
import { toArtistRating } from "~/utils/types.server";
import mockServer, { testApi } from "../../../test/mocks";
import { http, HttpResponse } from "msw";
import { createRoutesStub, data } from "react-router";
import { render, screen } from "@testing-library/react";
import { userEvent } from "@testing-library/user-event";

describe("loader", () => {
  test("loads ratings", async () => {
    const getRatingsSpy = vi.spyOn(ratingRequests, "getRatings");
    const response = await loader({
      request: new Request("http://app.com"),
      params: {},
      context: {},
    });

    expect(getRatingsSpy).toHaveBeenCalledTimes(1);
    expect(getRatingsSpy).toHaveBeenCalledWith(expect.anything());
    expect(response.data).toEqual({
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
      errorData = error;
    }
    expect(errorData).toMatchObject({ data: errorMessage });
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
    formData.append("festival_name", newRatingRequest.festival_name!);
    formData.append("rating", newRatingRequest.rating.toString());
    formData.append("year", newRatingRequest.year!.toString());
    formData.append("comment", newRatingRequest.comment!);
    formData.append("_action", "SAVE_RATING");

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

  test("updates rating", async () => {
    const updatedRatingRequest: RatingRequest = {
      artist_name: testArtistName,
      festival_name: testFestivalName,
      rating: 5,
      year: 2015,
      comment: "Old school swedish death metal",
    };
    const formData = new FormData();
    formData.append("artist_name", updatedRatingRequest.artist_name);
    formData.append("festival_name", updatedRatingRequest.festival_name!);
    formData.append("rating", updatedRatingRequest.rating.toString());
    formData.append("year", updatedRatingRequest.year!.toString());
    formData.append("comment", updatedRatingRequest.comment!);
    formData.append("_action", "UPDATE_RATING");

    const saveRatingSpy = vi.spyOn(ratingRequests, "updateRating");
    const response = await action({
      request: new Request("http://app.com", {
        method: "PUT",
        body: formData,
      }),
      params: {},
      context: {},
    });

    expect(saveRatingSpy).toHaveBeenCalledTimes(1);
    expect(saveRatingSpy).toHaveBeenCalledWith(expect.anything(), {
      artist_name: updatedRatingRequest.artist_name,
      festival_name: updatedRatingRequest.festival_name,
      rating: updatedRatingRequest.rating,
      year: updatedRatingRequest.year,
      comment: updatedRatingRequest.comment,
    });
    expect(response).toEqual({ ok: true });
  });
});

test("shows all rated bands", async () => {
  const ratings = testArtistRatingsData.map((r) => toArtistRating(r));
  const RemixStub = createRoutesStub([
    {
      path: "/ratings",
      // @ts-expect-error Type error by react-router (https://github.com/remix-run/react-router/issues/13579)
      Component: RatingsRoute,
      loader: async () => {
        return data({
          ok: true,
          data: ratings,
        });
      },
    },
  ]);
  render(<RemixStub initialEntries={["/ratings"]} />);

  for (const rating of ratings) {
    expect(await screen.findByText(rating.artistName)).toBeVisible();
  }
});

test("can update rating", async () => {
  const user = userEvent.setup();
  const rating = toArtistRating(testArtistRatingData);
  const updatedFestivalName = "Dong";
  const RemixStub = createRoutesStub([
    {
      path: "/ratings",
      // @ts-expect-error Type error by react-router (https://github.com/remix-run/react-router/issues/13579)
      Component: RatingsRoute,
      loader: async () => {
        return data({
          ok: true,
          data: [rating],
        });
      },
      action: async ({ request }) => {
        const formData = await request.formData();
        expect(formData.get("_action")).toEqual("UPDATE_RATING");
        expect(formData.get("artist_name")).toEqual(rating.artistName);
        expect(formData.get("festival_name")).toEqual(updatedFestivalName);
        expect(formData.get("year")).toEqual(rating.year?.toString());
        expect(formData.get("rating")).toEqual(rating.rating.toString());
        expect(formData.get("comment")).toEqual(rating.comment);
        return data({ ok: true });
      },
    },
  ]);
  render(<RemixStub initialEntries={["/ratings"]} />);

  await user.dblClick(await screen.findByText(rating.festivalName!));
  const festivalInputField = screen.getByDisplayValue(rating.festivalName!);
  await user.clear(festivalInputField);
  await user.type(festivalInputField, "Dong");
  await user.tab();

  expect(screen.queryByText(/AssertionError/i)).not.toBeInTheDocument();
});
