package server

import "net/http"

type Country struct {
	Name     string
	Languaje string
}

var countries []*Country = []*Country{}

func NewServer(addr string) *http.Server {
	InitRoutes()
	return &http.Server{
		Addr: addr,
	}
}
