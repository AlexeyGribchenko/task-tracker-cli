package main

import (
	"fmt"
	"os"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/config"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/infrastructure/repository/sqlite"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/colors"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/writer"
	"github.com/fatih/color"
)

const configPath = ".env"

// TODO: configuring from config-file not from .env?
// Example: debug on\off, format: ...
// To avoid excess building

// TODO: add statistics: percentage of tasks by status | by time (today, weekly, montly)
func main() {

	if err := run(); err != nil {
		fmt.Println(color.RedString("Error:"), err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {

	cfg := config.ParseConfig(configPath)

	colors.Init(cfg.Color)

	db, err := sqlite.New(cfg.Sqlite)
	if err != nil {
		panic("failed to initialize db: " + err.Error())
	}

	writer := writer.New(cfg.Format)
	app := cli.New(db, *writer)

	return app.Run()
}
