package builtins

import (
	"fmt"
	"io"
	"strings"

	"github.com/azuyamat/hermit/internal/types"
)

type Export struct{}

func (e *Export) Name() string {
	return "export"
}

func (e *Export) Execute(args []string, context *types.ExecutionContext, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	if len(args) == 0 {
		for _, env := range context.Env {
			_, err := fmt.Fprintf(stdout, "%s=%s\n", env, context.GetEnv(env))
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("export: invalid argument: %s", arg)
		}
		context.SetEnv(parts[0], parts[1])
	}
	return nil
}
