package writer

type Config struct {
	MaxColumnWidth int      `env:"FORMAT_MAX_WIDTH" env-default:"40"`
	ExtraColumns   []string `env:"FORMAT_EXTRA_COLUMNS" env-default:""`
}
