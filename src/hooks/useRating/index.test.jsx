import React from 'react';
import { render, screen } from '@testing-library/react';
import { API } from 'aws-amplify';
import useUser from '../useUser';
import useRating from './index';

jest.mock('../useUser');

const UseRatingHookExample = () => {
  const { data: ratings } = useRating();
  return (
    <>
      <p>
        Ratings:
        {JSON.stringify(ratings)}
      </p>
    </>
  );
};

describe('useRating', () => {
  it('should return mocked ratings', async () => {
    useUser.mockReturnValue({ userId: { data: 'userId' }, token: { data: 'token' } });
    render(<UseRatingHookExample />);
    expect(await screen.findByText(`Ratings:${JSON.stringify([{
      artist: 'Bloodbath',
      festival: 'Wacken',
      year: 2015,
      rating: 5,
      comment: 'comment',
    }])}`));
    expect(API.get).toHaveBeenCalledWith('musicrating',
      '/api/v1/ratings/bands/userId',
      { header: { Authorization: 'Bearer token' } });
    expect(API.get).toHaveBeenCalledTimes(1);
  });
  it('should return no ratings when userId is undefined', () => {
    useUser.mockReturnValue({ userId: { data: undefined }, token: { data: 'token' } });
    render(<UseRatingHookExample />);
    expect(screen.getByText('Ratings:'));
  });
  it('should return no ratings when token is undefined', () => {
    useUser.mockReturnValue({ userId: { data: 'userId' }, token: { data: undefined } });
    render(<UseRatingHookExample />);
    expect(screen.getByText('Ratings:'));
  });
});
