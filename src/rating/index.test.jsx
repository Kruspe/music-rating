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
      const expectedInit = {
        header: { Authorization: 'Bearer Token' },
        body: { bandName: 'bandName', rating: 4, comment: 'comment' },
      };
      const postSpy = jest.spyOn(API, 'post').mockImplementation(f => f);
      jest.spyOn(Auth, 'currentSession').mockImplementation(() => currentSessionMock);
      wrapper.find('#bandName').prop('onChange')({ target: { value: 'bandName' } });
      wrapper.find(RatingMaterialUI).prop('onChange')({}, 4);
      wrapper.find('#comment').prop('onChange')({ target: { value: 'comment' } });

      await wrapper.find('#rating-form').prop('onSubmit')({ preventDefault: preventDefaultMock });
      expect(preventDefaultMock).toHaveBeenCalled();
      expect(postSpy).toHaveBeenCalledWith('musicrating', '/bands', expectedInit);
    });

    it('should not submit data if bandName is not filled', async () => {
      const preventDefaultMock = jest.fn();
      const postSpy = jest.spyOn(API, 'post').mockImplementation(f => f);
      wrapper.find('#bandName').prop('onChange')({ target: { value: '' } });
      await wrapper.find('#rating-form').prop('onSubmit')({ preventDefault: preventDefaultMock });
      wrapper.find('#bandName').prop('onChange')({ target: { value: '  ' } });
      await wrapper.find('#rating-form').prop('onSubmit')({ preventDefault: preventDefaultMock });
      wrapper.find('#bandName').prop('onChange')({ target: { value: undefined } });
      await wrapper.find('#rating-form').prop('onSubmit')({ preventDefault: preventDefaultMock });
      expect(preventDefaultMock).toHaveBeenCalledTimes(3);
      expect(postSpy).not.toHaveBeenCalled();
    });
  });

  describe('bandName', () => {
    it('should have empty string as initial value', () => {
      expect(wrapper.find('#bandName').prop('value')).toBe('');
    });
    it('should update TextField onChange', () => {
      wrapper.find('#bandName').prop('onChange')({ target: { value: 'bandName' } });
      expect(wrapper.find('#bandName').prop('value')).toBe('bandName');
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
