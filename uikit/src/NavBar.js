import tw from 'tailwind-styled-components';
import React from 'react';

const NavBar = tw.div`
  bg-white
  border-b
  border-gray-100
  px-4
  py-2
  mb-4
`;

const InnerItem = tw.div`
  inline-block
  py-2
  px-4
  rounded
`;

const Item = ({ NavLink = tw.a``, children, ...props }) => (
  <NavLink
    activeClassName='inline-block bg-gray-100 rounded'
    {...props}
  >
    <InnerItem>
      {children}
    </InnerItem>  
  </NavLink>
);

NavBar.Item = Item;

export default NavBar;
