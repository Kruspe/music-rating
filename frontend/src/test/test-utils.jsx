/* eslint react/prop-types: 0 */
/* eslint import/no-extraneous-dependencies: 0 */

import { render } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from 'react-query';

const QueryClientProviderWrapper = ({ children }) => (
  <QueryClientProvider client={new QueryClient()}>
    {children}
  </QueryClientProvider>
);

const customRender = (ui, options) => render(ui, { wrapper: QueryClientProviderWrapper, ...options });

export * from '@testing-library/react';
export { customRender as render };
