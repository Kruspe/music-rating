import { isRouteErrorResponse, useRouteError } from "@remix-run/react";

export function ErrorBoundary() {
  const error = useRouteError();

  // when true, this is what used to go to `CatchBoundary`
  if (isRouteErrorResponse(error)) {
    return (
      <div>
        <h1>Oops</h1>
        <p>Status: {error.status}</p>
        <p>{error.data.message}</p>
        <p>{error.statusText}</p>
      </div>
    );
  }

  // // Don't forget to typecheck with your own logic.
  // // Any value can be thrown, not just errors!
  // let errorMessage = "Unknown error";
  // if (error) {
  //   errorMessage = error.message ? error.message ;
  // }

  return (
    <div>
      <h1>Uh oh ...</h1>
      <p>Something went wrong.</p>
    </div>
  );
}
