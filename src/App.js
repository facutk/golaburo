import React from 'react';

const App = () => {
  const handleClick = () => {
    fetch('/ping').then(r => r.text()).then(console.log);
  };

  return (
    <div>
      hello react
      <button onClick={handleClick}>click me</button>
    </div>
  );
};

export default App;
