package writer

type Config struct {
	MinWidth int `env:"FORMAT_MIN_WIDTH"`
	TabWidth int `env:"FORMAT_TAB_WIDTH"`
	Padding  int `env:"FORMAT_PADDING"`
}
