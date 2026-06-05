import { render, screen } from "@testing-library/react";
import App from "~/root";
import { createRoutesStub } from "react-router";
import { userEvent } from "@testing-library/user-event";

vi.mock("~/utils/auth.server", () => ({}));

test("should have correct routes in header", async () => {
  const user = userEvent.setup();
  const RemixStub = createRoutesStub([
    {
      path: "/",
      Component: App,
      children: [
        {
          path: "/festivals/:name",
          Component: () => <p>FestivalRoute</p>,
        },
        {
          path: "/ratings",
          Component: () => <p>RatingsRoute</p>,
        },
        {
          path: "/ratings/:festivalName",
          Component: () => <p>FestivalRatingsRoute</p>,
        },
      ],
    },
  ]);
  render(<RemixStub />);

  await user.click(await screen.findByText("Wacken"));
  await user.click(screen.getByText("Rate"));
  expect(await screen.findByText("FestivalRoute")).toBeVisible();

  await user.click(await screen.findByText("Wacken"));
  await user.click(screen.getByText("Overview"));
  expect(await screen.findByText("FestivalRatingsRoute")).toBeVisible();

  await user.click(await screen.findByText("Dong"));
  await user.click(screen.getByText("Rate"));
  expect(await screen.findByText("FestivalRoute")).toBeVisible();

  await user.click(await screen.findByText("Dong"));
  await user.click(screen.getByText("Overview"));
  expect(await screen.findByText("FestivalRatingsRoute")).toBeVisible();

  await user.click(screen.getByText("RUDE"));
  await user.click(screen.getByText("Rate"));
  expect(await screen.findByText("FestivalRoute")).toBeVisible();

  await user.click(await screen.findByText("RUDE"));
  await user.click(screen.getByText("Overview"));
  expect(await screen.findByText("FestivalRatingsRoute")).toBeVisible();

  await user.click(screen.getByText("My Ratings"));
  expect(await screen.findByText("RatingsRoute")).toBeVisible();
});
