package core

import (
	"io"

	"github.com/azuyamat/hermit/internal/types"
)

type Echo struct{}

func (e *Echo) Name() string {
	return "echo"
}

func (e *Echo) Execute(args []string, context *types.ExecutionContext, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	for _, arg := range args {
		_, err := io.WriteString(stdout, arg+" ")
		if err != nil {
			return err
		}
	}
	_, err := io.WriteString(stdout, "\n")
	if err != nil {
		return err
	}
	return nil
}
