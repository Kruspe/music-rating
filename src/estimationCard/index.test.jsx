import React from 'react';
import { render } from '@testing-library/react';
import EstimationCard from './index';

describe('EstimationCard', () => {
  describe('without image', () => {
    it('should show artist name', () => {
      const { getByText } = render(<EstimationCard artist="Bloodbath" />);
      const artistName = getByText('Bloodbath');
      expect(artistName).toBeVisible();
    });
    it('should show artist name and image', () => {
      const { getByText, getByRole } = render(<EstimationCard artist="Bloodbath" image="BloodbathImage" />);
      expect(getByText('Bloodbath')).toBeVisible();
      expect(getByRole('img')).toHaveAttribute('src', 'BloodbathImage');
    });
  });
});
