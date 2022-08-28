import { render, screen } from '@testing-library/react';
import { QueryClientProvider } from 'react-query';
import { Auth0Provider } from '@auth0/auth0-react';
import Login from './login';
import App from './App';

jest.mock('react-query');
jest.mock('@auth0/auth0-react');
jest.mock('./login');

beforeEach(() => {
  QueryClientProvider.mockImplementation(({ children }) => (
    <>
      QueryClientProvider
      {children}
    </>
  ));
  Auth0Provider.mockImplementation(({
    clientId, domain, redirectUri, children,
  }) => (
    <div>
      <p>Auth0Provider</p>
      <p>{`ClientId: ${clientId}`}</p>
      <p>{`Domain: ${domain}`}</p>
      <p>{`RedirectUri: ${redirectUri}`}</p>
      {children}
    </div>
  ));
  Login.mockImplementation(() => (
    <p>LoginComponent</p>
  ));
});

it('should use QueryClientProvider and Auth0Provider', () => {
  render(<App />);
  expect(screen.getByText('QueryClientProvider')).toBeVisible();
  expect(screen.getByText('Auth0Provider')).toBeVisible();
  expect(screen.getByText('ClientId: prjn715M1O1ysyL8yxOF8gjdcWpnq9a4')).toBeVisible();
  expect(screen.getByText('Domain: https://musicrating.eu.auth0.com')).toBeVisible();
  expect(screen.getByText('RedirectUri: http://localhost:3000')).toBeVisible();
  expect(screen.getByText('LoginComponent')).toBeVisible();
});
