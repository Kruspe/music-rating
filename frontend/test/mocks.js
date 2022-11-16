import { rest } from 'msw';
import { setupServer } from 'msw/node';

export const bloodbathRating = {
  artist_name: 'Bloodbath',
  comment: 'old school swedish death metal',
  festival_name: 'Wacken',
  rating: 5,
  year: 2015,
};
export const hypocrisyRating = {
  artist_name: 'Hypocrisy',
  festival_name: 'Wacken',
  rating: 5,
  year: 2022,
};

const handlers = [
  rest.get('http://localhost:8080/api/ratings', (req, res, ctx) => res(
    ctx.status(200),
    ctx.json([bloodbathRating, hypocrisyRating]),
  )),
  rest.post('http://localhost:8080/api/ratings', (req, res, ctx) => res(ctx.status(201))),
];

export const mockServer = setupServer(...handlers);
