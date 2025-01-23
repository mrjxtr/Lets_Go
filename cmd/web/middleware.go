package main

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mrjxtr/Lets_Go/config"
)

// Neutered implements http.FileSystem
// to disable directory listing
type Neutered struct {
	fs http.FileSystem
}

func (nfs Neutered) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

// initLogger returns a new structured logger
func InitLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// AddSource: true, // ? This is to add the file and line number
		Level: slog.LevelDebug,
	}))

	return logger
}

func InitApp() *config.Application {
	return &config.Application{
		Logger: InitLogger(),
	}
}
