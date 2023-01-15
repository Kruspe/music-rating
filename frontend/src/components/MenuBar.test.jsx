import { useAuth0 } from '@auth0/auth0-react';
import userEvent from '@testing-library/user-event';
import { render, screen } from '../../test/test-utils';
import MenuBar from './MenuBar';

jest.mock('@auth0/auth0-react');

let user;
beforeEach(() => {
  user = userEvent.setup();
});

it('should allow to login', async () => {
  const mockLoginWithRedirect = jest.fn();
  useAuth0.mockImplementation(() => ({
    isAuthenticated: false,
    loginWithRedirect: mockLoginWithRedirect,
  }));
  render(<MenuBar />);
  await user.click(screen.getByRole('button', { name: 'Log in' }));

  expect(mockLoginWithRedirect).toHaveBeenCalled();
});

it('should allow to logout', async () => {
  const mockLogout = jest.fn();
  useAuth0.mockImplementation(() => ({
    isAuthenticated: true,
    logout: mockLogout,
  }));
  render(<MenuBar />);
  await user.click(screen.getByRole('button', { name: 'Log out' }));

  expect(mockLogout).toHaveBeenCalled();
});
