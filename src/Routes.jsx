import React from 'react';
import { Route, Switch } from 'react-router-dom';
import Authentication from './Authentication';

const Routes = () => (
  <Switch>
    <Route exact path="/" component={Authentication} />
  </Switch>
);

export default Routes;
