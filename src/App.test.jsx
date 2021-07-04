import { render, screen } from '@testing-library/react';
import { QueryClientProvider } from 'react-query';
import App from './App';

jest.mock('react-query');

it('should wrap everything with QueryClientProvider', () => {
  QueryClientProvider.mockImplementation(({ children }) => (
    <>
      QueryClientProvider
      {children}
    </>
  ));

  render(<App />);
  expect(screen.getByText('QueryClientProvider')).toBeVisible();
  expect(screen.getByText(/hello world/i)).toBeVisible();
});
