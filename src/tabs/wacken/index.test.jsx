import React from 'react';
import {
  fireEvent, render, screen, wait,
} from '@testing-library/react';
import Wacken from './index';

jest.mock('../../components/rating', () => (() => (<p>Rating</p>)));

describe('Wacken', () => {
  it('should display unrated artists and rating', async () => {
    render(<Wacken />);

    await wait(() => {
      expect(window.fetch).toHaveBeenCalledWith('wackenLink',
        { headers: { 'Content-Type': 'application/json' } });
    });
    expect(window.fetch).toHaveBeenCalledTimes(1);

    const megadethTile = await screen.findByText('Megadeth');
    expect(megadethTile).toBeVisible();
    expect(screen.getByText('Vader')).toBeVisible();
    expect(screen.queryByText(/bloodbath/i)).toBeNull();
    expect(screen.getByAltText('Megadeth')).toBeVisible();
    expect(screen.getByAltText('Vader')).toBeVisible();
    expect(screen.queryByAltText(/bloodbath/i)).toBeNull();

    fireEvent.click(megadethTile);
    expect(screen.getByText('Rating')).toBeInTheDocument();
  });
});
