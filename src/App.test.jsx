import { render, screen } from '@testing-library/react';
import { QueryClientProvider } from 'react-query';
import { AmplifyAuthenticator, AmplifySignUp } from '@aws-amplify/ui-react';
import { onAuthUIStateChange } from '@aws-amplify/ui-components';
import App from './App';

jest.mock('react-query');
jest.mock('@aws-amplify/ui-react');
jest.mock('@aws-amplify/ui-components');

beforeEach(() => {
  QueryClientProvider.mockImplementation(({ children }) => (
    <>
      QueryClientProvider
      {children}
    </>
  ));

  AmplifyAuthenticator.render.mockImplementation(({ children }) => (
    <>
      AmplifyAuthenticator
      <div>
        {children}
      </div>
    </>
  ));
  AmplifySignUp.render.mockImplementation(({ formFields }) => {
    const enabledFormFields = formFields.map((formField) => formField.type).join(',');
    return (
      <p>
        AmplifySignUp:
        {enabledFormFields}
      </p>
    );
  });
});

it('should only see login and customized sign up when unauthenticated', () => {
  render(<App />);
  expect(screen.getByText('AmplifyAuthenticator')).toBeVisible();
  expect(screen.getByText(/AmplifySignUp:username,password,email/)).toBeVisible();
  expect(screen.queryByText(/hello world/i)).toBeNull();
});

it('should use QueryClientProvider and show app when authenticated', () => {
  onAuthUIStateChange.mockImplementation((authStateHandler) => authStateHandler('signedin'));
  render(<App />);
  expect(screen.getByText('QueryClientProvider')).toBeVisible();
  expect(screen.getByText(/hello world/i)).toBeVisible();
  expect(screen.queryByText('AmplifyAuthenticator')).toBeNull();
});
