import { fireEvent, render, screen } from '@testing-library/react';
import React from 'react';
import TabBar from './TabBar';

describe('TabBar', () => {
  it('should render overview', async () => {
    render(<TabBar />);
    expect(await screen.findByText(/OverviewPage/i)).toBeVisible();
  });

  it('should switch to wacken tab when clicked', async () => {
    render(<TabBar />);

    fireEvent.click(screen.getByText(/estimate wacken/i));
    expect(await screen.findByAltText('Vader')).toBeVisible();
    expect(screen.getByAltText('Megadeth')).toBeVisible();
    expect(screen.getByText('Vader')).toBeVisible();
    expect(screen.getByText('Megadeth')).toBeVisible();
  });
});
