package main

import (
	"log"
	"net/http"
)

func main() {
	cfg := DevConfig() // ? Uncomment this line to use development config
	// cfg := ProdConfig() // ? Uncomment this line to use production config

	mux := http.NewServeMux()

	// Disabled FileServer directory listings using Neutered struct
	fileServer := http.FileServer(Neutered{http.Dir(cfg.staticDir)})
	mux.Handle("GET /static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("Starting server on %s", cfg.addr)

	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)
}
