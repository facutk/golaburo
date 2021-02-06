import React, { useEffect, useState } from 'react';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';

import Lexorank from './util/lexorank';
import { Button, Input } from './components';

const reorder = (list, startIndex, endIndex) => {
  const result = Array.from(list);
  const [removed] = result.splice(startIndex, 1);
  result.splice(endIndex, 0, removed);

  return result;
};

const grid = 8;

const getItemStyle = (isDragging, draggableStyle) => ({
  // some basic styles to make the items look a bit nicer
  userSelect: "none",
  padding: grid,
  margin: `0 0 ${grid}px 0`,

  // // change background colour if dragging
  background: isDragging ? "#fafafa" : "white",

  // styles we need to apply on draggables
  ...draggableStyle
});

const getListStyle = isDraggingOver => ({
  background: isDraggingOver ? "lightblue" : "white",
  padding: grid,
  width: 250
});

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

  const onDragEnd = (result) => {
    // dropped outside the list
    if (!result.destination) {
      return;
    }

    const newTodos = reorder(
      todos,
      result.source.index,
      result.destination.index
    );

    const { draggableId, destination } = result;
    const { index } = destination;
    const prevIndex = index - 1;

    let prevRank = '';
    if (prevIndex > -1 && prevIndex <= newTodos.length) {
      prevRank = newTodos[prevIndex].Rank;
    }

    const nextIndex = index + 1;
    let nextRank = '';
    if (nextIndex <= newTodos.length) {
      nextRank = newTodos[nextIndex].Rank;
    }

    const lexorank = new Lexorank();
    
    const [Rank, ok] = lexorank.insert(prevRank, nextRank);

    console.log({ ok, index, prevIndex, nextIndex, prevRank, nextRank, Rank });

    fetch(`/api/v1/todos/${draggableId}`, {
      method: 'PUT',
      body: JSON.stringify({ Rank })
    })

    const newNewTodos = newTodos.map((item) => {
      if (item.ID === draggableId) {
        return {
          ...item,
          Rank
        }
      }
      return item
    });

    setTodos(newNewTodos);
  }

  const isDisabled = !draft;

  return (
    <>
      <h2>Todos</h2>
      <form onSubmit={handleAddTodo} className='flex'>
        <Input
          value={draft}
          onChange={handleDraftChange}
          placeholder='Add Todo...'
        />
        <Button disabled={isDisabled}>
          Add
        </Button>
      </form>

      <DragDropContext onDragEnd={onDragEnd}>
        <Droppable droppableId="list">
          {(provided, snapshot) => (
            <div
            {...provided.droppableProps}
            ref={provided.innerRef}
            style={getListStyle(snapshot.isDraggingOver)}
          >
            {todos.map((item, index) => (
              <Draggable key={item.ID} draggableId={item.ID} index={index}>
                {(provided, snapshot) => (
                  <div
                    ref={provided.innerRef}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}
                    style={getItemStyle(
                      snapshot.isDragging,
                      provided.draggableProps.style
                    )}
                  >
                    <div
                      style={{
                        display: "flex",
                        justifyContent: "space-between"
                      }}
                    >
                      <button onClick={() => handleDelete(item)}>-</button>
                      <strong>
                        {item.Description}
                      </strong>
                      <span>
                        {item.Rank} {item.Rank ? '★' : '☆'}
                      </span>
                    </div>
                    <small>
                      {item.Created}
                    </small>
                  </div>
                )}
              </Draggable>
            ))}
            {provided.placeholder}
          </div>
          )}
        </Droppable>
      </DragDropContext>
    </>
  );
};

export default Todos;
