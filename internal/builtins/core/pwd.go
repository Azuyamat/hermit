package core

import (
	"fmt"

	"github.com/azuyamat/hermit/internal/command"
	"github.com/azuyamat/hermit/internal/types"
)

type Pwd struct{}

func (p *Pwd) Metadata() command.Metadata {
	return command.NewMetadataBuilder("pwd", "Print the current working directory").
		Usage("pwd").
		ExactArgs(0).
		Build()
}

func (p *Pwd) Execute(ctx *command.Context, shell *types.ExecutionContext) error {
	dir, err := shell.GetWorkingDir()
	if err != nil {
		return fmt.Errorf("pwd: failed to get working directory: %w", err)
	}
	ctx.Println(dir)
	return nil
}
