import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { Auth } from 'aws-amplify';
import App from './App';

describe('App', () => {
  const username = 'username';
  const password = 'password';

  it('should allow users to sign in', async () => {
    render(<App />);
    const usernameField = screen.getByLabelText(/username/i);
    const passwordField = screen.getByLabelText(/password/i);
    await userEvent.type(usernameField, username);
    await userEvent.type(passwordField, password);

    expect(usernameField).toHaveValue(username);
    expect(passwordField).toHaveValue(password);
    userEvent.click(screen.getByText(/sign in/i));

    await waitFor(async () => {
      expect(Auth.signIn).toHaveBeenCalledWith(username, password);
      expect(await screen.findByText(/overviewpage/i)).toBeVisible();
    });
  });
});
