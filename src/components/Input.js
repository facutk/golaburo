import React from 'react';

const Input = (props) => (
  <div className="mb-3 pt-0">
    <input
      type="text"
      placeholder="Placeholder"
      className="px-3 py-3 placeholder-gray-400 text-gray-700 relative bg-white bg-white rounded text-sm border border-gray-400 outline-none focus:outline-none focus:shadow-outline w-full"
      {...props}
    />
  </div>
);

export default Input;
