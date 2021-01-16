package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "POSTGRES_URI:", os.Getenv("POSTGRES_URI"))
	})

	http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), nil)
}
