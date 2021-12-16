package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/caarlos0/env"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return err
	}
	a := newApp(cfg)
	fmt.Fprintf(os.Stdout, "Listening on port %d\n...\n", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), a)
}
