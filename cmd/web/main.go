package main

import (
	"net/http"
	"os"

	"github.com/mrjxtr/Lets_Go/config"
)

func main() {
	cfg := config.DevConfig() // ? use ProdConfig() for production
	app := InitApp()
	// logger := InitLogger()
	// app := &config.Application{
	// 	Logger: logger,
	// }
	mux := http.NewServeMux()

	// Disabled FileServer directory listings using Neutered struct
	fileServer := http.FileServer(Neutered{http.Dir(cfg.StaticDir)})
	mux.Handle("GET /static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	app.Logger.Info("Starting server", "addr", cfg.Addr)

	err := http.ListenAndServe(cfg.Addr, mux)
	app.Logger.Error(err.Error())
	os.Exit(1)
}
