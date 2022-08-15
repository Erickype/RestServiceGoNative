package server

import (
	"fmt"
	"net/http"
)

type country struct {
	Name     string
	Languaje string
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}

		fmt.Fprintf(w, "Hello there %s", "visitor")
	})
}

func GetCountries(w *http.ResponseWriter, r *http.Request) {

}

func AddCountrie(w *http.ResponseWriter, r *http.Response) {

}
