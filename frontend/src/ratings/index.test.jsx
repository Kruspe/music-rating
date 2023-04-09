import { rest } from 'msw';
import { render, screen } from '../../test/test-utils';
import Ratings from './index';
import { bloodbathRating, hypocrisyRating, mockServer } from '../../test/mocks';

test('should show all ratings of user', async () => {
  render(<Ratings />);
  expect(await screen.findByText(bloodbathRating.artist_name)).toBeVisible();
  expect(screen.getByText(hypocrisyRating.artist_name)).toBeVisible();
});

test('should show message when no bands are rated', async () => {
  mockServer.use(rest.get('http://localhost:8080/api/ratings', (req, res, ctx) => res(
    ctx.status(200),
    ctx.json([]),
  )));
  render(<Ratings />);

  expect(await screen.findByText('You have not rated any bands so far')).toBeVisible();
});
