import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import './App.css';
import { Auth0Provider } from '@auth0/auth0-react';
import { createTheme, ThemeProvider } from '@mui/material';
import Content from './Content';

const queryClient = new QueryClient();

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

function App() {
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
          <Content />
        </QueryClientProvider>
      </ThemeProvider>
    </Auth0Provider>
  );
}

export default App;
