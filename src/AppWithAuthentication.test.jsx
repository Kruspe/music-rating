import React from 'react';
import { render, waitForElement, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import AppWithAuthentication from './AppWithAuthentication';

describe('Authentication', () => {
  it('should display SignIn form', async () => {
    const { findByText } = render(<AppWithAuthentication />);
    expect(await waitForElement(() => findByText(/sign in$/i))).toBeInTheDocument();
  });

  it('should display correct SignUp form', async () => {
    const { findByText } = render(<AppWithAuthentication />);
    const signUpLink = await waitForElement(() => findByText('Create account'));
    fireEvent.click(signUpLink);
    expect(await waitForElement(() => findByText(/create account/i))).toBeInTheDocument();
  });
});
