import * as ratingRequests from "~/utils/.server/requests/rating";
import FestivalRatingsRoute, {
  loader,
} from "~/routes/ratings.$festivalName/route";
import {
  testArtistRatingsData,
  testFestivalName,
} from "../../../test/mock-data/rating";
import { ArtistRating, toArtistRating } from "~/utils/types.server";
import mockServer, { testApi } from "../../../test/mocks";
import { http, HttpResponse } from "msw";
import { createRemixStub } from "@remix-run/testing";
import { json, TypedResponse } from "@remix-run/node";
import { FetchResponse } from "~/utils/.server/requests/util";
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
    });
    const responseData = await response.json();

    expect(getFestivalRatingsSpy).toHaveBeenCalledTimes(1);
    expect(getFestivalRatingsSpy).toHaveBeenCalledWith(
      expect.anything(),
      testFestivalName,
    );
    expect(responseData).toEqual({
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
        params: {},
        context: {},
      });
    } catch (error) {
      errorData = await (error as Response).json();
    }
    expect(errorData).toEqual(errorMessage);
  });
});

test("shows all artists from festival with ratings", async () => {
  const festivalRatings = testArtistRatingsData.map((r) => toArtistRating(r));
  const RemixStub = createRemixStub([
    {
      path: "/ratings/:festivalName",
      Component: FestivalRatingsRoute,
      loader: async (): Promise<
        TypedResponse<FetchResponse<ArtistRating[]>>
      > => {
        return json({
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
