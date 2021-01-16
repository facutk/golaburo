import React from 'react';
import {
  HashRouter as Router,
  Switch,
  Route
} from 'react-router-dom';

import Nav from './Nav';
import Hits from './Hits';

const App = () => (
  <Router>
    <Nav />
    <Switch>
      <Route path='/hits'>
        <Hits />
      </Route>
      <Route path='/'>
        <div>hello</div>
      </Route>
    </Switch>
  </Router>
  
);

export default App;
