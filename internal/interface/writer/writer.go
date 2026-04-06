package writer

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type CLIWriter struct {
	writer *tabwriter.Writer
}

func New(cfg Config) *CLIWriter {
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
