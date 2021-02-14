import React from 'react';

const Button = ({
  children, onClick
}) => {
  return (
    <button
      onClick={onClick}
      className='bg-yellow-600'
    >
      {children}
    </button>
  );
}

export default Button;
