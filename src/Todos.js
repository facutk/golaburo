import React, { useEffect, useState } from 'react';

const Todos = () => {
  const [todos, setTodos] = useState([]);
  const [draft, setDraft] = useState('');

  useEffect(() => {
    fetch('/api/v1/todos')
      .then(r => r.json())
      .then(setTodos);
  }, []);

  const handleDraftChange = (e) => {
    setDraft(e.target.value);
  }

  const handleAddTodo = (e) => {
    e.preventDefault();
    const newTodo = {
      Description: draft
    };

    fetch('/api/v1/todos', {
      method: 'POST',
      body: JSON.stringify(newTodo)
    })
      .then(r => r.json())
      .then((response) => setTodos(todos.concat(response)));
    
    setDraft('');
  }

  const handleDelete = (todoItem) => {
    fetch(`/api/v1/todos/${todoItem.ID}`, {
      method: 'DELETE',
    })
      .then(() => {
        const newTodos = todos.filter((todo) => todo.ID !== todoItem.ID);
        setTodos(newTodos);
      })
  }

  return (
    <>
      <h2>Todos</h2>
      <form onSubmit={handleAddTodo}>
        <input value={draft} onChange={handleDraftChange} />
        <button>
          Add
        </button>
      </form>
      <ul>
        {todos.map((todo) => (
          <li key={todo.ID}>
            <button onClick={() => handleDelete(todo)}>-</button> {todo.Description}
          </li>
        ))}
      </ul>
    </>
  );
};

export default Todos;
