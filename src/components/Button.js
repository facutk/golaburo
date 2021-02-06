import React from 'react';

const Button = (props) => (
  <button
    className="bg-pink-500 text-white active:bg-gray-700 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full"
    type="button"
    style={{ transition: "all .15s ease" }}
    {...props}
  />
)

export default Button;
