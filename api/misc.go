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

// CacheControlWrapper adds cache
func CacheControlWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var cache = "max-age=31536000" // 1 year

		if r.URL.String() == "/" {
			cache = "max-age=0"
		}

		w.Header().Set("Cache-Control", cache)

		h.ServeHTTP(w, r)
	})
}
