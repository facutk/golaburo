import tw from 'tailwind-styled-components';

const Button = tw.button`
  bg-red-500
  hover:bg-red-400
  text-white
  font-bold
  py-2
  px-4
  ml-2
  first:ml-0
  rounded
  ${p => p.disabled ? 'bg-blue-500 hover:bg-blue-500 cursor-not-allowed' : ''}
`;

export default Button;
