import React, { useEffect, useState } from 'react';

const App = () => {
  const [hits, setHits] = useState();
  useEffect(() => {
    fetch('/hits').then(r => r.text()).then(setHits);
  }, []);

  return (
    <h1>{hits}</h1>
  );
};

export default App;
