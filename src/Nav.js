import React from 'react';
import { Link } from 'react-router-dom';

const Nav = () => (
  <nav>
    <div>
      <Link to='/'>golaburo</Link>
    </div>
    <div>
      <Link to='/hits'>Hits</Link>
      <Link to='/dnd'>Dnd</Link>
    </div>
  </nav>
);

export default Nav;
