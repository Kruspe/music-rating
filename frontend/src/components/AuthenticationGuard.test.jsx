import { withAuthenticationRequired } from '@auth0/auth0-react';
import { waitFor } from '@testing-library/react';
import { render, screen } from '../../test/test-utils';
import AuthenticationGuard from './AuthenticationGuard';

jest.mock('@auth0/auth0-react');

function TestComponent() {
  return <p>TestComponent</p>;
}

test('should authenticate component', async () => {
  const mockWithAuthenticationRequired = jest.fn();
  withAuthenticationRequired.mockImplementation((component) => {
    mockWithAuthenticationRequired.mockReturnValue(component());
    return mockWithAuthenticationRequired;
  });
  render(<AuthenticationGuard component={TestComponent} />);

  await waitFor(() => expect(mockWithAuthenticationRequired).toHaveBeenCalled());
  expect(screen.getByText('TestComponent')).toBeVisible();
});
