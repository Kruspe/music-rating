import React from 'react';
import { fireEvent, render, screen } from '@testing-library/react';
import EstimationCard from './index';

jest.mock('../rating', () => () => (<p>Rating</p>));

describe('EstimationCard', () => {
  it('should show artist name and display rating when clicked', () => {
    render(<EstimationCard artist="Bloodbath" />);
    fireEvent.click(screen.getByText('Bloodbath'));
    expect(screen.getByText('Rating')).toBeVisible();
  });
  it('should show artist name and image', () => {
    render(<EstimationCard artist="Bloodbath" image="BloodbathImage" />);
    fireEvent.click(screen.getByText('Bloodbath'));
    expect(screen.getByRole('img')).toHaveAttribute('src', 'BloodbathImage');
    expect(screen.getByText('Rating')).toBeVisible();
  });
});
