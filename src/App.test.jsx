import React from 'react';
import ReactDOM from 'react-dom';
import { shallow } from 'enzyme';
import toJSON from 'enzyme-to-json';
import App from './App';
import TabBar from './TabBar';

describe('App', () => {
  it('renders without crashing', () => {
    const div = document.createElement('div');
    ReactDOM.render(<App authState="signedIn" />, div);
    ReactDOM.unmountComponentAtNode(div);
  });

  it('should renders content correctly', () => {
    expect(toJSON(shallow(<App authState="signedIn" />))).toMatchSnapshot();
  });

  it('should render no content when not signedIn', () => {
    expect(shallow(<App authState="signIn" />).find(TabBar)).toHaveLength(0);
  });

  it('should render content when signedIn', () => {
    expect(shallow(<App authState="signedIn" />).find(TabBar)).toHaveLength(1);
  });
});
