package config

import (
	"flag"
	"log/slog"
)

type Application struct {
	Logger *slog.Logger
}

type ServerCfg struct {
	Addr      string
	StaticDir string
}

// DevConfig returns the config for the development environment
func DevConfig() *ServerCfg {
	cfg := &ServerCfg{}

	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(
		&cfg.StaticDir,
		"static-dir",
		"./ui/static/",
		"Path to static assets",
	)

	flag.Parse()
	return cfg
}

// ProdConfig returns the config for the production environment
func ProdConfig() *ServerCfg {
	cfg := &ServerCfg{}

	flag.StringVar(&cfg.Addr, "addr", ":9999", "HTTP network address")
	flag.StringVar(
		&cfg.StaticDir,
		"static-dir",
		"./ui/static/",
		"Path to static assets",
	)

	flag.Parse()
	return cfg
}
