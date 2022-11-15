import { useAuth0 } from '@auth0/auth0-react';
import { render, screen } from '@testing-library/react';
import App from './App';

jest.mock('@auth0/auth0-react');

beforeEach(() => {
  useAuth0.mockImplementation(() => ({
    isAuthenticated: true,
  }));
});

it('should show log in when user is not logged in', () => {
  useAuth0.mockImplementation(() => ({
    isAuthenticated: false,
  }));
  render(<App />);

  expect(screen.getByText('You need to log in in order to rate your music')).toBeVisible();
  expect(screen.getByRole('button', { name: 'Log in' })).toBeVisible();
});

it('should show content when user is logged in', () => {
  render(<App />);

  expect(screen.getByText('Hello World')).toBeVisible();
});
