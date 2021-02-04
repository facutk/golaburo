import styled from 'styled-components';

const Dots = styled.span`
  &::after {
    display: inline-block;
    animation: ellipsis 1s infinite;
    content: ".";
    width: 1em;
    text-align: left;
    white-space: pre;
  }
  @keyframes ellipsis {
    0% {
      content: ".";
    }
    16% {
      content: "..";
    }
    32% {
      content: "...";
    }
    48% {
      content: " ..";
    }
    64% {
      content: "  .";
    }
    80% {
      content: "";
    }
  }
`;

export default Dots;
