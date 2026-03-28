package main

import (
	"fmt"
	"os"

	"github.com/AlexeyGribchenko/task-tracker-cli/internal/config"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/infrastructure/repository/sqlite"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/cli"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/usecase"
)

const configPath = ".env"

// TODO: configuring from config-file not from .env?
// Example: debug on\off, format: ...
// To avoid excess building

// TODO: add statistics: percentage of tasks by status | by time (today, weekly, montly)
func main() {

	cfg := config.ParseConfig(configPath)

	db, err := sqlite.New(cfg.Sqlite)
	if err != nil {
		panic("failed to initialize db: " + err.Error())
	}

	getUC := usecase.NewGetTasksUseCase(db)
	createUC := usecase.NewCreateTaskUse(db)
	updateUC := usecase.NewUpdateTaskUseCase(db)

	writer := cli.NewWriter(cfg.Format)

	app := cli.New(createUC, getUC, updateUC, writer)

	if err := app.Run(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	os.Exit(0)
}
