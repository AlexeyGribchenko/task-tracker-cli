package config

import (
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/infrastructure/repository/sqlite"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/colors"
	"github.com/AlexeyGribchenko/task-tracker-cli/internal/interface/writer"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Sqlite sqlite.Config
	Format writer.Config
	Color  colors.Config
}

func ParseConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(configPath, cfg)

	return cfg, err
}
