import React from 'react';
import { shallow } from 'enzyme';
import toJSON from 'enzyme-to-json';
import { Authenticator } from 'aws-amplify-react';
import Authentication from './Authentication';

describe('Authentication', () => {
  it('should render correctly', () => {
    expect(toJSON(shallow(<Authentication />))).toMatchSnapshot();
  });

  it('should have right signUpConfig', () => {
    const expectedSignUpConfig = { defaultCountryCode: 49, hiddenDefaults: ['phone_number'] };
    const wrapper = shallow(<Authentication />);
    expect(wrapper.find(Authenticator).prop('signUpConfig')).toEqual(expectedSignUpConfig);
  });
});
