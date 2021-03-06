import React, { Suspense } from 'react';
import {
  HashRouter as Router,
  Switch,
  Route
} from 'react-router-dom';
import { ThemeWrapper } from '@golaburo/uikit';

import '@golaburo/uikit/dist/style.css';

import Nav from './Nav';
import Dots from './Dots';

const Dnd = React.lazy(() => import('./Dnd'));
const Hits = React.lazy(() => import('./Hits'));
const Todos = React.lazy(() => import('./Todos'));

const App = () => (
  <Suspense fallback={<Dots />}>
    <Router>
      <ThemeWrapper>
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
      </ThemeWrapper>
    </Router>
  </Suspense>
);

export default App;
