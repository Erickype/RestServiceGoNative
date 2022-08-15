package server

import (
	"fmt"
	"net/http"
)

func Index(w *http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}

		fmt.Fprintf(w, "Hello there %s", "visitor")
	})
}
