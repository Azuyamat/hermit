package builtins

import (
	"io"

	"github.com/azuyamat/hermit/internal/types"
)

type Clear struct{}

func (c *Clear) Name() string {
	return "clear"
}

func (c *Clear) Execute(args []string, context *types.ExecutionContext, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	_, err := stdout.Write([]byte("\033[2J\033[H\n"))
	return err
}
