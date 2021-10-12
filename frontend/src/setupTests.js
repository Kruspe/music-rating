/* eslint import/no-extraneous-dependencies: 0 */
import '@testing-library/jest-dom';
import fetchMock from './test/mocks';

beforeEach(() => {
  window.fetch = jest.fn().mockImplementation(fetchMock);
});

afterEach(() => {
  jest.restoreAllMocks();
});
