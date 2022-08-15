package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}

		fmt.Fprintf(w, "Hello there %s", "visitor")
	})

	// server init config
	server := http.Server{
		Addr: ":8080",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
