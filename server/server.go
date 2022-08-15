package server

import "net/http"

func NewServer(addr string) *http.Server {
	InitRoutes()
	return &http.Server{
		Addr: addr,
	}
}
