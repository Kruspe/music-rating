import React from 'react';
import { withAuthenticator } from 'aws-amplify-react';
import Amplify from 'aws-amplify';
import PropTypes from 'prop-types';

import awsExports from './aws-exports';
import './App.css';

Amplify.configure(awsExports);

export const App = (props) => {
  const { authState } = props;
  return <div>{authState}</div>;
};
App.propTypes = {
  authState: PropTypes.string.isRequired,
};
export default withAuthenticator(App);
