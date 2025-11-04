import * as ratingRequests from "~/utils/.server/requests/rating";
import FestivalRatingsRoute, { loader } from "~/routes/ratings/festival";
import {
  testArtistRatingsData,
  testFestivalName,
} from "../../../test/mock-data/rating";
import { toArtistRating } from "~/utils/types.server";
import mockServer, { testApi } from "../../../test/mocks";
import { http, HttpResponse } from "msw";
import { createRoutesStub, data } from "react-router";
import { render, screen } from "@testing-library/react";

describe("loader", () => {
  test("loads festival artists and ratings", async () => {
    const getFestivalRatingsSpy = vi.spyOn(
      ratingRequests,
      "getFestivalRatings",
    );
    const response = await loader({
      request: new Request("http://app.com"),
      params: { festivalName: testFestivalName },
      context: {},
      unstable_pattern: "",
    });

    expect(getFestivalRatingsSpy).toHaveBeenCalledTimes(1);
    expect(getFestivalRatingsSpy).toHaveBeenCalledWith(
      expect.anything(),
      testFestivalName,
    );
    expect(response.data).toEqual({
      ok: true,
      data: testArtistRatingsData.map((r) => toArtistRating(r)),
    });
  });

  test("throws error if festival is not supported", async () => {
    const errorMessage = "Error loading ratings";
    mockServer.use(
      http.get(`${testApi}/ratings/:festivalName`, () => {
        return HttpResponse.json({ error: errorMessage }, { status: 500 });
      }),
    );
    let errorData;
    try {
      await loader({
        request: new Request("http://app.com"),
        params: { festivalName: testFestivalName },
        context: {},
        unstable_pattern: "",
      });
    } catch (error) {
      errorData = error;
    }
    expect(errorData).toMatchObject({ data: errorMessage });
  });
});

test("shows all artists from festival with ratings", async () => {
  const festivalRatings = testArtistRatingsData.map((r) => toArtistRating(r));
  const RemixStub = createRoutesStub([
    {
      path: "/ratings/:festivalName",
      // @ts-expect-error Type error by react-router (https://github.com/remix-run/react-router/issues/13579)
      Component: FestivalRatingsRoute,
      loader: async () => {
        return data({
          ok: true,
          data: festivalRatings,
        });
      },
    },
  ]);
  render(<RemixStub initialEntries={[`/ratings/${testFestivalName}`]} />);

  expect(await screen.findByText(festivalRatings[0].artistName)).toBeVisible();
  expect(await screen.findByText(festivalRatings[1].artistName)).toBeVisible();
  expect(await screen.findByText(festivalRatings[2].artistName)).toBeVisible();
});
