import React from 'react';
import { shallow, mount } from 'enzyme';
import toJSON from 'enzyme-to-json';
import { MemoryRouter } from 'react-router-dom';
import Routes from './Routes';
import App from './App';

describe('Routes', () => {
  it('should render correctly', () => {
    expect(toJSON(shallow(<Routes />))).toMatchSnapshot();
  });

  it('should render app if path is /', () => {
    expect(mount(<MemoryRouter initialEntries={['/']}><Routes /></MemoryRouter>).find(App)).toHaveLength(1);
  });
});
