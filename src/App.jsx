import { useEffect, useState } from 'react';
import Amplify from 'aws-amplify';
import { QueryClient, QueryClientProvider } from 'react-query';
import { AmplifyAuthenticator, AmplifySignUp } from '@aws-amplify/ui-react';
import { AuthState, onAuthUIStateChange } from '@aws-amplify/ui-components';

import awsExports from './aws-exports';
import './App.css';

Amplify.configure(awsExports);

const queryClient = new QueryClient();

const App = () => {
  const [authState, setAuthState] = useState(AuthState.SignIn);
  useEffect(() => onAuthUIStateChange((nextAuthState) => {
    setAuthState(nextAuthState);
  }), []);

  return authState === AuthState.SignedIn ? (
    <QueryClientProvider client={queryClient}>
      <div>Hello World</div>
    </QueryClientProvider>
  ) : (
    <AmplifyAuthenticator initalAuthState={AuthState.SignedIn}>
      <AmplifySignUp
        slot="sign-up"
        formFields={[
          { type: 'username' },
          { type: 'password' },
          { type: 'email' },
        ]}
      />
    </AmplifyAuthenticator>
  );
};

export default App;
