import { QueryClientProvider } from '@tanstack/react-query';
import { Auth0Provider } from '@auth0/auth0-react';
import { RouterProvider } from 'react-router-dom';
import { render, screen } from '@testing-library/react';
import App from './App';
import MenuBar from './components/MenuBar';

jest.mock('@tanstack/react-query');
jest.mock('@auth0/auth0-react');
jest.mock('react-router-dom');
jest.mock('./components/MenuBar');

test('should load all Providers and MenuBar', () => {
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
  RouterProvider.mockImplementation(() => <p>RouterProvider</p>);
  MenuBar.mockImplementation(() => <p>MenuBar</p>);

  render(<App />);

  expect(screen.getByText('QueryClientProvider')).toBeVisible();

  expect(screen.getByText('Auth0Provider')).toBeVisible();
  expect(screen.getByText('ClientId: prjn715M1O1ysyL8yxOF8gjdcWpnq9a4')).toBeVisible();
  expect(screen.getByText('Domain: https://musicrating.eu.auth0.com')).toBeVisible();
  expect(screen.getByText('RedirectUri: http://localhost:3000/wacken')).toBeVisible();
  expect(screen.getByText('Audience: undefined')).toBeVisible();

  expect(screen.getByText('RouterProvider')).toBeVisible();

  expect(screen.getByText('MenuBar')).toBeVisible();
});
