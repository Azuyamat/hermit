package core

import (
	"github.com/azuyamat/hermit/internal/command"
	"github.com/azuyamat/hermit/internal/types"
)

type Clear struct{}

func (c *Clear) Metadata() command.Metadata {
	return command.NewMetadataBuilder("clear", "Clear the terminal screen").
		Usage("clear").
		Flags().
		Build()
}

func (c *Clear) Execute(ctx *command.Context, shell *types.ExecutionContext) error {
	ctx.Print("\033[H\033[2J")
	return nil
}
