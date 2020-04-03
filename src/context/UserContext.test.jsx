import React from 'react';
import { render } from '@testing-library/react';
import UserContext from './UserContext';

describe('UserContext', () => {
  it('should have correct default values', async () => {
    const { getByTestId } = render(
      <UserContext.Consumer>
        {(userContext) => (
          <>
            <p data-testid="userId">{userContext.userId}</p>
            <p data-testid="jwtToken">{userContext.jwtToken}</p>
            <p data-testid="ratedArtists">{userContext.ratedArtists}</p>
          </>
        )}
      </UserContext.Consumer>,
    );
    expect(getByTestId('userId')).toHaveTextContent('');
    expect(getByTestId('jwtToken')).toHaveTextContent('');
    expect(getByTestId('ratedArtists')).toHaveTextContent('');
  });
});
