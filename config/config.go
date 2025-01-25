package config

import (
	"flag"
)

type ServerCfg struct {
	Addr      string
	StaticDir string
	Sqlite3   string
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
	flag.StringVar(
		&cfg.Sqlite3,
		"sqlite3",
		"./data/snippetbox.db",
		"Sqlite database",
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
