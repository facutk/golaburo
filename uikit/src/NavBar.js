import tw from 'tailwind-styled-components';
import React from 'react';

const NavBar = tw.div`
  bg-red-500
`;

const InnerItem = tw.div`
  border-gray-400
  inline-block
  py-2
  px-4
`;

const Item = ({ NavLink = tw.a``, children, ...props }) => (
  <NavLink
    activeClassName='bg-blue-500 inline-block'
    {...props}
  >
    <InnerItem>
      {children}
    </InnerItem>  
  </NavLink>
);

NavBar.Item = Item;

export default NavBar;
