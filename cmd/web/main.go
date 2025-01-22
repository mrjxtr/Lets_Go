package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := DevConfig() // ? Uncomment this line to use development config
	// cfg := ProdConfig() // ? Uncomment this line to use production config

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// AddSource: true, // ? This is to add the file and line number
		Level: slog.LevelDebug,
	}))

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
