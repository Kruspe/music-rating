import React from 'react';
import { render, screen } from '@testing-library/react';
import { API, Auth } from 'aws-amplify';
import useUser from './index';

const mockGetJwtToken = jest.fn();
jest.mock('aws-amplify', () => ({
  Auth:
    {
      currentSession: async () => ({
        getAccessToken: () => (
          { getJwtToken: mockGetJwtToken }),
      }),
      currentUserInfo: jest.fn(),
    },
  API: {
    get: jest.fn(),
  },
}));


const UseUserHookExample = () => {
  const { userId, token, ratedArtists } = useUser();
  return (
    <>
      <p>{userId}</p>
      <p>{token}</p>
      <p>{ratedArtists && ratedArtists[0].band}</p>
    </>
  );
};

describe('useUser', () => {
  beforeEach(() => {
    mockGetJwtToken.mockReset();
    Auth.currentUserInfo.mockReset();
    API.get.mockReset();
  });

  it('should get userId, token and ratedArtists', async () => {
    API.get.mockResolvedValueOnce([{
      band: 'Bloodbath', festival: 'Wacken', year: 2015, rating: 5, comment: 'comment',
    }]);
    Auth.currentUserInfo.mockResolvedValueOnce({ id: 'userId' });
    mockGetJwtToken.mockReturnValueOnce('token');
    render(<UseUserHookExample />);

    expect(await screen.findByText('userId')).toBeInTheDocument();
    expect(await screen.findByText('token')).toBeInTheDocument();
    expect(await screen.findByText('Bloodbath')).toBeInTheDocument();
    expect(Auth.currentUserInfo).toHaveBeenCalledTimes(1);
    expect(mockGetJwtToken).toHaveBeenCalledTimes(1);
    expect(API.get).toHaveBeenCalledWith('musicrating', '/api/v1/ratings/bands/userId', { header: { Authorization: 'Bearer token' } });
    expect(API.get).toHaveBeenCalledTimes(1);
  });

  it('should not get ratedArtists when userId was not fetched', async () => {
    Auth.currentUserInfo.mockResolvedValueOnce({ id: '' });
    mockGetJwtToken.mockReturnValueOnce('token');
    render(<UseUserHookExample />);

    expect(await screen.findByText('token')).toBeInTheDocument();
    expect(API.get).not.toHaveBeenCalled();
  });

  it('should not get ratedArtists when token was not fetched', async () => {
    Auth.currentUserInfo.mockResolvedValueOnce({ id: 'userId' });
    mockGetJwtToken.mockReturnValueOnce('');
    render(<UseUserHookExample />);

    expect(await screen.findByText('userId')).toBeInTheDocument();
    expect(API.get).not.toHaveBeenCalled();
  });
});
