import React from 'react';
import Amplify from 'aws-amplify';
import PropTypes from 'prop-types';

import awsExports from './aws-exports';
import './App.css';
import TabBar from './TabBar';

Amplify.configure(awsExports);

const App = (props) => {
  const { authState } = props;
  return (authState === 'signedIn' && <TabBar />);
};

App.propTypes = {
  authState: PropTypes.string,
};

export default App;
