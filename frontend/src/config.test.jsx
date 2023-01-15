import { renderWithRouteProvider, screen } from '../test/test-utils';
import Wacken from './wacken';
import Home from './home';
import AuthenticationGuard from './components/AuthenticationGuard';

jest.mock('./Home');
jest.mock('./Wacken');

jest.mock('./components/AuthenticationGuard');

function TestComponent() {
  return <div />;
}

beforeEach(() => {
  AuthenticationGuard.mockImplementation(({ component }) => <p>{`Authenticated ${component()}`}</p>);
});

test('should have correct / route', () => {
  Home.mockImplementation(() => <p>HomeComponent</p>);

  renderWithRouteProvider(<TestComponent />);
  expect(screen.getByText('HomeComponent')).toBeVisible();
});

test('should have correct /wacken route', () => {
  Wacken.mockImplementation(() => 'WackenComponent');

  renderWithRouteProvider(<TestComponent />, { initialEntries: ['/wacken'] });
  expect(screen.getByText('Authenticated WackenComponent')).toBeVisible();
});
