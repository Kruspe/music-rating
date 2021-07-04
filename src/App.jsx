import Amplify from 'aws-amplify';

import awsExports from './aws-exports';
import './App.css';

Amplify.configure(awsExports);

const App = () => (<div>Hello World</div>);

export default App;
