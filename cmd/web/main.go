package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/mrjxtr/Lets_Go/config"
)

type application struct {
	logger *slog.Logger
}

var cfg = config.DevConfig() // ? use ProdConfig() for production

func main() {
	// initLogger returns a new structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// AddSource: true, // ? This is to add the file and line number
		Level: slog.LevelDebug,
	}))

	app := &application{
		logger: logger,
	}

	app.logger.Info("Starting server", "addr", cfg.Addr)

	err := http.ListenAndServe(cfg.Addr, app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)
}
