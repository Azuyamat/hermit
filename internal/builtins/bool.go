package builtins

import (
	"github.com/azuyamat/hermit/internal/command"
	"github.com/azuyamat/hermit/internal/types"
)

type True struct{}

func (t *True) Metadata() command.Metadata {
	return command.NewMetadataBuilder("true", "Do nothing, successfully").Build()
}

func (t *True) Execute(ctx *command.Context, shell *types.ExecutionContext) error {
	return nil
}

type False struct{}

func (f *False) Metadata() command.Metadata {
	return command.NewMetadataBuilder("false", "Do nothing, unsuccessfully").Build()
}

func (f *False) Execute(ctx *command.Context, shell *types.ExecutionContext) error {
	return types.NewErrExitCode(1)
}
