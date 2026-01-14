package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/azuyamat/hermit/internal/command"
	"github.com/azuyamat/hermit/internal/types"
)

type Wc struct{}

func (w *Wc) Metadata() command.Metadata {
	return command.NewMetadataBuilder("wc", "Count lines, words, and bytes in files").
		Usage("wc [options] [file ...]").
		Flags(
			command.NewBoolFlag("lines", "l", "Print the newline counts").Build(),
			command.NewBoolFlag("words", "w", "Print the word counts").Build(),
			command.NewBoolFlag("bytes", "c", "Print the byte counts").Build(),
			command.NewBoolFlag("chars", "m", "Print the character counts").Build(),
		).
		MinArgs(0).
		MaxArgs(-1).
		Build()
}

func (w *Wc) Execute(ctx *command.Context, shell *types.ExecutionContext) error {
	showLines := ctx.Bool("lines")
	showWords := ctx.Bool("words")
	showBytes := ctx.Bool("bytes")
	showChars := ctx.Bool("chars")

	if !showLines && !showWords && !showBytes && !showChars {
		showLines = true
		showWords = true
		showBytes = true
	}

	if len(ctx.Args()) == 0 {
		counts := w.count(ctx.Stdin())
		w.printCounts(counts, "", showLines, showWords, showBytes, showChars, ctx.Stdout())
		return nil
	}

	var total *counts
	for _, filename := range ctx.Args() {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		fileCounts := w.count(file)
		w.printCounts(fileCounts, filename, showLines, showWords, showBytes, showChars, ctx.Stdout())

		if total == nil {
			total = &counts{}
		}
		total.lines += fileCounts.lines
		total.words += fileCounts.words
		total.bytes += fileCounts.bytes
		total.chars += fileCounts.chars
	}

	if len(ctx.Args()) > 1 {
		w.printCounts(total, "total", showLines, showWords, showBytes, showChars, ctx.Stdout())
	}

	return nil
}

type counts struct {
	lines int
	words int
	bytes int
	chars int
}

func (w *Wc) count(r io.Reader) *counts {
	c := &counts{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		c.lines++
		c.bytes += len(scanner.Bytes()) + 1
		c.chars += len([]rune(line)) + 1

		fields := strings.Fields(line)
		c.words += len(fields)
	}

	return c
}

func (w *Wc) printCounts(c *counts, name string, showLines, showWords, showBytes, showChars bool, stdout io.Writer) {
	var parts []string

	if showLines {
		parts = append(parts, fmt.Sprintf("%7d", c.lines))
	}
	if showWords {
		parts = append(parts, fmt.Sprintf("%7d", c.words))
	}
	if showBytes {
		parts = append(parts, fmt.Sprintf("%7d", c.bytes))
	}
	if showChars {
		parts = append(parts, fmt.Sprintf("%7d", c.chars))
	}

	output := strings.Join(parts, " ")
	if name != "" {
		output += " " + name
	}

	fmt.Fprintln(stdout, output)
}
