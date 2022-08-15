package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
