package writer

type Config struct {
	MaxColumnWidth int      `env:"FORMAT_MAX_WIDTH"`
	ExtraColumns   []string `env:"FORMAT_EXTRA_COLUMNS"`
}
