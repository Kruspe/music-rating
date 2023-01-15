import AuthenticationGuard from './components/AuthenticationGuard';
import Home from './home';
import Wacken from './wacken';

const routesConfig = [
  {
    path: '/',
    element: <Home />,
  },
  {
    path: '/wacken',
    element: <AuthenticationGuard component={Wacken} />,
  },
];

export default routesConfig;
