import React from 'react';
import { Authenticator } from 'aws-amplify-react';
import Amplify from 'aws-amplify';
import App from './App';

import awsExports from './aws-exports';
import './App.css';

Amplify.configure(awsExports);

const Authentication = () => (
  <Authenticator>
    <App />
  </Authenticator>
);

export default Authentication;
