import { rest } from 'msw';
import userEvent from '@testing-library/user-event';
import { useQuery } from '@tanstack/react-query';
import { render, screen, waitFor } from '../../test/test-utils';
import Wacken from './index';
import { mockServer, TestToken, unratedArtist } from '../../test/mocks';

let user;
beforeEach(() => {
  user = userEvent.setup();
});

function AllRatingsTestComponent() {
  const { data: ratings, isFetched } = useQuery(['ratings'], async () => {
    const response = await fetch(`${process.env.REACT_APP_API_ENDPOINT}/ratings`, {
      headers: {
        authorization: `Bearer ${TestToken}`,
      },
    });
    return response.json();
  });

  return isFetched && <p>{`Found ${ratings.length} ratings`}</p>;
}

test('should show unrated band and allow to rate it', async () => {
  render(<Wacken />);

  expect(await screen.findByAltText(`${unratedArtist.artist_name} image`)).toBeVisible();

  await user.type(screen.getByLabelText(/Festival\/Concert/), 'Wacken');
  await user.type(screen.getByLabelText(/Year/), '2015');
  await user.click(screen.getByLabelText(/4.5 Stars/));
  await user.type(screen.getByLabelText(/Comment/), 'Swedish death metal');
  await user.click(screen.getByRole('button', { name: 'Rate' }));

  await waitFor(() => expect(screen.queryByAltText(`${unratedArtist.artist_name} image`)).not.toBeInTheDocument());
  render(<AllRatingsTestComponent />);
  expect(await screen.findByText('Found 3 ratings')).toBeVisible();
});

test('should display message that all bands have been rated', async () => {
  mockServer.use(rest.get('http://localhost:8080/festivals/wacken', (req, res, ctx) => res(
    ctx.status(200),
    ctx.json([]),
  )));
  render(<Wacken />);

  expect(await screen.findByText('You rated all bands that are announced at the moment')).toBeVisible();
});
