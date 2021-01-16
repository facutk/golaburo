import React, { useEffect, useState } from 'react';

const Todos = () => {
  const [todos, setTodos] = useState([]);
  useEffect(() => {
    fetch('/api/v1/todos').then(r => r.json()).then(setTodos);
  }, []);

  return (
    <>
      <h2>Todos</h2>
      <ul>
        {todos.map((todo) => (
          <li key={todo.ID}>
            {todo.Description}
          </li>
        ))}
      </ul>
    </>
  );
};

export default Todos;
