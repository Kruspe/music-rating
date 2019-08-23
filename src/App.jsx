import React from 'react';
import Amplify from 'aws-amplify';
import PropTypes from 'prop-types';

import awsExports from './aws-exports';
import './App.css';

Amplify.configure(awsExports);

const App = (props) => {
  const { authState } = props;
  return (authState === 'signedIn' && <div id="content">blub</div>);
};

App.propTypes = {
  authState: PropTypes.string,
};

export default App;
