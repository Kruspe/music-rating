/* eslint react/prop-types: 0 */
/* eslint import/no-extraneous-dependencies: 0 */

import { render } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { mockServer } from './mocks';

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

export * from '@testing-library/react';
export { customRender as render };
