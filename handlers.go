package main

import (
	"fmt"
	"net/http"
)

func (a *app) handleWordCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Counting...")
	}
}
