import React from 'react';
import { shallow } from 'enzyme';
import toJSON from 'enzyme-to-json';
import { Rating as RatingMaterialUI } from '@material-ui/lab';
import Rating from './Rating';

let wrapper;
describe('Rating', () => {
  beforeEach(() => {
    wrapper = shallow(<Rating />);
  });
  it('should render correctly', () => {
    expect(toJSON(wrapper)).toMatchSnapshot();
  });

  describe('bandName', () => {
    it('should have correct initial value', () => {
      expect(wrapper.find('#bandName').prop('value')).toBe('');
    });
    it('should update TextField onChange', () => {
      wrapper.find('#bandName').prop('onChange')({ target: { value: 'bandName' } });
      expect(wrapper.find('#bandName').prop('value')).toBe('bandName');
    });
  });

  describe('rating', () => {
    it('should have correct initial state', () => {
      expect(wrapper.find(RatingMaterialUI).prop('value')).toBe(0);
    });
    it('should update rating onChange', () => {
      wrapper.find(RatingMaterialUI).prop('onChange')({}, 2);
      expect(wrapper.find(RatingMaterialUI).prop('value')).toBe(2);
    });
  });

  describe('comment', () => {
    it('should have correct initial state', () => {
      expect(wrapper.find('#comment').prop('value')).toBe('');
    });
    it('should update rating onChange', () => {
      wrapper.find('#comment').prop('onChange')({ target: { value: 'comment' } });
      expect(wrapper.find('#comment').prop('value')).toBe('comment');
    });
  });
});
