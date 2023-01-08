import './App.css';
import { useAuth0 } from '@auth0/auth0-react';
import Wacken from './wacken';
import Home from './home';
import MenuBar from './components/menu-bar';

function App() {
  const { isAuthenticated } = useAuth0();
  return (
    <>
      <MenuBar />
      {isAuthenticated ? <Wacken /> : <Home />}
    </>
  );
}

export default App;
