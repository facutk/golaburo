package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/facutk/golaburo/db"
)

// HandleHits middleware
func HandleHits(w http.ResponseWriter, r *http.Request) {
	var err error
	var hits int
	err = db.Pool.QueryRow(context.Background(), "select hits from visits where id=1").Scan(&hits)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	_, err = db.Pool.Exec(context.Background(), "update visits set hits=$1 where id=1", hits+1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "update Exec failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprint(w, hits)
}
