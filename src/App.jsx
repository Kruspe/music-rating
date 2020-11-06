import { useState } from 'react';
import Amplify, { Auth } from 'aws-amplify';
import { Button, Grid, TextField } from '@material-ui/core';

import awsExports from './aws-exports';
import './App.css';
import TabBar from './tabs/TabBar';

Amplify.configure(awsExports);

const App = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [authState, setAuthState] = useState();
  const login = async (event) => {
    event.preventDefault();
    setAuthState(await Auth.signIn(username, password));
  };

  return (
    <>
      {authState ? <TabBar /> : (
        <form onSubmit={login}>
          <Grid container direction="column">
            <Grid item>
              <TextField
                label="Username"
                name="username"
                id="username"
                value={username}
                onChange={(event) => setUsername(event.target.value)}
              />
            </Grid>
            <Grid item>
              <TextField
                type="password"
                label="Password"
                name="password"
                id="password"
                value={password}
                onChange={(event) => setPassword(event.target.value)}
              />
            </Grid>
            <Grid item>
              <Button type="submit">Sign In</Button>
            </Grid>
          </Grid>
        </form>
      )}
    </>
  );
};

export default App;
