package core

import (
	"strings"

	"github.com/azuyamat/hermit/internal/command"
	"github.com/azuyamat/hermit/internal/types"
)

type Echo struct{}

func (e *Echo) Metadata() command.Metadata {
	return command.NewMetadataBuilder("echo", "Display a line of text").
		Usage("echo [options] [string ...]").
		Flags(command.NewBoolFlag("no-newline", "n", "Do not output the trailing newline").Build()).
		Build()
}

func (e *Echo) Execute(ctx *command.Context, shell *types.ExecutionContext) error {
	output := strings.Join(ctx.Args(), " ")

	if ctx.Bool("no-newline") {
		ctx.Print(output)
	} else {
		ctx.Println(output)
	}

	return nil
}
