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
    it('should display unrated bands', async () => {
      global.fetch = jest.fn().mockResolvedValueOnce(new Response(JSON.stringify(
        [
          { artist: 'Bloodbath', image: 'bloodbathImage' },
          { artist: 'Megadeth', image: 'megadethImage' },
          { artist: 'Vader', image: 'vaderImage' }],
      )));
      const {
        getByText, queryByText, getByAltText, queryByAltText,
      } = render(
        <UserContext.Provider value={{ ratedBands: [{ band: 'Bloodbath' }] }}>
          <EstimateWacken />
        </UserContext.Provider>,
      );

      await wait(() => {
        expect(global.fetch).toHaveBeenCalledTimes(1);
        expect(global.fetch).toHaveBeenCalledWith('www.link-to-json.com',
          { headers: { 'Content-Type': 'application/json' } });
      });
      expect(queryByText(/bloodbath/i)).toBeNull();
      expect(getByText('Megadeth')).toBeVisible();
      expect(getByText('Vader')).toBeVisible();
      expect(queryByAltText(/bloodbath/i)).toBeNull();
      expect(getByAltText('Megadeth')).toBeVisible();
      expect(getByAltText('Vader')).toBeVisible();
    });
    it('should remove unrated band after rating', async () => {
      global.fetch = jest.fn().mockResolvedValueOnce(new Response(JSON.stringify(
        [
          { artist: 'Vader', image: 'vaderImage' }],
      )));
      const postSpy = jest.spyOn(API, 'post').mockImplementation((f) => f);
      const expectedInit = {
        header: { Authorization: 'Bearer token' },
        body: {
          user: 'userId', band: 'Vader', festival: 'Wacken', year: '2015', rating: 5, comment: 'comment',
        },
      };
      const {
        findByText, getByLabelText, getByText, queryByText,
      } = render(
        <UserContext.Provider value={{ userId: 'userId', jwtToken: 'token', ratedBands: [] }}>
          <EstimateWacken />
        </UserContext.Provider>,
      );
      expect(await findByText(/vader/i)).toBeVisible();
      fireEvent.click(getByText(/vader/i));
      fireEvent.change(getByLabelText(/festival \*/i), { target: { value: 'Wacken' } });
      fireEvent.change(getByLabelText(/year \*/i), { target: { value: '2015' } });
      fireEvent.click(getByLabelText(/5 star/i));
      fireEvent.change(getByLabelText(/comment/i), { target: { value: 'comment' } });
      await act(async () => {
        fireEvent.submit(getByText(/submit/i));
      });
      await wait(() => expect(postSpy).toHaveBeenCalledTimes(1));
      await wait(() => expect(postSpy).toHaveBeenCalledWith('musicrating', '/api/v1/ratings/bands', expectedInit));
      expect(queryByText(/band/i)).not.toBeInTheDocument();
    });
  });
});
