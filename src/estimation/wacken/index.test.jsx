import React from 'react';
import { act, fireEvent, render, wait, } from '@testing-library/react';
import { API, Auth, Storage } from 'aws-amplify';
import EstimateWacken from './index';


describe('EstimateWacken', () => {
  describe('Wacken', () => {
    beforeEach(() => {
      jest.spyOn(Storage, 'get').mockResolvedValueOnce('www.link-to-json.com');
      jest.spyOn(API, 'get').mockResolvedValueOnce([{
        band: 'Bloodbath', festival: 'Wacken', year: 2015, rating: 5, comment: 'comment',
      }]);
      API.post = jest.fn();
      const currentSessionMock = {
        getAccessToken: () => ({ getJwtToken: () => ('Token') }),
      };
      const currentUserInfoMock = { id: 'userId' };
      jest.spyOn(Auth, 'currentUserInfo').mockResolvedValue(currentUserInfoMock);
      jest.spyOn(Auth, 'currentSession').mockResolvedValue(currentSessionMock);
    });
    const expectAllRatingFieldsToBeVisible = (getAllByLabelText) => {
      getAllByLabelText(/festival \*/i).forEach((festivalField) => {
        expect(festivalField).toBeVisible();
      });
      getAllByLabelText(/year \*/i).forEach((yearField) => {
        expect(yearField).toBeVisible();
      });
      getAllByLabelText(/1 star/i).forEach((ratingField) => {
        expect(ratingField).toBeVisible();
      });
      getAllByLabelText(/comment/i).forEach((commentField) => {
        expect(commentField).toBeVisible();
      });
    };

    it('should display unrated bands', async () => {
      global.fetch = jest.fn().mockResolvedValueOnce(new Response(JSON.stringify(['Bloodbath', 'Megadeth', 'Vader'])));
      const { getAllByLabelText } = render(<EstimateWacken />);

      await wait(() => {
        expect(global.fetch).toHaveBeenCalledTimes(1);
        expect(global.fetch).toHaveBeenCalledWith('www.link-to-json.com',
          { headers: { 'Content-Type': 'application/json' } });
      });
      expectAllRatingFieldsToBeVisible(getAllByLabelText);
      const bandField = getAllByLabelText(/band/i);
      expect(bandField).toHaveLength(2);
      expect(bandField[0]).toHaveValue('Megadeth');
      expect(bandField[1]).toHaveValue('Vader');
    });
    it('should remove unrated band after rating', async () => {
      global.fetch = jest.fn().mockResolvedValueOnce(new Response(JSON.stringify(['Bloodbath', 'Vader'])));
      const {
        findByLabelText, getByLabelText, getByText, queryByLabelText,
      } = render(<EstimateWacken />);

      expect(await findByLabelText(/band/i)).toHaveValue('Vader');
      fireEvent.change(getByLabelText(/festival \*/i), { target: { value: 'Wacken' } });
      fireEvent.change(getByLabelText(/year \*/i), { target: { value: '2015' } });
      fireEvent.click(getByLabelText(/5 star/i));
      fireEvent.change(getByLabelText(/comment/i), { target: { value: 'comment' } });
      await act(async () => {
        fireEvent.submit(getByText(/submit/i));
      });
      expect(queryByLabelText(/band/i)).not.toBeInTheDocument();
    });
  });
});
