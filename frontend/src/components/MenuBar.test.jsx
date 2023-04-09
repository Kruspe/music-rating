import { useAuth0, withAuthenticationRequired } from '@auth0/auth0-react';
import userEvent from '@testing-library/user-event';
import { renderWithRouteProvider, screen } from '../../test/test-utils';
import MenuBar from './MenuBar';

jest.mock('@auth0/auth0-react');
jest.mock('../home');

let user;
beforeEach(() => {
  user = userEvent.setup();
});

test('should allow to login', async () => {
  const mockLoginWithRedirect = jest.fn();
  useAuth0.mockImplementation(() => ({
    isAuthenticated: false,
    loginWithRedirect: mockLoginWithRedirect,
  }));
  renderWithRouteProvider(<MenuBar />);
  await user.click(screen.getByRole('button', { name: 'Log in' }));

  expect(mockLoginWithRedirect).toHaveBeenCalled();
});

test('should allow to logout', async () => {
  const mockLogout = jest.fn();
  useAuth0.mockImplementation(() => ({
    isAuthenticated: true,
    logout: mockLogout,
  }));
  withAuthenticationRequired.mockImplementation((component) => component);
  renderWithRouteProvider(<MenuBar />);
  await user.click(screen.getByRole('button', { name: 'Log out' }));

  expect(mockLogout).toHaveBeenCalled();
});

test('should have correct hrefs set', async () => {
  useAuth0.mockImplementation(() => ({
    isAuthenticated: true,
  }));
  withAuthenticationRequired.mockImplementation((component) => component);
  renderWithRouteProvider(<MenuBar />);

  expect(await screen.findByText('My Ratings')).toHaveAttribute('href', '/ratings');
  expect(screen.getByText('Wacken')).toHaveAttribute('href', '/wacken');
});
