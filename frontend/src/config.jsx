import AuthenticationGuard from './components/AuthenticationGuard';
import Home from './home';
import Wacken from './wacken';
import Ratings from './ratings';
import MenuBar from './components/MenuBar';

const routesConfig = [
  {
    path: '/',
    element: <MenuBar />,
    children: [
      {
        path: '/',
        element: <Home />,
      },
      {
        path: '/ratings',
        element: <AuthenticationGuard component={Ratings} />,
      },
      {
        path: '/wacken',
        element: <AuthenticationGuard component={Wacken} />,
      },
    ],
  },
];

export default routesConfig;
