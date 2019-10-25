import React from 'react';
import { shallow } from 'enzyme';
import toJSON from 'enzyme-to-json';
import { Rating as RatingMaterialUI } from '@material-ui/lab';
import { API, Auth } from 'aws-amplify';
import Rating from './index';

let wrapper;
describe('Rating', () => {
  beforeEach(() => {
    wrapper = shallow(<Rating />);
    jest.clearAllMocks();
  });
  it('should render correctly', () => {
    expect(toJSON(wrapper)).toMatchSnapshot();
  });

  describe('form', () => {
    it('should sent data to api and should call event.preventDefault', async () => {
      const preventDefaultMock = jest.fn();
      const currentSessionMock = Promise.resolve({
        getAccessToken: () => ({ getJwtToken: () => ('Token') }),
      });
      const currentUserInfoMock = Promise.resolve({ id: 'userId' });
      const expectedInit = {
        header: { Authorization: 'Bearer Token' },
        body: {
          user: 'userId', band: 'band', festival: 'festival', year: 2018, rating: 4, comment: 'comment',
        },
      };
      const postSpy = jest.spyOn(API, 'post').mockImplementation((f) => f);
      jest.spyOn(Auth, 'currentSession').mockImplementation(() => currentSessionMock);
      jest.spyOn(Auth, 'currentUserInfo').mockImplementation(() => currentUserInfoMock);
      wrapper.find('#band').prop('onChange')({ target: { value: 'band' } });
      wrapper.find('#festival').prop('onChange')({ target: { value: 'festival' } });
      wrapper.find('#year').prop('onChange')({ target: { value: 2018 } });
      wrapper.find(RatingMaterialUI).prop('onChange')({}, 4);
      wrapper.find('#comment').prop('onChange')({ target: { value: 'comment' } });

      await wrapper.find('#rating-form').prop('onSubmit')({ preventDefault: preventDefaultMock });
      expect(preventDefaultMock).toHaveBeenCalled();
      expect(postSpy).toHaveBeenCalledWith('musicrating', '/bands', expectedInit);
    });

    it('should not submit data if band is not filled', async () => {
      const preventDefaultMock = jest.fn();
      const postSpy = jest.spyOn(API, 'post').mockImplementation((f) => f);
      wrapper.find('#band').prop('onChange')({ target: { value: '' } });
      await wrapper.find('#rating-form').prop('onSubmit')({ preventDefault: preventDefaultMock });
      wrapper.find('#band').prop('onChange')({ target: { value: '  ' } });
      await wrapper.find('#rating-form').prop('onSubmit')({ preventDefault: preventDefaultMock });
      wrapper.find('#band').prop('onChange')({ target: { value: undefined } });
      await wrapper.find('#rating-form').prop('onSubmit')({ preventDefault: preventDefaultMock });
      expect(preventDefaultMock).toHaveBeenCalledTimes(3);
      expect(postSpy).not.toHaveBeenCalled();
    });
  });

  describe('band', () => {
    it('should have empty string as initial value', () => {
      expect(wrapper.find('#band').prop('value')).toBe('');
    });
    it('should update TextField onChange', () => {
      wrapper.find('#band').prop('onChange')({ target: { value: 'band' } });
      expect(wrapper.find('#band').prop('value')).toBe('band');
    });
  });

  describe('festival', () => {
    it('should have empty string as initial value', () => {
      expect(wrapper.find('#festival').prop('value')).toBe('');
    });
    it('should update TextField onChange', () => {
      wrapper.find('#festival').prop('onChange')({ target: { value: 'festival' } });
      expect(wrapper.find('#festival').prop('value')).toBe('festival');
    });
  });

  describe('year', () => {
    it('should have empty string as initial value', () => {
      expect(wrapper.find('#year').prop('value')).toBe('');
    });
    it('should update TextField onChange', () => {
      wrapper.find('#year').prop('onChange')({ target: { value: 2018 } });
      expect(wrapper.find('#year').prop('value')).toBe(2018);
    });
  });

  describe('rating', () => {
    it('should have 0 as initial state', () => {
      expect(wrapper.find(RatingMaterialUI).prop('value')).toBe(1);
    });
    it('should update rating onChange', () => {
      wrapper.find(RatingMaterialUI).prop('onChange')({}, 2);
      expect(wrapper.find(RatingMaterialUI).prop('value')).toBe(2);
    });
  });

  describe('comment', () => {
    it('should have empty string as initial state', () => {
      expect(wrapper.find('#comment').prop('value')).toBe('');
    });
    it('should update rating onChange', () => {
      wrapper.find('#comment').prop('onChange')({ target: { value: 'comment' } });
      expect(wrapper.find('#comment').prop('value')).toBe('comment');
    });
  });
});
