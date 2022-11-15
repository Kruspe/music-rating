import PropTypes from 'prop-types';
import { createTheme, ThemeProvider } from '@mui/material';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Auth0Provider } from '@auth0/auth0-react';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

const queryClient = new QueryClient();

export default function Providers({ children }) {
  return (
    <Auth0Provider
      domain="https://musicrating.eu.auth0.com"
      clientId={process.env.REACT_APP_CLIENT_ID}
      redirectUri={process.env.REACT_APP_DOMAIN_NAME.includes('localhost')
        ? process.env.REACT_APP_DOMAIN_NAME : `https://${process.env.REACT_APP_DOMAIN_NAME}`}
      audience={process.env.REACT_APP_DOMAIN_NAME.includes('localhost')
        ? process.env.REACT_APP_DOMAIN_NAME : `https://api.${process.env.REACT_APP_DOMAIN_NAME}`}
    >
      <ThemeProvider theme={darkTheme}>
        <QueryClientProvider client={queryClient}>
          {children}
        </QueryClientProvider>
      </ThemeProvider>
    </Auth0Provider>
  );
}

Providers.propTypes = {
  children: PropTypes.node.isRequired,
};
