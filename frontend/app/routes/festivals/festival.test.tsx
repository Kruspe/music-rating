import * as festivalRequests from "~/utils/.server/requests/festival";
import FestivalRoute, { loader } from "~/routes/festivals/festival";
import { testFestivalArtistsData } from "../../../test/mock-data/festival";
import { toFestivalArtist } from "~/utils/types.server";
import { createRoutesStub, data } from "react-router";
import { fireEvent, render, screen } from "@testing-library/react";
import { type RatingRequest } from "~/utils/.server/requests/rating";
import {
  testArtistName,
  testFestivalName,
} from "../../../test/mock-data/rating";
import { userEvent } from "@testing-library/user-event";
import mockServer, { testApi } from "../../../test/mocks";
import { http, HttpResponse } from "msw";

describe("loader", () => {
  test("fetches unrated artists for festival", async () => {
    const unratedArtistsRequestSpy = vi.spyOn(
      festivalRequests,
      "getUnratedFestivalArtists",
    );
    const response = await loader({
      request: new Request("http://app.com"),
      params: { name: testFestivalName },
      context: {},
    });

    expect(unratedArtistsRequestSpy).toHaveBeenCalledTimes(1);
    expect(unratedArtistsRequestSpy).toHaveBeenCalledWith(
      expect.anything(),
      testFestivalName,
    );
    expect(response.data).toEqual({
      ok: true,
      data: testFestivalArtistsData.map((a) => toFestivalArtist(a)),
    });
  });

  test("throws error if festival is not supported", async () => {
    const errorMessage = "Festival not supported";
    mockServer.use(
      http.get(`${testApi}/festivals/${testFestivalName}`, () => {
        return HttpResponse.json({ error: errorMessage }, { status: 404 });
      }),
    );
    let errorData;
    try {
      await loader({
        request: new Request("http://app.com"),
        params: { name: testFestivalName },
        context: {},
      });
    } catch (error) {
      errorData = error as { data: string };
    }
    expect(errorData).toMatchObject({ data: errorMessage });
  });
});

test("shows RatingCards for unrated artists", async () => {
  const RemixStub = createRoutesStub([
    {
      path: "/festivals/:name",
      // @ts-expect-error Type error by react-router (https://github.com/remix-run/react-router/issues/13579)
      Component: FestivalRoute,
      loader: async () => {
        return data({
          ok: true,
          data: testFestivalArtistsData.map((a) => toFestivalArtist(a)),
        });
      },
    },
  ]);
  render(<RemixStub initialEntries={[`/festivals/${testFestivalName}`]} />);

  expect(
    await screen.findByText(testFestivalArtistsData[0].artist_name),
  ).toBeVisible();
  expect(
    screen.getByText(testFestivalArtistsData[1].artist_name),
  ).toBeVisible();
  expect(
    screen.getByText(testFestivalArtistsData[2].artist_name),
  ).toBeVisible();
});

test("rate unrated artist", async () => {
  const newRatingRequest: RatingRequest = {
    artist_name: testArtistName,
    festival_name: testFestivalName,
    year: 2015,
    rating: 5,
    comment: "Old school swedish death metal",
  };
  const user = userEvent.setup();
  const RemixStub = createRoutesStub([
    {
      path: "/festivals/:name",
      // @ts-expect-error Type error by react-router (https://github.com/remix-run/react-router/issues/13579)
      Component: FestivalRoute,
      loader: async () => {
        return data({
          ok: true,
          data: [toFestivalArtist(testFestivalArtistsData[0])],
        });
      },
    },
    {
      path: "/ratings",
      action: async ({ request }) => {
        const formData = await request.formData();
        expect(formData.get("_action")).toEqual("SAVE_RATING");
        expect(formData.get("artist_name")).toEqual(
          newRatingRequest.artist_name,
        );
        expect(formData.get("festival_name")).toEqual(
          newRatingRequest.festival_name,
        );
        expect(formData.get("year")).toEqual(newRatingRequest.year!.toString());
        expect(formData.get("rating")).toEqual(
          newRatingRequest.rating.toString(),
        );
        expect(formData.get("comment")).toEqual(newRatingRequest.comment);
        return data({ ok: true });
      },
    },
  ]);
  render(<RemixStub initialEntries={[`/festivals/${testFestivalName}`]} />);

  await user.type(
    await screen.findByLabelText(/festival/i),
    newRatingRequest.festival_name!,
  );
  await user.type(
    screen.getByLabelText(/year/i),
    newRatingRequest.year!.toString(),
  );
  // this only works with fireEvent and not userEvent
  fireEvent.click(screen.getByLabelText("5 Stars"));
  await user.type(screen.getByLabelText(/comment/i), newRatingRequest.comment!);
  await user.click(screen.getByText(/rate/i));

  expect(screen.queryByText(/AssertionError/i)).not.toBeInTheDocument();
});
