import { render, screen } from '@testing-library/react';
import { QueryClientProvider } from 'react-query';
import App from './App';

jest.mock('react-query');

beforeEach(() => {
  QueryClientProvider.mockImplementation(({ children }) => (
    <>
      QueryClientProvider
      {children}
    </>
  ));
});

it('should use QueryClientProvider and show app when authenticated', () => {
  render(<App />);
  expect(screen.getByText('QueryClientProvider')).toBeVisible();
  expect(screen.getByText(/hello world/i)).toBeVisible();
  expect(screen.queryByText('AmplifyAuthenticator')).toBeNull();
});
