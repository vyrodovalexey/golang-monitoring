package main

import (
	"github.com/vyrodovalexey/golang-monitoring/internal/handlers"
	"github.com/vyrodovalexey/golang-monitoring/storage"
	"net/http"
)

func run() error {
	return http.ListenAndServe(`:8080`, nil)
}

func main() {
	gauge := make(map[string]storage.GaugeItem)
	counter := make(map[string][]storage.CounterItem)

	storage.Storage = storage.MemStorage{gauge, counter}

	http.HandleFunc("/update/", handlers.Update)
	if err := run(); err != nil {
		panic(err)
	}
}