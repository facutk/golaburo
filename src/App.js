import React from 'react';
import {
  HashRouter as Router,
  Switch,
  Route
} from 'react-router-dom';

import Nav from './Nav';
import Hits from './Hits';
import Todos from './Todos';

const App = () => (
  <Router>
    <Nav />
    <Switch>
      <Route path='/hits'>
        <Hits />
      </Route>
      <Route path='/'>
        <div>Landing</div>
        <Todos />
      </Route>
    </Switch>
  </Router>
  
);

export default App;
