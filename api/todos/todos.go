package todos

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/facutk/golaburo/db"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type todo struct {
	ID          uuid.UUID
	Description string
	Created     time.Time
}

// HandleGetAll middleware
func HandleGetAll(w http.ResponseWriter, r *http.Request) {
	todos := []todo{}

	rows, err := db.Pool.Query(context.Background(),
		"SELECT todo.id, todo.description, todo.created FROM todos todo ORDER BY todo.created DESC")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var todo = todo{}
		err := rows.Scan(&todo.ID, &todo.Description, &todo.Created)
		todos = append(todos, todo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}
	}
	defer rows.Close()
	marshalledTodos, _ := json.MarshalIndent(todos, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalledTodos)
}

// HandleCreate middleware
func HandleCreate(w http.ResponseWriter, r *http.Request) {
	t := todo{}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	errDecode := json.NewDecoder(r.Body).Decode(&t)
	if errDecode != nil {
		http.Error(w, errDecode.Error(), http.StatusBadRequest)
		return
	}
	t.ID = uuid.New()
	_, err := db.Pool.Exec(context.Background(), "INSERT INTO todos (id, description) VALUES ($1, $2);", t.ID, t.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	marshalledTodo, _ := json.MarshalIndent(t, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalledTodo)
}

// HandleGet middleware
func HandleGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID := params["todoId"]

	todo := todo{}
	err := db.Pool.QueryRow(context.Background(), "select todo.id, todo.description FROM todos todo where id=$1", todoID).Scan(&todo.ID, &todo.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	marshalledTodo, _ := json.MarshalIndent(todo, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalledTodo)
}

// HandleDelete middleware
func HandleDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID := params["todoId"]

	_, err := db.Pool.Exec(context.Background(), "delete FROM todos where id=$1", todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// HandleUpdate middleware
func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// todoID := params["todoId"]
	// Update an existing record.
}
