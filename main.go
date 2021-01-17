package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/facutk/golaburo/dummy"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

// Todo type
type Todo struct {
	ID          uuid.UUID
	Description string
}

var conn *pgx.Conn

func main() {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todos := []Todo{}

			rows, _ := pool.Query(context.Background(), "SELECT todo.id, todo.description FROM todos todo")
			for rows.Next() {
				var todo = Todo{}
				err := rows.Scan(&todo.ID, &todo.Description)
				todos = append(todos, todo)
				if err != nil {
					fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
					os.Exit(1)
				}
			}
			marshalledTodos, _ := json.MarshalIndent(todos, "", "  ")
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
			_, err := pool.Exec(context.Background(), "INSERT INTO todos (id, description) VALUES ($1, $2);", t.ID, t.Description)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			marshalledTodo, _ := json.MarshalIndent(t, "", "  ")
			w.Header().Set("Content-Type", "application/json")
			w.Write(marshalledTodo)
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
			err = pool.QueryRow(context.Background(), "select todo.id, todo.description FROM todos todo where id=$1", todoID).Scan(&todo.ID, &todo.Description)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			marshalledTodo, _ := json.Marshal(todo)
			w.Header().Set("Content-Type", "application/json")
			w.Write(marshalledTodo)
		case http.MethodDelete:
			_, err = pool.Exec(context.Background(), "delete FROM todos where id=$1", todoID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		case http.MethodPut:
			// Update an existing record.
		default:
			// Give an error message.
		}
	})

	r.HandleFunc("/api/v1/hits", func(w http.ResponseWriter, r *http.Request) {
		var hits int
		err = pool.QueryRow(context.Background(), "select hits from visits where id=1").Scan(&hits)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}

		_, err := pool.Exec(context.Background(), "update visits set hits=$1 where id=1", hits+1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "update Exec failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprint(w, hits)
	})

	r.HandleFunc("/api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	r.HandleFunc("/api/v1/uuid", func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		fmt.Fprint(w, id.String())
	})

	r.HandleFunc("/api/v1/dummy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, dummy.Foo())
	})

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./build")))

	http.Handle("/", r)
	http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), nil)
}
