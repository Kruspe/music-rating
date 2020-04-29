export const ApiGetMock = (apiName, path, ...args) => {
  if (path === '/api/v1/ratings/bands/userId') {
    return Promise.resolve([{
      band: 'Bloodbath', festival: 'Wacken', year: 2015, rating: 5, comment: 'comment',
    }]);
  }
  console.warn(`API GET request not mocked for [${apiName}, ${path}]`, ...args);
  return new Error('API GET request not mocked');
};

export const ApiPostMock = (apiName, path, ...args) => {
  console.warn(`API POST request not mocked for [${apiName}, ${path}]`, ...args);
  return new Error('API POST request not mocked');
};