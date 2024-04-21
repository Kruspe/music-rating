import { vi } from "vitest";
import "@testing-library/jest-dom";

import mockServer, { testApi } from "./mocks";

vi.mock("~/utils/auth.server", () => ({
  authenticator: {
    isAuthenticated: () => true,
  },
}));

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
