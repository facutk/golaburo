import React from 'react';
import { Link } from 'react-router-dom';

const Nav = () => (
  <nav className="nav">
    <div className="nav-left">
      <Link to='/'>golaburo</Link>
    </div>
    <div className="nav-center">
      
    </div>
    <div className="nav-right">
      
      <Link to='/hits'>Hits</Link>
      <Link to='/dnd'>Dnd</Link>
    </div>
  </nav>
);

export default Nav;
