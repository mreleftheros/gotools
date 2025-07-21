package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/mreleftheros/gotools/srv/json"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		l := slog.New(slog.NewJSONHandler(os.Stderr, nil))

		err := json.Parse(w, r, &input)
		if err != nil {
			json.Write(l, w, r, 400, json.NewErrorResponse(err.Error(), nil), nil)
			return
		}
		json.Write(l, w, r, 200, json.NewDataResponse(input), nil)
	})

	log.Fatal(http.ListenAndServe(":3000", mux))
}
