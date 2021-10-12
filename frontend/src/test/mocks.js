/* eslint no-console: 0 */

const fetchMock = (path, ...args) => {
  console.warn('window.fetch is not mocked for this call', path, ...args);
  return new Error('This must be mocked!');
};

export default fetchMock;
