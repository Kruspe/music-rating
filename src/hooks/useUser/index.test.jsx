import React from 'react';
import { render, screen } from '@testing-library/react';
import { Auth } from 'aws-amplify';
import useUser from './index';

const UseUserHookExample = () => {
  const { userId, token } = useUser();
  return (
    <>
      <p>{userId.data}</p>
      <p>{token.data}</p>
    </>
  );
};

describe('useUser', () => {
  it('should get userId and token', async () => {
    const getJwtTokenMock = jest.fn().mockReturnValueOnce('token');
    Auth.currentSession.mockReturnValueOnce({
      getAccessToken: () => ({ getJwtToken: getJwtTokenMock }),
    });
    render(<UseUserHookExample />);

    expect(await screen.findByText('userId')).toBeInTheDocument();
    expect(await screen.findByText('token')).toBeInTheDocument();
    expect(Auth.currentUserInfo).toHaveBeenCalledTimes(1);
    expect(getJwtTokenMock).toHaveBeenCalledTimes(1);
  });
});
