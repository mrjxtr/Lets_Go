package main

import (
	"flag"
)

type config struct {
	addr      string
	staticDir string
}

// DevConfig returns the config for the development environment
func DevConfig() *config {
	cfg := &config{}

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(
		&cfg.staticDir,
		"static-dir",
		"./ui/static/",
		"Path to static assets",
	)

	flag.Parse()
	return cfg
}

// ProdConfig returns the config for the production environment
func ProdConfig() *config {
	cfg := &config{}

	flag.StringVar(&cfg.addr, "addr", ":9999", "HTTP network address")
	flag.StringVar(
		&cfg.staticDir,
		"static-dir",
		"./ui/static/",
		"Path to static assets",
	)

	flag.Parse()
	return cfg
}
