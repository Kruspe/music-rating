import { API, Auth } from 'aws-amplify';
import { render } from '@testing-library/react';
import React from 'react';
import UserProvider from './UserProvider';
import UserContext from '../context/UserContext';

describe('UserProvider', () => {
  it('should display userId, jwtToken and ratedBands', async () => {
    jest.spyOn(API, 'get').mockResolvedValue([{
      band: 'Bloodbath', festival: 'Wacken', year: 2015, rating: 5, comment: 'comment',
    }]);
    const currentSessionMock = {
      getAccessToken: () => ({ getJwtToken: () => ('token') }),
    };
    const currentUserInfoMock = { id: 'userId' };
    jest.spyOn(Auth, 'currentUserInfo').mockResolvedValue(currentUserInfoMock);
    jest.spyOn(Auth, 'currentSession').mockResolvedValue(currentSessionMock);

    const { findByText } = render(
      <UserProvider>
        <UserContext.Consumer>
          {(userContext) => (
            <>
              <p>{userContext.userId}</p>
              <p>{userContext.jwtToken}</p>
              {userContext.ratedBands.map((ratedBand) => (
                <p key={ratedBand.band}>{ratedBand.band}</p>
              ))}
            </>
          )}
        </UserContext.Consumer>
      </UserProvider>,
    );

    expect(await findByText('userId')).toBeVisible();
    expect(await findByText('token')).toBeVisible();
    expect(await findByText('Bloodbath')).toBeVisible();
  });
});
