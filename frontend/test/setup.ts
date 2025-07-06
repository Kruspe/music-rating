import { vi } from "vitest";
import "@testing-library/jest-dom";

import mockServer, { testApi } from "./mocks";

vi.mock("~/utils/session.server", () => ({
  sessionStorage: {
    getSession: async () => ({
      get: (name: string) => {
        if (name === "user") {
          return { token: "test-token" };
        }
      },
    }),
  },
}));
vi.mock("~/utils/auth.server", () => ({}));

beforeAll(() => {
  mockServer.listen();
  vi.stubEnv("API_ENDPOINT", testApi);
});
afterEach(() => {
  mockServer.resetHandlers();
});
afterAll(() => {
  mockServer.close();
});
