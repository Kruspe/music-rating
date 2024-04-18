import { authenticator } from "~/utils/auth.server";
import { ActionFunctionArgs, redirect } from "@remix-run/node";

export async function action({ request }: ActionFunctionArgs) {
  const { token } = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  const form = await request.formData();
  const artistName = form.get("artist_name") as string;
  const festivalName = form.get("festival_name") as string;
  const rating = form.get("rating") as string;
  const year = form.get("year") as string;
  const comment = form.get("comment") as string;

  await fetch(`${process.env.API_ENDPOINT}/ratings/${artistName}`, {
    method: "PUT",
    body: JSON.stringify({
      festival_name: festivalName,
      rating: parseFloat(rating),
      year: parseInt(year, 10),
      comment: comment,
    }),
    headers: {
      authorization: `Bearer ${token}`,
    },
  });
  return redirect("/ratings");
}
