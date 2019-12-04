import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { API, Auth } from 'aws-amplify';
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
        getAccessToken: () => ({ getJwtToken: () => ('Token') }),
      };
      const currentUserInfoMock = { id: 'userId' };
      jest.spyOn(Auth, 'currentUserInfo').mockResolvedValue(currentUserInfoMock);
      jest.spyOn(Auth, 'currentSession').mockResolvedValue(currentSessionMock);
    });

    const isOverviewVisible = async (findByPlaceholderText, findByText, findByLabelText) => {
      expect(await findByPlaceholderText(/search/i)).toBeVisible();
      expect(await findByText(/bloodbath/i)).toBeVisible();
      expect(await findByText(/wacken/i)).toBeVisible();
      expect(await findByText(/2015/)).toBeVisible();
      expect(await findByLabelText(/5 stars/i)).toBeVisible();
      expect(await findByText(/10 rows/i)).toBeVisible();
    };

    const isRatingVisible = (getByLabelText) => {
      expect(getByLabelText(/band \*/i)).toBeVisible();
      expect(getByLabelText(/festival \*/i)).toBeVisible();
      expect(getByLabelText(/year \*/i)).toBeVisible();
      expect(getByLabelText(/comment/i)).toBeVisible();
    };

    it('should render overview', async () => {
      const { findByPlaceholderText, findByText, findByLabelText } = render(<App authState="signedIn" />);
      await isOverviewVisible(findByPlaceholderText, findByText, findByLabelText);
    });

    it('should switch between overview and rating', async () => {
      const {
        findByText, getByText, getByLabelText, findByPlaceholderText, findByLabelText,
      } = render(<App authState="signedIn" />);

      fireEvent.click(getByText(/rating/i));
      isRatingVisible(getByLabelText);
      fireEvent.click(getByText(/overview/i));
      await isOverviewVisible(findByPlaceholderText, findByText, findByLabelText);
    });
  });
});
