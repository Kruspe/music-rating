import { QueryClient, QueryClientProvider } from 'react-query';

import './App.css';

const queryClient = new QueryClient();

const App = () => (
  <QueryClientProvider client={queryClient}>
    <div>Hello World</div>
  </QueryClientProvider>
);

export default App;
