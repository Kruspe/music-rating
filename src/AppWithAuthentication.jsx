import { Authenticator } from 'aws-amplify-react';
import Amplify from 'aws-amplify';
import App from './App';

import awsExports from './aws-exports';
import './App.css';

Amplify.configure(awsExports);

const AppWithAuthentication = () => (
  <Authenticator signUpConfig={{ defaultCountryCode: 49, hiddenDefaults: ['phone_number'] }}>
    <App />
  </Authenticator>
);

export default AppWithAuthentication;
