package main

import (
	"net/http"
	"os"
)

func main() {
	cfg := DevConfig() // ? use ProdConfig() for production
	logger := InitLogger()
	mux := http.NewServeMux()

	// Disabled FileServer directory listings using Neutered struct
	fileServer := http.FileServer(Neutered{http.Dir(cfg.staticDir)})
	mux.Handle("GET /static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	logger.Info("Starting server", "addr", cfg.addr)

	err := http.ListenAndServe(cfg.addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
