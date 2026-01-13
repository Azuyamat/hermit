package builtins

import (
	"fmt"
	"io"

	"github.com/azuyamat/hermit/internal/types"
)

type Pwd struct{}

func (p *Pwd) Name() string {
	return "pwd"
}

func (p *Pwd) Execute(args []string, context *types.ExecutionContext, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	dir, err := context.GetWorkingDir()
	if err != nil {
		return fmt.Errorf("pwd: failed to get working directory: %w", err)
	}
	_, err = io.WriteString(stdout, dir+"\n")
	if err != nil {
		return fmt.Errorf("pwd: failed to write output: %w", err)
	}
	return nil
}
