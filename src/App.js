import React, { Suspense } from 'react';
import {
  HashRouter as Router,
  Switch,
  Route
} from 'react-router-dom';

import '@golaburo/uikit/dist/style.css';

import Nav from './Nav';
import Dots from './Dots';

const Dnd = React.lazy(() => import('./Dnd'));
const Hits = React.lazy(() => import('./Hits'));
const Todos = React.lazy(() => import('./Todos'));

const App = () => (
  <Suspense fallback={<Dots />}>
    <Router>
      <Nav />
      <Switch>
        <Route path='/dnd'>
          <Dnd />
        </Route>
        <Route path='/hits'>
          <Hits />
        </Route>
        <Route path='/'>
          <Todos />
        </Route>
      </Switch>
    </Router>
  </Suspense>
);

export default App;
