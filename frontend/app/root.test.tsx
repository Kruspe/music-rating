import { render, screen } from "@testing-library/react";
import App from "~/root";
import { createRemixStub } from "@remix-run/testing";
import { userEvent } from "@testing-library/user-event";

test("should have correct routes in header", async () => {
  const user = userEvent.setup();
  const RemixStub = createRemixStub([
    {
      path: "/",
      Component: App,
      children: [
        {
          path: "/festivals/:name",
          Component: () => <p>WackenRoute</p>,
        },
        {
          path: "/ratings",
          Component: () => <p>RatingsRoute</p>,
        },
      ],
    },
  ]);
  render(<RemixStub />);

  await user.click(await screen.findByText("Wacken"));
  expect(await screen.findByText("WackenRoute")).toBeVisible();

  await user.click(screen.getByText("My Ratings"));
  expect(await screen.findByText("RatingsRoute")).toBeVisible();
});
