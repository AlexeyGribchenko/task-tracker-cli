package colors

import "github.com/fatih/color"

type Config struct {
	ColoredOutput bool `env:"COLORED_OUTPUT" env-default:"false"`
}

func Init(cfg Config) {
	color.NoColor = !cfg.ColoredOutput
}
