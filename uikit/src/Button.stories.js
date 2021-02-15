import React from 'react';

import Button from './Button';

const Template = (args) => (
  <>
    <Button {...args} />
    <Button {...args} disabled />
  </>
);

export const Primary = Template.bind({});
Primary.args = {
  primary: true,
  children: 'Primary',
  disabled: false
};

const Story = {
  title: 'Button',
  component: Button,
  argTypes: {
    backgroundColor: { control: 'color' },
  },
}

export default Story;