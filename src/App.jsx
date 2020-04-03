import React from 'react';
import Amplify from 'aws-amplify';
import PropTypes from 'prop-types';

import awsExports from './aws-exports';
import './App.css';
import TabBar from './tabs/TabBar';
import UserProvider from './provider/UserProvider';

Amplify.configure(awsExports);

const App = (props) => {
  const { authState } = props;
  return (authState === 'signedIn' && (
  <UserProvider>
    <TabBar />
  </UserProvider>
  ));
};

App.propTypes = {
  authState: PropTypes.string,
};

App.defaultProps = {
  authState: 'signedOut',
};

export default App;
