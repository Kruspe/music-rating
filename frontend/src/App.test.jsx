import { useAuth0 } from '@auth0/auth0-react';
import { render, screen } from '@testing-library/react';
import App from './App';
import Wacken from './wacken';

jest.mock('@auth0/auth0-react');
jest.mock('./wacken');

beforeEach(() => {
  useAuth0.mockImplementation(() => ({
    isAuthenticated: true,
  }));
  Wacken.mockImplementation(() => (<p>WackenComponent</p>));
});

it('should show log in when user is not logged in', () => {
  useAuth0.mockImplementation(() => ({
    isAuthenticated: false,
  }));
  render(<App />);

  expect(screen.getByText('You need to log in in order to rate your music')).toBeVisible();
  expect(screen.getByRole('button', { name: 'Log in' })).toBeVisible();
});

it('should show bands for wacken', () => {
  render(<App />);

  expect(screen.getByText('WackenComponent')).toBeVisible();
});
