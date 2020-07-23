import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

jest.mock('./tabs/TabBar', () => (() => (<p>TabBar</p>)));

describe('App', () => {
  it('should render no content when not signedIn', () => {
    const { container } = render(<App authState="signIn" />);
    expect(container).toBeEmptyDOMElement();
  });

  it('should render TabBar when signedIn', async () => {
    render(<App authState="signedIn" />);
    expect(await screen.findByText('TabBar')).toBeVisible();
  });
});
