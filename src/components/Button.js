import React from 'react';

const Button = (props) => (
  <button
    className="bg-pink-500 text-white active:bg-pink-600 font-bold uppercase text-sm px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1"
    type="button"
    style={{ transition: "all .15s ease" }}
    {...props}
  />
)

export default Button;
