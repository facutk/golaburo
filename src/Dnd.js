import React, { Component } from "react";
import { DragDropContext, Droppable, Draggable } from "react-beautiful-dnd";

import Lexorank from './util/lexorank';

// a little function to help us with reordering the result
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

class Dnd extends Component {
  constructor(props) {
    super(props);
    this.state = {
      items: []
    };
    this.onDragEnd = this.onDragEnd.bind(this);
  }

  componentDidMount() {
    fetch('/api/v1/todos')
      .then(r => r.json())
      .then((response) => {
        const todos = response.map(item => ({
          id: item.ID,
          content: item.Description,
          ...item
        }));
        this.setState({
          items: todos
        });
      });
  }

  onDragEnd(result) {
    // dropped outside the list
    if (!result.destination) {
      return;
    }

    const items = reorder(
      this.state.items,
      result.source.index,
      result.destination.index
    );

    const { draggableId, destination } = result;
    const { index } = destination;
    const prevIndex = index - 1;

    let prevRank = '';
    if (prevIndex > -1 && prevIndex <= items.length) {
      prevRank = items[prevIndex].Rank;
    }

    const nextIndex = index + 1;
    let nextRank = '';
    if (nextIndex <= items.length) {
      nextRank = items[nextIndex].Rank;
    }


    const lexorank = new Lexorank();
    
    const [Rank, ok] = lexorank.insert(prevRank, nextRank);

    console.log({ ok, index, prevIndex, nextIndex, prevRank, nextRank, Rank });

    fetch(`/api/v1/todos/${draggableId}`, {
      method: 'PUT',
      body: JSON.stringify({ Rank })
    })

    const newItems = items.map((item) => {
      if (item.ID === draggableId) {
        return {
          ...item,
          Rank
        }
      }
      return item
    });

    this.setState({
      items: newItems
    });
  }

  // Normally you would want to split things out into separate components.
  // But in this example everything is just done in one place for simplicity
  render() {
    return (
      <DragDropContext onDragEnd={this.onDragEnd}>
        <Droppable droppableId="droppable">
          {(provided, snapshot) => (
            <div
              {...provided.droppableProps}
              ref={provided.innerRef}
              style={getListStyle(snapshot.isDraggingOver)}
            >
              <table>
                {this.state.items.map((item, index) => (
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
              </table>
              {provided.placeholder}
            </div>
          )}
        </Droppable>
      </DragDropContext>
    );
  }
}

export default Dnd;
