import { rest } from 'msw';
import userEvent from '@testing-library/user-event';
import { render, screen } from '../../test/test-utils';
import Ratings from './index';
import { bloodbathRating, hypocrisyRating, mockServer } from '../../test/mocks';

test('should show all ratings of user', async () => {
  render(<Ratings />);
  expect(await screen.findByText(bloodbathRating.artist_name)).toBeVisible();
  expect(screen.getByText(hypocrisyRating.artist_name)).toBeVisible();
});

test('should be able to update rating', async () => {
  const user = userEvent.setup();
  render(<Ratings />);
  expect(await screen.findByText(bloodbathRating.artist_name)).toBeVisible();

  await user.dblClick(screen.getByText(bloodbathRating.year));
  await user.keyboard('{Backspace>20/}2023');
  await user.tab();
  expect(await screen.findByText('2023')).toBeVisible();

  await user.dblClick(screen.getAllByText(bloodbathRating.festival_name)[0]);
  await user.keyboard('{Backspace>20/}RUDE');
  await user.tab();
  expect(await screen.findByText('RUDE')).toBeVisible();
});

test('should show message when no bands are rated', async () => {
  mockServer.use(rest.get('http://localhost:8080/ratings', (req, res, ctx) => res(
    ctx.status(200),
    ctx.json([]),
  )));
  render(<Ratings />);

  expect(await screen.findByText('You have not rated any bands so far')).toBeVisible();
});
