import React from 'react';
import { shallow } from 'enzyme';
import toJSON from 'enzyme-to-json';
import { API, Auth } from 'aws-amplify';
import Overview from './index';

const expectedRatings = [{ band: 'band', rating: 5 }];

describe('Overview', () => {
  let getBandsSpy;
  beforeEach(() => {
    getBandsSpy = jest.spyOn(API, 'get').mockImplementation(() => expectedRatings);
    getBandsSpy.mockClear();
    const currentSessionMock = Promise.resolve({
      getAccessToken: () => ({ getJwtToken: () => ('Token') }),
    });
    const currentUserInfoMock = Promise.resolve({ id: 'userId' });
    jest.spyOn(Auth, 'currentUserInfo').mockImplementation(() => currentUserInfoMock);
    jest.spyOn(Auth, 'currentSession').mockImplementation(() => currentSessionMock);
  });
  it('should render correctly', (done) => {
    const wrapper = shallow(<Overview />);
    process.nextTick(() => {
      expect(toJSON(wrapper)).toMatchSnapshot();
      done();
    });
  });

  it('should get user data', async () => {
    const wrapper = shallow(<Overview />);

    process.nextTick(() => {
      expect(getBandsSpy).toHaveBeenCalledWith('musicrating', '/bands/userId', { header: { Authorization: 'Bearer Token' } });
      expect(wrapper.state().ratings).toEqual(expectedRatings);
    });
  });
});
