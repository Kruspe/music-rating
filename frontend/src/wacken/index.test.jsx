import { rest } from 'msw';
import userEvent from '@testing-library/user-event';
import { waitFor } from '@testing-library/react';
import { render, screen } from '../../test/test-utils';
import Wacken from './index';
import { mockServer, unratedArtist } from '../../test/mocks';

let user;
beforeEach(() => {
  user = userEvent.setup();
});

it('should show unrated band and allow to rate it', async () => {
  render(<Wacken />);

  expect(await screen.findByAltText(`${unratedArtist.artist_name} image`)).toBeVisible();

  await user.type(screen.getByLabelText(/Festival\/Concert/), 'Wacken');
  await user.type(screen.getByLabelText(/Year/), '2015');
  await user.click(screen.getByLabelText(/5 Stars/));
  await user.type(screen.getByLabelText(/Comment/), 'Swedish death metal');
  await user.click(screen.getByText('Rate'));

  await waitFor(() => expect(screen.queryByAltText(`${unratedArtist.artist_name} image`)).toBeNull());
});

it('should display message that all bands have been rated', async () => {
  mockServer.use(rest.get('http://localhost:8080/api/festivals/wacken', (req, res, ctx) => res(
    ctx.status(200),
    ctx.json([]),
  )));
  render(<Wacken />);

  expect(await screen.findByText('You rated all bands that are announced at the moment')).toBeVisible();
});
