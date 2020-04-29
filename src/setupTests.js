import '@testing-library/jest-dom/extend-expect';
import { queryCache } from 'react-query';
import { API, Auth } from 'aws-amplify';
import { ApiGetMock, ApiPostMock } from './test/mocks';

beforeEach(() => {
  queryCache.clear();

  jest.spyOn(API, 'get').mockImplementation(ApiGetMock);
  jest.spyOn(API, 'post').mockImplementation(ApiPostMock);
  jest.spyOn(Auth, 'currentUserInfo').mockResolvedValue({ id: 'userId' });
  const currentSessionMock = {
    getAccessToken: () => ({ getJwtToken: () => ('token') }),
  };
  jest.spyOn(Auth, 'currentSession').mockResolvedValue(currentSessionMock);
});

afterEach(() => {
  jest.restoreAllMocks();
});