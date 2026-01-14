package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/azuyamat/hermit/internal/command"
	"github.com/azuyamat/hermit/internal/types"
)

type Export struct{}

func (e *Export) Metadata() command.Metadata {
	return command.NewMetadataBuilder("export", "Set environment variables").
		Usage("export [name=value ...]").
		Flags().
		Build()
}

func (e *Export) Execute(ctx *command.Context, shell *types.ExecutionContext) error {
	if ctx.ArgCount() == 0 {
		for name, value := range shell.Env {
			fmt.Fprintf(os.Stdout, "%s=%s\n", name, value)
		}
		return nil
	}

	for i := 0; i < ctx.ArgCount(); i++ {
		arg := ctx.ArgOr(i, "")
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("export: invalid argument: %s", arg)
		}
		shell.SetEnv(parts[0], parts[1])
		os.Setenv(parts[0], parts[1])
	}
	return nil
}
