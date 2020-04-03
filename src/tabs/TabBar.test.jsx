import { fireEvent, render } from '@testing-library/react';
import { Storage } from 'aws-amplify';
import React from 'react';
import TabBar from './TabBar';
import UserContext from '../context/UserContext';

describe('TabBar', () => {
  function renderTabBar() {
    return render(
      <UserContext.Provider value={{
        userId: 'userId',
        jwtToken: 'token',
        ratedArtists: [{
          band: 'Bloodbath', festival: 'Wacken', year: 2015, rating: 5, comment: 'comment',
        }],
      }}
      >

        <TabBar authState="signedIn" />
      </UserContext.Provider>,
    );
  }

  it('should render overview', async () => {
    const { findByText } = renderTabBar();
    expect(await findByText(/OverviewPage/i)).toBeVisible();
  });

  it('should switch to wacken tab when clicked', async () => {
    jest.spyOn(Storage, 'get').mockResolvedValueOnce('www.link-to-json.com');
    global.fetch = jest.fn().mockResolvedValueOnce(new Response(JSON.stringify(
      [
        { artist: 'Vader', image: 'vaderImage' }],
    )));
    const { getByText, findByText, findByAltText } = renderTabBar();

    fireEvent.click(getByText(/estimate wacken/i));
    expect(await findByAltText('Vader')).toBeVisible();
    expect(await findByText('Vader')).toBeVisible();
  });
});
