import React from 'react';
import {
  act, fireEvent, render, wait,
} from '@testing-library/react';
import { API, Storage } from 'aws-amplify';
import EstimateWacken from './index';
import UserContext from '../../context/UserContext';

describe('EstimateWacken', () => {
  describe('Wacken', () => {
    beforeEach(() => {
      jest.spyOn(Storage, 'get').mockResolvedValueOnce('www.link-to-json.com');
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
      global.fetch = jest.fn().mockResolvedValueOnce(new Response(JSON.stringify(
        [
          { name: 'Bloodbath', image: 'bloodbathImage' },
          { name: 'Megadeth', image: 'megadethImage' },
          { name: 'Vader', image: 'vaderImage' }],
      )));
      const { getAllByLabelText } = render(
        <UserContext.Provider value={{ ratedBands: [{ band: 'Bloodbath' }] }}>
          <EstimateWacken />
        </UserContext.Provider>,
      );

      await wait(() => {
        expect(global.fetch).toHaveBeenCalledTimes(1);
        expect(global.fetch).toHaveBeenCalledWith('www.link-to-json.com',
          { headers: { 'Content-Type': 'application/json' } });
      });
      expectAllRatingFieldsToBeVisible(getAllByLabelText);
      const bandFields = getAllByLabelText(/band/i);
      expect(bandFields).toHaveLength(2);
      expect(bandFields[0]).toHaveValue('Megadeth');
      expect(bandFields[1]).toHaveValue('Vader');
    });
    it('should remove unrated band after rating', async () => {
      global.fetch = jest.fn().mockResolvedValueOnce(new Response(JSON.stringify(
        [
          { name: 'Bloodbath', image: 'bloodbathImage' },
          { name: 'Vader', image: 'vaderImage' }],
      )));
      const postSpy = jest.spyOn(API, 'post').mockImplementation((f) => f);
      const expectedInit = {
        header: { Authorization: 'Bearer token' },
        body: {
          user: 'userId', band: 'Vader', festival: 'Wacken', year: '2015', rating: 5, comment: 'comment',
        },
      };
      const {
        findByLabelText, getByLabelText, getByText, queryByLabelText,
      } = render(
        <UserContext.Provider value={{ userId: 'userId', jwtToken: 'token', ratedBands: [{ band: 'Bloodbath' }] }}>
          <EstimateWacken />
        </UserContext.Provider>,
      );

      expect(await findByLabelText(/band/i)).toHaveValue('Vader');
      fireEvent.change(getByLabelText(/festival \*/i), { target: { value: 'Wacken' } });
      fireEvent.change(getByLabelText(/year \*/i), { target: { value: '2015' } });
      fireEvent.click(getByLabelText(/5 star/i));
      fireEvent.change(getByLabelText(/comment/i), { target: { value: 'comment' } });
      await act(async () => {
        fireEvent.submit(getByText(/submit/i));
      });
      await wait(() => expect(postSpy).toHaveBeenCalledTimes(1));
      await wait(() => expect(postSpy).toHaveBeenCalledWith('musicrating', '/bands', expectedInit));
      expect(queryByLabelText(/band/i)).not.toBeInTheDocument();
    });
  });
});
