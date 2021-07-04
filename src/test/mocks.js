/* eslint no-console: 0 */

const getOperationAndName = (options) => {
  const split = options.query.split(/\s+/);
  return { operation: split[1], name: split[2].slice(0, -1) };
};

export const graphqlMock = (options, ...args) => {
  const { operation, name } = getOperationAndName(options);
  if (name === 'ListRatings') {
    return Promise.resolve([{
      artist: 'Bloodbath', festival: 'Wacken', year: 2015, rating: 5, comment: 'comment',
    }]);
  }
  console.warn(`API GraphQl ${operation} not mocked for [${name} with ${options.variables}]`, ...args);
  return new Error('API GraphQl request not mocked');
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
