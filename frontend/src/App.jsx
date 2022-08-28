import { QueryClient, QueryClientProvider } from 'react-query';

import './App.css';
import { Auth0Provider } from '@auth0/auth0-react';
import Login from './login';

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Auth0Provider
        domain="https://musicrating.eu.auth0.com"
        clientId={process.env.REACT_APP_CLIENT_ID}
        redirectUri={process.env.REACT_APP_DOMAIN_NAME}
      >
        <Login />
      </Auth0Provider>
    </QueryClientProvider>
  );
}

export default App;
