package main

import (
	"github.com/Erickype/RestServiceGoNative/server"
)

func main() {

	// server init config
	server := server.NewServer(":8080")

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
