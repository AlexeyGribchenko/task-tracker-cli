package colors

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

type Colorer interface {
	Bold(msg string) string
	Green(msg string) string
	Red(msg string) string
	Yellow(msg string) string
	Blue(msg string) string
}

func New(cfg Config) Colorer {
	if cfg.ColoredOutput {
		return &RealColorer{}
	}
	return &NoopColorer{}
}
