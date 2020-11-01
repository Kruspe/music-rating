import Amplify from 'aws-amplify';
import PropTypes from 'prop-types';

import awsExports from './aws-exports';
import './App.css';
import TabBar from './tabs/TabBar';

Amplify.configure(awsExports);

const App = ({ authState }) => (authState === 'signedIn' && (<TabBar />));

App.propTypes = {
  authState: PropTypes.string,
};

App.defaultProps = {
  authState: 'signedOut',
};

export default App;
