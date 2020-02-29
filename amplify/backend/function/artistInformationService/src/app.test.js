import request from 'supertest';
import app from './app';
import fetch from 'node-fetch';

import 'whatwg-fetch';

jest.mock('node-fetch');

describe('artistInformationService', () => {
  describe('api', () => {
    afterEach(() => {
      app.stopServer()
    });

    it('should return artist', (done) => {
      //given
      process.env.SPOTIFY_CLIENT_ID = 'CLIENT_ID';
      process.env.SPOTIFY_CLIENT_SECRET = 'CLIENT_SECRET';
      const tokenEndpointUrl = 'https://accounts.spotify.com/api/token';
      const tokenEndpointResponse = {
        'access_token': 'token',
        'token_type': 'bearer',
        'expires_in': 3600,
      };
      const searchUrl = 'https://api.spotify.com/v1/search?q=bloodbath&type=artist&limit=1';
      const expectedImageUrl = 'https://320.com';
      const searchResponse = {
        'artists': {
          'items': [
            {
              'images': [
                {
                  'height': 640,
                  'url': 'https://image_640.com',
                  'width': 640
                },
                {
                  'height': 320,
                  'url': expectedImageUrl,
                  'width': 320
                },
                {
                  'height': 160,
                  'url': 'https://image_160.com',
                  'width': 160
                }
              ],
            }
          ],
        }
      };
      fetch.mockResolvedValueOnce(new Response(JSON.stringify(tokenEndpointResponse)));
      fetch.mockResolvedValueOnce(new Response(JSON.stringify(searchResponse)));

      //when & then
      request(app)
        .get('/api/v1/information/artists/bloodbath')
        .set('x-apigateway-event', encodeURIComponent(JSON.stringify({})))
        .set('x-apigateway-context', encodeURIComponent(JSON.stringify({})))
        .expect(200)
        .expect({ image: expectedImageUrl })
        .then(() => {
          expect(fetch).toHaveBeenCalledTimes(2);
          expect(fetch).toHaveBeenNthCalledWith(1, tokenEndpointUrl,
            {
              method: 'POST',
              headers: {
                'Authorization': `Basic ${Buffer.from('CLIENT_ID:CLIENT_SECRET').toString('base64')}`,
                'Content-Type': 'application/x-www-form-urlencoded'
              },
              body: 'grant_type=client_credentials'
            });
          expect(fetch).toHaveBeenNthCalledWith(2, searchUrl,
            { headers: { 'Authorization': 'Bearer token' } });
          done();
        });
    });
  });
});