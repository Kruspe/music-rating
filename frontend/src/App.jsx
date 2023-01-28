import './App.css';
import { Auth0Provider } from '@auth0/auth0-react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { createTheme, CssBaseline, ThemeProvider } from '@mui/material';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import routesConfig from './config';
import MenuBar from './components/MenuBar';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

const router = createBrowserRouter(routesConfig);

const queryClient = new QueryClient();

function App() {
  return (
    <Auth0Provider
      domain="https://musicrating.eu.auth0.com"
      clientId={process.env.REACT_APP_CLIENT_ID}
      redirectUri={process.env.REACT_APP_DOMAIN_NAME.includes('localhost')
        ? `${process.env.REACT_APP_DOMAIN_NAME}/wacken` : `https://${process.env.REACT_APP_DOMAIN_NAME}/wacken`}
      audience={process.env.REACT_APP_DOMAIN_NAME.includes('localhost')
        ? undefined : `https://api.${process.env.REACT_APP_DOMAIN_NAME}`}
    >
      <ThemeProvider theme={darkTheme}>
        <CssBaseline />
        <QueryClientProvider client={queryClient}>
          <MenuBar />
          <RouterProvider router={router} />
        </QueryClientProvider>
      </ThemeProvider>
    </Auth0Provider>
  );
}

export default App;
