import { getFestivalArtists } from "~/utils/requests/festival.server";

describe("getFestivalArtists", () => {
  test("returns artists for the festival", async () => {
    console.log("test");
    const festivalArtists = await getFestivalArtists(
      new Request("http://app.com"),
    );

    expect(festivalArtists).toHaveLength(3);
  });
});
