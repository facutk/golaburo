package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/facutk/golaburo/api"
	"github.com/facutk/golaburo/api/todos"
	"github.com/facutk/golaburo/db"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var err error
	db.Pool, err = pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Pool.Close()

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/todos", todos.HandleGetAll).Methods("GET")
	r.HandleFunc("/api/v1/todos", todos.HandleCreate).Methods("POST")
	r.HandleFunc("/api/v1/todos/{todoId}", todos.HandleGet).Methods("GET")
	r.HandleFunc("/api/v1/todos/{todoId}", todos.HandleDelete).Methods("DELETE")
	r.HandleFunc("/api/v1/todos/{todoId}", todos.HandleUpdate).Methods("PUT")

	r.HandleFunc("/api/v1/hits", api.HandleHits)
	r.HandleFunc("/api/v1/ping", api.HandlePing)
	r.HandleFunc("/api/v1/uuid", api.HandleUUID)
	r.HandleFunc("/api/v1/dummy", api.HandleDummy)
	r.PathPrefix("/").Handler(api.CacheControlWrapper(http.FileServer(http.Dir("./build"))))

	http.Handle("/", r)
	http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), nil)
}
