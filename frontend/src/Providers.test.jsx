import { render, screen } from '@testing-library/react';
import { QueryClientProvider } from '@tanstack/react-query';
import { Auth0Provider } from '@auth0/auth0-react';
import Providers from './Providers';

jest.mock('@tanstack/react-query');
jest.mock('@auth0/auth0-react');

beforeEach(() => {
  QueryClientProvider.mockImplementation(({ children }) => (
    <>
      QueryClientProvider
      {children}
    </>
  ));
  Auth0Provider.mockImplementation(({
    audience, clientId, domain, redirectUri, children,
  }) => (
    <div>
      <p>Auth0Provider</p>
      <p>{`ClientId: ${clientId}`}</p>
      <p>{`Domain: ${domain}`}</p>
      <p>{`RedirectUri: ${redirectUri}`}</p>
      <p>{`Audience: ${audience}`}</p>
      {children}
    </div>
  ));
});

it('should load all Providers and ContentComponent', () => {
  render(<Providers><p>Content</p></Providers>);
  expect(screen.getByText('QueryClientProvider')).toBeVisible();
  expect(screen.getByText('Auth0Provider')).toBeVisible();
  expect(screen.getByText('ClientId: prjn715M1O1ysyL8yxOF8gjdcWpnq9a4')).toBeVisible();
  expect(screen.getByText('Domain: https://musicrating.eu.auth0.com')).toBeVisible();
  expect(screen.getByText('RedirectUri: http://localhost:3000')).toBeVisible();
  expect(screen.getByText('Audience: http://localhost:3000')).toBeVisible();
  expect(screen.getByText('Content')).toBeVisible();
});
