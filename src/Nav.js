import React from 'react';
import { NavLink } from 'react-router-dom';
import { NavBar } from '@golaburo/uikit';

const Nav = () => (
  <NavBar>
    <NavBar.Item NavLink={NavLink} to='/' exact>golaburo</NavBar.Item>
    <NavBar.Item NavLink={NavLink} to='/hits'>Hits</NavBar.Item>
    <NavBar.Item NavLink={NavLink} to='/dnd'>Dnd</NavBar.Item>
  </NavBar>
);

export default Nav;
