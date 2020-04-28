import React from 'react';
import { fireEvent, render, screen } from '@testing-library/react';
import AppWithAuthentication from './AppWithAuthentication';

describe('Authentication', () => {
  it('should display SignIn form', async () => {
    render(<AppWithAuthentication />);
    expect(await screen.findByText(/sign in to your account/i)).toBeVisible();
    expect(screen.getByPlaceholderText(/enter your username/i)).toBeVisible();
    expect(screen.getByPlaceholderText(/enter your password/i)).toBeVisible();
    expect(screen.getByText(/reset password/i)).toBeVisible();
    expect(screen.getByText(/create account/i)).toBeVisible();
    expect(screen.getByText('Sign In')).toBeVisible();
  });

  it('should display correct SignUp form', async () => {
    render(<AppWithAuthentication />);
    fireEvent.click(await screen.findByText(/create account/i));

    expect(await screen.findByText(/create a new account/i)).toBeVisible();
    expect(screen.getByPlaceholderText(/username/i)).toBeVisible();
    expect(screen.getByPlaceholderText(/password/i)).toBeVisible();
    expect(screen.getByPlaceholderText(/email/i)).toBeVisible();
    expect(screen.queryByPlaceholderText(/phone number/i)).toBeNull();
    expect(screen.getByText(/sign in/i)).toBeVisible();
    expect(screen.getByText('Create Account')).toBeVisible();
  });
});
