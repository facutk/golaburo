import React from 'react';

import Input from './Input';

const Template = (args) => <Input {...args} />;

export const Primary = Template.bind({});
Primary.args = {
  value: null,
  placeholder: 'Placeholder ...',
  type: 'text'
};

const Story = {
  title: 'Input',
  component: Input,
  argTypes: {
    value: { control: 'text' },
  },
}

export default Story;