package main

import (
	"database/sql"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/form/v4"
	_ "github.com/mattn/go-sqlite3"

	"github.com/mrjxtr/Lets_Go/config"
	"github.com/mrjxtr/Lets_Go/internal/models"
)

type application struct {
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
}

// Configuration
var cfg = config.DevConfig() // ? use ProdConfig() for production

func main() {
	// Dependencies
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// AddSource: true, // ? This is to add the file and line number
		Level: slog.LevelDebug,
	}))

	db, err := openDB(cfg.Sqlite3)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	app := &application{
		logger:        logger,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
		formDecoder:   formDecoder,
	}

	// Running server
	app.logger.Info("Starting server", "addr", cfg.Addr)

	err = http.ListenAndServe(cfg.Addr, app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)
}

func openDB(sqlite3 string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqlite3)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
