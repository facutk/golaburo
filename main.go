package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

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
