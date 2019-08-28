import React from 'react';
import { shallow } from 'enzyme';
import toJSON from 'enzyme-to-json';
import { Rating as RatingMaterialUI } from '@material-ui/lab';
import Rating from './Rating';


describe('Rating', () => {
  it('should render correctly', () => {
    expect(toJSON(shallow(<Rating />))).toMatchSnapshot();
  });
  it('should update rating onChange', () => {
    const wrapper = shallow(<Rating />);
    wrapper.find(RatingMaterialUI).prop('onChange')({}, 2);
    expect(wrapper.find(RatingMaterialUI).prop('value')).toBe(2);
  });
});
