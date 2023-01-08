import { rest } from 'msw';
import { setupServer } from 'msw/node';

export const TestUserId = 'MetalLover666';
export const TestToken = 'eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJNZXRhbExvdmVyNjY2IiwiaWF0IjoxNTE2MjM5MDIyfQ.JZ3R_3it-1K9ttH5NA80fpIsBhnW6DNsIzwB2zEFRmo7hgE-HhW3jJbArXNS0fw2Pcj-xrU-DMF8KoLr8_EJh2XdTDjaRqz859p0RJ1gPLovsQ8N1HeqeQXKi2mwDJe2rZhWILHdWZ1zmduCY5fF8jUYyBIqLRh1B44L_CBlgeEejKoJfw7V3WoZhxdLeW8SlS2PQ7kN0XIyzm-_TPq1j5QnNHRnXRIh8V7o9rBtdM7PVGEFTpzb1jC6bZ3W-aHuZEWk5e1kRTV8IOXiLf-xtPQ42Hn4j2F27mDg0h2PsgVWmNjr2eqc9y0izps-rmoXHnzmBzvbtGS2yytEFw_WAA';

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

export const unratedArtist = { artist_name: 'God Dethroned', image_url: 'https://god-dethrond-image.com' };

export const checkToken = (req) => {
  if (req.headers.get('authorization') !== `Bearer ${TestToken}`) {
    console.error('missing token');
    expect(true).toBeFalsy();
  }
};
const api = 'http://localhost:8080/api';

let ratedArtists;
let unratedArtists;
beforeEach(() => {
  ratedArtists = [bloodbathRating, hypocrisyRating];
  unratedArtists = [unratedArtist];
});
const handlers = [
  rest.get(`${api}/ratings`, (req, res, ctx) => {
    checkToken(req);
    return res(
      ctx.status(200),
      ctx.json(ratedArtists),
    );
  }),
  rest.post(`${api}/ratings`, async (req, res, ctx) => {
    checkToken(req);
    const body = await req.json();
    unratedArtists = unratedArtists.filter((a) => a.artist_name !== body.artist_name);
    if (ratedArtists.filter((a) => a.artist_name === body.artist_name).length === 0) {
      ratedArtists.push(body);
    }
    return res(ctx.status(201));
  }),
  rest.get(`${api}/festivals/wacken`, (req, res, ctx) => {
    checkToken(req);
    return res(
      ctx.status(200),
      ctx.json(unratedArtists),
    );
  }),
];

export const mockServer = setupServer(...handlers);
