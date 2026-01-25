// @ts-expect-error this is where the compiled server lives
import * as build from "./build/server";
import type {
  APIGatewayProxyEventHeaders,
  APIGatewayProxyEventV2,
  APIGatewayProxyHandlerV2,
  APIGatewayProxyResultV2,
} from "aws-lambda";
import type { AppLoadContext, ServerBuild } from "react-router";
import { createRequestHandler as createRemixRequestHandler } from "react-router";
import { readableStreamToString } from "@react-router/node";

export const handler = createRequestHandler({
  build,
});

export type GetLoadContextFunction = (
  event: APIGatewayProxyEventV2,
) => Promise<AppLoadContext> | AppLoadContext;

function createRequestHandler({
  build,
  getLoadContext,
  mode = process.env.NODE_ENV,
}: {
  build: ServerBuild;
  getLoadContext?: GetLoadContextFunction;
  mode?: string;
}): APIGatewayProxyHandlerV2 {
  const handleRequest = createRemixRequestHandler(
    {
      ...build,
      allowedActionOrigins: process.env.DOMAIN_NAME
        ? [process.env.DOMAIN_NAME]
        : false,
    },
    mode,
  );

  return async (event) => {
    const request = createRemixRequest(event);
    const loadContext = await getLoadContext?.(event);

    const response = await handleRequest(request, loadContext);

    return sendRemixResponse(response);
  };
}

function createRemixRequest(event: APIGatewayProxyEventV2): Request {
  const host = event.headers["x-forwarded-host"] || event.headers.host;
  const search = event.rawQueryString.length ? `?${event.rawQueryString}` : "";
  const scheme = event.headers["x-forwarded-proto"] || "http";

  const url = new URL(event.rawPath + search, `${scheme}://${host}`);
  const isFormData = event.headers["content-type"]?.includes(
    "multipart/form-data",
  );

  console.error(url.href, event.headers);
  return new Request(url.href, {
    method: event.requestContext.http.method,
    headers: createRemixHeaders(event.headers, event.cookies),
    body:
      event.body && event.isBase64Encoded
        ? isFormData
          ? Buffer.from(event.body, "base64")
          : Buffer.from(event.body, "base64").toString()
        : event.body,
  });
}

function createRemixHeaders(
  requestHeaders: APIGatewayProxyEventHeaders,
  requestCookies?: string[],
): Headers {
  const headers = new Headers();

  for (const [header, value] of Object.entries(requestHeaders)) {
    if (value) {
      headers.append(header, value);
    }
  }

  if (requestCookies) {
    headers.append("Cookie", requestCookies.join("; "));
  }

  return headers;
}

async function sendRemixResponse(
  response: Response,
): Promise<APIGatewayProxyResultV2> {
  const cookies: string[] = [];

  for (const [key, values] of Object.entries(response.headers)) {
    if (key.toLowerCase() === "set-cookie") {
      for (const value of values) {
        cookies.push(value);
      }
    }
  }

  if (cookies.length) {
    response.headers.delete("Set-Cookie");
  }

  const contentType = response.headers.get("Content-Type");
  const isBase64Encoded = isBinaryType(contentType);
  let body: string | undefined;

  if (response.body) {
    if (isBase64Encoded) {
      body = await readableStreamToString(response.body, "base64");
    } else {
      body = await response.text();
    }
  }

  return {
    statusCode: response.status,
    headers: Object.fromEntries(response.headers.entries()),
    cookies,
    body,
    isBase64Encoded,
  };
}

/**
 * Common binary MIME types
 * @see https://github.com/architect/functions/blob/45254fc1936a1794c185aac07e9889b241a2e5c6/src/http/helpers/binary-types.js
 */
const binaryTypes = [
  "application/octet-stream",
  // Docs
  "application/epub+zip",
  "application/msword",
  "application/pdf",
  "application/rtf",
  "application/vnd.amazon.ebook",
  "application/vnd.ms-excel",
  "application/vnd.ms-powerpoint",
  "application/vnd.openxmlformats-officedocument.presentationml.presentation",
  "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
  "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
  // Fonts
  "font/otf",
  "font/woff",
  "font/woff2",
  // Images
  "image/avif",
  "image/bmp",
  "image/gif",
  "image/jpeg",
  "image/png",
  "image/tiff",
  "image/vnd.microsoft.icon",
  "image/webp",
  // Audio
  "audio/3gpp",
  "audio/aac",
  "audio/basic",
  "audio/mpeg",
  "audio/ogg",
  "audio/wav",
  "audio/webm",
  "audio/x-aiff",
  "audio/x-midi",
  "audio/x-wav",
  // Video
  "video/3gpp",
  "video/mp2t",
  "video/mpeg",
  "video/ogg",
  "video/quicktime",
  "video/webm",
  "video/x-msvideo",
  // Archives
  "application/java-archive",
  "application/vnd.apple.installer+xml",
  "application/x-7z-compressed",
  "application/x-apple-diskimage",
  "application/x-bzip",
  "application/x-bzip2",
  "application/x-gzip",
  "application/x-java-archive",
  "application/x-rar-compressed",
  "application/x-tar",
  "application/x-zip",
  "application/zip",
];

function isBinaryType(contentType: string | null | undefined) {
  if (!contentType) return false;
  const [test] = contentType.split(";");
  return binaryTypes.includes(test);
}
