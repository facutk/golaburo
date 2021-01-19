package api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// HandlePing middleware
func HandlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

// HandleUUID middleware
func HandleUUID(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	fmt.Fprint(w, id.String())
}

// HandleDummy middleware
func HandleDummy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, Foo())
}
