/* eslint no-console: 0 */
export const apiGetMock = (apiName, path, ...args) => {
  if (path === '/api/v1/ratings/bands/userId') {
    return Promise.resolve([{
      artist: 'Bloodbath', festival: 'Wacken', year: 2015, rating: 5, comment: 'comment',
    }]);
  }
  console.warn(`API GET request not mocked for [${apiName}, ${path}]`, ...args);
  return new Error('API GET request not mocked');
};

export const apiPostMock = (apiName, path, ...args) => {
  if (path === '/api/v1/ratings/bands') {
    return Promise.resolve('Success');
  }
  console.warn(`API POST request not mocked for [${apiName}, ${path}]`, ...args);
  return new Error('API POST request not mocked');
};

export const storageGetMock = (key, ...args) => {
  if (key === 'wacken.json') {
    return Promise.resolve('wackenLink');
  }
  console.warn(`STORAGE GET request not mocked for [${key}]`, ...args);
  return new Error('STORAGE GET request not mocked');
};

export const fetchMock = (path, ...args) => {
  if (path === 'wackenLink') {
    return new Response(JSON.stringify([
      { artist: 'Bloodbath', image: 'bloodbathImage' },
      { artist: 'Megadeth', image: 'megadethImage' },
      { artist: 'Vader', image: 'vaderImage' },
    ]));
  }
  console.warn('window.fetch is not mocked for this call', path, ...args);
  return new Error('This must be mocked!');
};
