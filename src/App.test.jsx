import React from 'react';
import { render } from '@testing-library/react';
import { API, Auth } from 'aws-amplify';
import App from './App';

jest.mock('./tabs/TabBar', () => (() => (<p>TabBar</p>)));

describe('App', () => {
  it('should render no content when not signedIn', () => {
    const { container } = render(<App authState="signIn" />);
    expect(container).toBeEmpty();
  });

  it('should render TabBar when signedIn', async () => {
    const getRatedArtistsSpy = jest.spyOn(API, 'get').mockResolvedValue([{
      band: 'Bloodbath', festival: 'Wacken', year: 2015, rating: 5, comment: 'comment',
    }]);
    const currentSessionMock = {
      getAccessToken: () => ({ getJwtToken: () => ('token') }),
    };
    const currentUserInfoMock = { id: 'userId' };
    jest.spyOn(Auth, 'currentUserInfo').mockResolvedValue(currentUserInfoMock);
    jest.spyOn(Auth, 'currentSession').mockResolvedValue(currentSessionMock);

    const { findByText } = render(<App authState="signedIn" />);
    expect(await findByText('TabBar')).toBeVisible();
    expect(getRatedArtistsSpy).toHaveBeenCalledTimes(1);
  });
});
