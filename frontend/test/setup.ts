import { vi } from "vitest";
import "@testing-library/jest-dom";

import mockServer, { testApi } from "./mocks";

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
