package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	_ "github.com/joho/godotenv/autoload"
)

var conn *pgx.Conn

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	http.Handle("/", http.FileServer(http.Dir("./build")))

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	http.HandleFunc("/uuid", func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		fmt.Fprint(w, id.String())
	})

	http.HandleFunc("/api/v1/todos", func(w http.ResponseWriter, r *http.Request) {
		type Todo struct {
			ID          uuid.UUID
			Description string
		}

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
			fmt.Fprint(w, string(marshalledTodos))
		case http.MethodPost:
			id := uuid.New()

			_, err := conn.Exec(context.Background(), "INSERT INTO todos (id, description) VALUES ($1, $2);", id, "hello todos!")
			if err != nil {
				fmt.Fprintf(os.Stderr, "insert failed: %v\n", err)
				os.Exit(1)
			}

			fmt.Fprint(w, id.String())
		case http.MethodPut:
			// Update an existing record.
		case http.MethodDelete:
			// Remove the record.
		default:
			// Give an error message.
		}
	})

	http.HandleFunc("/api/v1/hits", func(w http.ResponseWriter, r *http.Request) {
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

	http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), nil)
}
