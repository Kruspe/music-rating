import { vi } from "vitest";
import "@testing-library/jest-dom";

import mockServer, { testApi } from "./mocks";

beforeAll(() => {
  mockServer.listen();
  vi.stubEnv("API_ENDPOINT", testApi);
  vi.mock("../app/utils/auth.server", () => ({
    authenticator: {
      isAuthenticated: () => ({
        token: "test-token",
      }),
    },
  }));
});
afterEach(() => {
  mockServer.resetHandlers();
});
afterAll(() => {
  mockServer.close();
});
