package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	_ "github.com/joho/godotenv/autoload"
)

// Todo structure
type Todo struct {
	ID          uuid.UUID
	Description string
}

var conn *pgx.Conn

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./build")))

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	r.HandleFunc("/uuid", func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		fmt.Fprint(w, id.String())
	})

	r.HandleFunc("/api/v1/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todos := []Todo{}

			rows, _ := conn.Query(context.Background(), "SELECT todo.id, todo.description FROM todos todo")
			for rows.Next() {
				var todo = Todo{}
				err := rows.Scan(&todo.ID, &todo.Description)
				todos = append(todos, todo)
				if err != nil {
					fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
					os.Exit(1)
				}
			}
			marshalledTodos, _ := json.Marshal(todos)
			w.Header().Set("Content-Type", "application/json")
			w.Write(marshalledTodos)
		case http.MethodPost:
			t := Todo{}

			// Try to decode the request body into the struct. If there is an error,
			// respond to the client with the error message and a 400 status code.
			errDecode := json.NewDecoder(r.Body).Decode(&t)
			if errDecode != nil {
				http.Error(w, errDecode.Error(), http.StatusBadRequest)
				return
			}
			t.ID = uuid.New()
			_, err := conn.Exec(context.Background(), "INSERT INTO todos (id, description) VALUES ($1, $2);", t.ID, t.Description)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			marshalledTodo, _ := json.Marshal(t)
			w.Header().Set("Content-Type", "application/json")
			w.Write(marshalledTodo)
		case http.MethodPut:
			// Update an existing record.
		case http.MethodDelete:
			// Remove the record.
		default:
			// Give an error message.
		}
	})

	r.HandleFunc("/api/v1/todos/{todoId}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		todoID := params["todoId"]

		switch r.Method {
		case http.MethodGet:
			todo := Todo{}
			err = conn.QueryRow(context.Background(), "select todo.id, todo.description FROM todos todo where id=$1", todoID).Scan(&todo.ID, &todo.Description)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			marshalledTodo, _ := json.Marshal(todo)
			w.Header().Set("Content-Type", "application/json")
			w.Write(marshalledTodo)
		case http.MethodPut:
			// Update an existing record.
		case http.MethodDelete:
			_, err = conn.Exec(context.Background(), "delete FROM todos where id=$1", todoID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		default:
			// Give an error message.
		}
	})

	r.HandleFunc("/api/v1/hits", func(w http.ResponseWriter, r *http.Request) {
		var hits int
		err = conn.QueryRow(context.Background(), "select hits from visits where id=1").Scan(&hits)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}

		_, err := conn.Exec(context.Background(), "update visits set hits=$1 where id=1", hits+1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "update Exec failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprint(w, hits)
	})
	http.Handle("/", r)
	http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), nil)
}
