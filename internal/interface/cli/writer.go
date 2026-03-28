package cli

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Config struct {
	MinWidth int `env:"FORMAT_MIN_WIDTH"`
	TabWidth int `env:"FORMAT_TAB_WIDTH"`
	Padding  int `env:"FORMAT_PADDING"`
}

type CLIWriter struct {
	writer *tabwriter.Writer
}

func NewWriter(cfg Config) *CLIWriter {
	wr := tabwriter.NewWriter(
		os.Stdout,
		cfg.MinWidth,
		cfg.TabWidth,
		cfg.Padding,
		' ',
		tabwriter.TabIndent,
	)
	return &CLIWriter{writer: wr}
}

func (cli *CLIWriter) Print(content string) {
	fmt.Fprintln(cli.writer, content)
}

func (cli *CLIWriter) Flush() error {
	return cli.writer.Flush()
}
