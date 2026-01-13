package core

import (
	"io"
	"os"
	"strconv"

	"github.com/azuyamat/hermit/internal/types"
)

type Exit struct{}

func (e *Exit) Name() string {
	return "exit"
}

func (e *Exit) Execute(args []string, context *types.ExecutionContext, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	exitCode := 0

	if len(args) > 0 {
		code, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		exitCode = code
	}

	os.Exit(exitCode)
	return nil
}
