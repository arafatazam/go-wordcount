package main

import (
	"log"
	"net/http"
)

type app struct {
	config *Config
	log    *log.Logger
	router *http.ServeMux
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.log.Println("New request")
	a.router.ServeHTTP(w, r)
}

func newApp(cfg *Config) *app {
	a := &app{
		config: cfg,
		log:    log.Default(),
		router: http.NewServeMux(),
	}
	a.routes()
	return a
}
