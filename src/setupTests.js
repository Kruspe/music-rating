// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import { API, Auth, Storage } from 'aws-amplify';
import { graphqlMock, fetchMock, storageGetMock } from './test/mocks';

beforeEach(() => {
  window.fetch = jest.fn().mockImplementation(fetchMock);

  jest.spyOn(API, 'graphql').mockImplementation(graphqlMock);

  jest.spyOn(Auth, 'currentUserInfo').mockResolvedValue({ id: 'userId' });
  const currentSessionMock = {
    getAccessToken: () => ({ getJwtToken: () => ('token') }),
  };
  jest.spyOn(Auth, 'currentSession').mockResolvedValue(currentSessionMock);
  const Username = 'username';
  jest.spyOn(Auth, 'signIn').mockResolvedValueOnce({ Username });
  jest.spyOn(Auth, 'signUp').mockResolvedValueOnce({ user: { Username } });

  jest.spyOn(Storage, 'get').mockImplementation(storageGetMock);
});

afterEach(() => {
  jest.restoreAllMocks();
});
