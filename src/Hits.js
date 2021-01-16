import React, { useEffect, useState } from 'react';

const Hits = () => {
  const [hits, setHits] = useState();
  useEffect(() => {
    fetch('/api/v1/hits').then(r => r.text()).then(setHits);
  }, []);

  return (
    <h1>{hits}</h1>
  );
};

export default Hits;
