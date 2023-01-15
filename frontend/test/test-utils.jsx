/* eslint react/prop-types: 0 */
/* eslint import/no-extraneous-dependencies: 0 */

import { render } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { useAuth0 } from '@auth0/auth0-react';
import { createMemoryRouter, RouterProvider } from 'react-router-dom';
import { mockServer, TestToken, TestUserId } from './mocks';
import routesConfig from '../src/config';

function QueryClientProviderWrapper({ children }) {
  return (
    <QueryClientProvider client={new QueryClient()}>
      {children}
    </QueryClientProvider>
  );
}

const customRender = (ui, options) => render(ui, { wrapper: QueryClientProviderWrapper, ...options });
const renderWithRouteProvider = (ui, options = {}) => render(ui, {
  wrapper: () => (
    <QueryClientProviderWrapper>
      <RouterProvider router={createMemoryRouter(routesConfig, { initialEntries: options.initialEntries })} />
    </QueryClientProviderWrapper>
  ),
  ...options,
});

beforeAll(() => mockServer.listen());
afterEach(() => mockServer.resetHandlers());
afterAll(() => mockServer.close());

jest.mock('@auth0/auth0-react');
beforeEach(() => {
  useAuth0.mockImplementation(() => ({
    getAccessTokenSilently: () => Promise.resolve(TestToken),
    user: {
      sub: TestUserId,
    },
  }));
});

export * from '@testing-library/react';
export { customRender as render, renderWithRouteProvider };
