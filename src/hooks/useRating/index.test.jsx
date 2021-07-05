import { render, screen } from '../../test/test-utils';
import useRating from './index';

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
    render(<UseRatingHookExample />);
    expect(await screen.findByText(`Ratings:${JSON.stringify([{
      artist: 'Bloodbath',
      festival: 'Wacken',
      year: 2015,
      rating: 5,
      comment: 'comment',
    }])}`)).toBeVisible();
  });
});
