import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { API, Auth, Storage } from 'aws-amplify';
import App from './App';

describe('App', () => {
  it('should render no content when not signedIn', () => {
    const { container } = render(<App authState="signIn" />);
    expect(container).toBeEmpty();
  });

  describe('signedIn', () => {
    beforeEach(() => {
      jest.spyOn(API, 'get').mockResolvedValue([{
        band: 'Bloodbath', festival: 'Wacken', year: 2015, rating: 5, comment: 'comment',
      }]);
      const currentSessionMock = {
        getAccessToken: () => ({ getJwtToken: () => ('token') }),
      };
      const currentUserInfoMock = { id: 'userId' };
      jest.spyOn(Auth, 'currentUserInfo').mockResolvedValue(currentUserInfoMock);
      jest.spyOn(Auth, 'currentSession').mockResolvedValue(currentSessionMock);
    });

    const isEstimateWackenVisible = async (findByText, findByAltText) => {
      expect(await findByAltText('Vader')).toBeVisible();
      expect(await findByText('Vader')).toBeVisible();
    };

    it('should render overview', async () => {
      const { findByText } = render(<App authState="signedIn" />);
      expect(await findByText(/OverviewPage/i)).toBeVisible();
    });

    it('should switch to wacken tab when clicked', async () => {
      jest.spyOn(Storage, 'get').mockResolvedValueOnce('www.link-to-json.com');
      global.fetch = jest.fn().mockResolvedValueOnce(new Response(JSON.stringify(
        [
          { artist: 'Vader', image: 'vaderImage' }],
      )));
      const { getByText, findByText, findByAltText } = render(<App authState="signedIn" />);

      fireEvent.click(getByText(/estimate wacken/i));
      await isEstimateWackenVisible(findByText, findByAltText);
    });
  });
});
