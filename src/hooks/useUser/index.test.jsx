import React from 'react';
import { render, screen } from '@testing-library/react';
import { API, Auth } from 'aws-amplify';
import useUser from './index';

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
  it('should get userId, token and ratedArtists', async () => {
    const getJwtTokenMock = jest.fn().mockReturnValueOnce('token');
    Auth.currentSession.mockReturnValueOnce({
      getAccessToken: () => ({ getJwtToken: getJwtTokenMock }),
    });
    render(<UseUserHookExample />);

    expect(await screen.findByText('userId')).toBeInTheDocument();
    expect(await screen.findByText('token')).toBeInTheDocument();
    expect(await screen.findByText('Bloodbath')).toBeInTheDocument();
    expect(Auth.currentUserInfo).toHaveBeenCalledTimes(1);
    expect(getJwtTokenMock).toHaveBeenCalledTimes(1);
    expect(API.get).toHaveBeenCalledWith('musicrating', '/api/v1/ratings/bands/userId', { header: { Authorization: 'Bearer token' } });
    expect(API.get).toHaveBeenCalledTimes(1);
  });

  it('should not get ratedArtists when userId was not fetched', async () => {
    Auth.currentUserInfo.mockResolvedValueOnce({ id: '' });
    render(<UseUserHookExample />);

    expect(await screen.findByText('token')).toBeInTheDocument();
    expect(screen.queryByText('userId')).toBeNull();
    expect(API.get).not.toHaveBeenCalled();
  });

  it('should not get ratedArtists when token was not fetched', async () => {
    const getJwtTokenMock = jest.fn().mockReturnValueOnce('');
    Auth.currentSession.mockReturnValueOnce({
      getAccessToken: () => ({ getJwtToken: getJwtTokenMock }),
    });
    render(<UseUserHookExample />);

    expect(await screen.findByText('userId')).toBeInTheDocument();
    expect(screen.queryByText('token')).toBeNull();
    expect(API.get).not.toHaveBeenCalled();
  });
});
