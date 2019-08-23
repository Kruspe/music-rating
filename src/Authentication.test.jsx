import React from 'react';
import { shallow } from 'enzyme';
import toJSON from 'enzyme-to-json';
import Authentication from './Authentication';

describe('Authentication', () => {
  it('should render correctly', () => {
    expect(toJSON(shallow(<Authentication />))).toMatchSnapshot();
  });
});
