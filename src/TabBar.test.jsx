import React from 'react';
import { shallow } from 'enzyme';
import toJSON from 'enzyme-to-json';
import { Tabs } from '@material-ui/core';
import TabBar from './TabBar';
import Overview from './overview';
import Rating from './rating';

describe('TabBar', () => {
  it('should render correctly', () => {
    expect(toJSON(shallow(<TabBar />))).toMatchSnapshot();
  });

  it('should show overview as initial tab', () => {
    const wrapper = shallow(<TabBar />);
    expect(wrapper.find(Overview)).toHaveLength(1);
    expect(wrapper.find(Rating)).toHaveLength(0);
  });

  describe('Tabs', () => {
    it('should change tab to Overview when index is 0', () => {
      const wrapper = shallow(<TabBar />);
      wrapper.find(Tabs).prop('onChange')({}, 1);
      expect(wrapper.find(Rating)).toHaveLength(1);
      wrapper.find(Tabs).prop('onChange')({}, 0);
      expect(wrapper.find(Rating)).toHaveLength(0);
      expect(wrapper.find(Overview)).toHaveLength(1);
    });
    it('should change tab to Rating when index is 1', () => {
      const wrapper = shallow(<TabBar />);
      expect(wrapper.find(Overview)).toHaveLength(1);
      wrapper.find(Tabs).prop('onChange')({}, 1);
      expect(wrapper.find(Rating)).toHaveLength(1);
      expect(wrapper.find(Overview)).toHaveLength(0);
    });
  });
});
