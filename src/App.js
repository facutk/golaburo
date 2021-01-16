import React from 'react';
import {
  HashRouter as Router,
  Switch,
  Route
} from 'react-router-dom';

import Hits from './Hits';

const App = () => (
  <Router>
    <Switch>
      <Route path='/hello'>
        <div>hello</div>
      </Route>
      <Route path='/'>
        <Hits />
      </Route>
    </Switch>
  </Router>
  
);

export default App;
