import { API, Auth } from 'aws-amplify';
import { render } from '@testing-library/react';
import React from 'react';
import UserProvider from './UserProvider';
import UserContext from '../context/UserContext';

describe('UserProvider', () => {
  afterEach(() => {
    jest.restoreAllMocks();
  });
  it('should display userId, jwtToken and ratedBands', async () => {
    const ratedBandsGetSpy = jest.spyOn(API, 'get').mockResolvedValue([{
      band: 'Bloodbath', festival: 'Wacken', year: 2015, rating: 5, comment: 'comment',
    }]);
    const currentSessionMock = {
      getAccessToken: () => ({ getJwtToken: () => ('token') }),
    };
    const currentUserInfoMock = { id: 'userId' };
    jest.spyOn(Auth, 'currentUserInfo').mockResolvedValueOnce(currentUserInfoMock);
    jest.spyOn(Auth, 'currentSession').mockResolvedValueOnce(currentSessionMock);

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
    expect(ratedBandsGetSpy).toHaveBeenCalledTimes(1);
    expect(ratedBandsGetSpy).toHaveBeenCalledWith('musicrating', '/api/v1/ratings/bands/userId', { header: { Authorization: 'Bearer token' } });
  });
  it('should not get bands if userId and jwtToken are not available', () => {
    const getSpy = jest.spyOn(API, 'get');
    const currentSessionMock = {
      getAccessToken: () => ({ getJwtToken: () => ('') }),
    };
    const currentUserInfoMock = { id: '' };
    jest.spyOn(Auth, 'currentUserInfo').mockResolvedValueOnce(currentUserInfoMock);
    jest.spyOn(Auth, 'currentSession').mockResolvedValueOnce(currentSessionMock);

    render(<UserProvider />);

    expect(getSpy).not.toHaveBeenCalled();
  });
});
