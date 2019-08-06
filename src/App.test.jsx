import React from 'react';
import ReactDOM from 'react-dom';
import { shallow } from 'enzyme';
import toJSON from 'enzyme-to-json';
import { App } from './App';

describe('App', () => {
  it('renders without crashing', () => {
    const div = document.createElement('div');
    ReactDOM.render(<App authState="loggedIn" />, div);
    ReactDOM.unmountComponentAtNode(div);
  });

  it('should render correctly when logged in', () => {
    expect(toJSON(shallow(<App authState="loggedIn" />))).toMatchSnapshot();
  });
});
