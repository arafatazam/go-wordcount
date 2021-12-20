package main

import (
	"bufio"
	"net/http"
)

func (a *app) handleWordCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scanner := bufio.NewScanner(r.Body)
		counter := &Counter{mp: make(map[string]int)}
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			counter.Count(scanner.Text())
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(counter.Top())
	}
}
