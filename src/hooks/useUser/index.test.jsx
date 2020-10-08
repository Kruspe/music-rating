import React from 'react';
import { render, screen } from '@testing-library/react';
import { Auth } from 'aws-amplify';
import useUser from './index';

const UseUserHookExample = () => {
  const { userId } = useUser();
  return (
    <>
      <p>{userId.data}</p>
    </>
  );
};

describe('useUser', () => {
  it('should get userId', async () => {
    render(<UseUserHookExample />);

    expect(await screen.findByText('userId')).toBeInTheDocument();
    expect(Auth.currentUserInfo).toHaveBeenCalledTimes(1);
  });
});
