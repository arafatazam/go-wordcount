package main

func (a *app) routes() {
	a.router.HandleFunc("/wordcount", a.handleWordCount())
}
