package core

import (
	"os"

	"github.com/azuyamat/hermit/internal/command"
	"github.com/azuyamat/hermit/internal/types"
)

type Exit struct{}

func (e *Exit) Metadata() command.Metadata {
	return command.NewMetadataBuilder("exit", "Exit the shell with a status code").
		Usage("exit [status]").
		Flags().
		MinArgs(0).
		MaxArgs(1).
		Build()
}

func (e *Exit) Execute(ctx *command.Context, shell *types.ExecutionContext) error {
	exitCode := ctx.ArgIntOr(0, 0)
	if exitCode < 0 {
		exitCode = 255 + (exitCode % 256)
	}
	os.Exit(exitCode)
	return nil
}
