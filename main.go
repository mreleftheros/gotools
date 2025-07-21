package main

import (
	"log"
	"net/http"

	"github.com/mreleftheros/gotools/srv/json"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m := map[string]string{
			"nameError":     "name is invalid",
			"passwordError": "password is invalid too",
		}
		json.Write(w, 400, json.NewErrorResponse("some error", m), nil)
	})

	log.Fatal(http.ListenAndServe(":3000", mux))
}
