import '@testing-library/jest-dom/extend-expect';
import { queryCache } from 'react-query';
import { API, Auth, Storage } from 'aws-amplify';
import { apiGetMock, apiPostMock, fetchMock, storageGetMock } from './test/mocks';

beforeEach(() => {
  queryCache.clear();

  window.fetch = jest.fn().mockImplementation(fetchMock);

  jest.spyOn(API, 'get').mockImplementation(apiGetMock);
  jest.spyOn(API, 'post').mockImplementation(apiPostMock);
  jest.spyOn(Auth, 'currentUserInfo').mockResolvedValue({ id: 'userId' });
  const currentSessionMock = {
    getAccessToken: () => ({ getJwtToken: () => ('token') }),
  };
  jest.spyOn(Auth, 'currentSession').mockResolvedValue(currentSessionMock);
  jest.spyOn(Storage, 'get').mockImplementation(storageGetMock);
});

afterEach(() => {
  jest.restoreAllMocks();
});