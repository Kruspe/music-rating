import Amplify from 'aws-amplify';
import { QueryClient, QueryClientProvider } from 'react-query';

import awsExports from './aws-exports';
import './App.css';

Amplify.configure(awsExports);

const queryClient = new QueryClient();

const App = () => (
  <QueryClientProvider client={queryClient}>
    <div>Hello World</div>
  </QueryClientProvider>
);

export default App;
