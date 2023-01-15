import { useAuth0, withAuthenticationRequired } from '@auth0/auth0-react';
import { render, renderWithRouteProvider, screen } from '../../test/test-utils';
import Home from './index';
import Wacken from '../wacken';

jest.mock('@auth0/auth0-react');
jest.mock('../wacken');

test('should show empty page while auth information are still loading', () => {
  useAuth0.mockImplementation(() => ({
    isLoading: true,
  }));

  render(<Home />);

  expect(screen.queryByText('MusicRating')).not.toBeInTheDocument();
});

test('should show information that user has to login', () => {
  useAuth0.mockImplementation(() => ({
    isAuthenticated: false,
  }));
  render(<Home />);

  expect(screen.getByText('MusicRating')).toBeVisible();
  expect(screen.getByText('You need to log in in order to rate your music')).toBeVisible();
});

test('should navigate to Wacken view when user is logged in', () => {
  useAuth0.mockImplementation(() => ({
    isAuthenticated: true,
  }));
  withAuthenticationRequired.mockImplementation((component) => component);
  Wacken.mockImplementation(() => <p>WackenComponent</p>);

  renderWithRouteProvider(<Home />);

  expect(screen.getByText('WackenComponent')).toBeVisible();
});
