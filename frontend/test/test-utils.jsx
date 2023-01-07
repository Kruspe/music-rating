/* eslint react/prop-types: 0 */
/* eslint import/no-extraneous-dependencies: 0 */

import { render } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { useAuth0 } from '@auth0/auth0-react';
import { mockServer, TestToken, TestUserId } from './mocks';

function QueryClientProviderWrapper({ children }) {
  return (
    <QueryClientProvider client={new QueryClient()}>
      {children}
    </QueryClientProvider>
  );
}

const customRender = (ui, options) => render(ui, { wrapper: QueryClientProviderWrapper, ...options });

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
export { customRender as render };
