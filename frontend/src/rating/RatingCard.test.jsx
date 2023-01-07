import { render, screen } from '../../test/test-utils';
import { unratedArtist } from '../../test/mocks';
import RatingCard from './RatingCard';

it('should show image', () => {
  render(<RatingCard artistName={unratedArtist.artist_name} imageUrl={unratedArtist.image_url} />);

  expect(screen.getByText(unratedArtist.artist_name)).toBeVisible();
  expect(screen.getByAltText(`${unratedArtist.artist_name} image`)).toBeVisible();
});

it('should show artist name when no image is supplied', () => {
  render(<RatingCard artistName={unratedArtist.artist_name} />);

  expect(screen.getByText(unratedArtist.artist_name)).toBeVisible();
});
