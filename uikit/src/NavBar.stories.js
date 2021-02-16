import React from 'react';

import NavBar from './NavBar';

const Template = (args) => (
  <NavBar {...args}>
    <NavBar.Item>
      Home
    </NavBar.Item>
    <NavBar.Item>
      About
    </NavBar.Item>
    <NavBar.Item>
      Settings
    </NavBar.Item>
  </NavBar>
);

export const Default = Template.bind({});

const Story = {
  title: 'NavBar',
  component: NavBar
}

export default Story;