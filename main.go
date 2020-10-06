package main

import (
	"flag"

	"go.uber.org/fx"

	"github.com/cronui/cronui/internal/config"
	"github.com/cronui/cronui/internal/database"
	"github.com/cronui/cronui/internal/server"
)

func main() {
	cf := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	app := fx.New(
		// fx.NopLogger,
		config.Module(*cf),
		database.Module,
		server.Module,
	)
	app.Run()
}
