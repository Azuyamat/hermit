package command

import "github.com/azuyamat/hermit/internal/types"

type Command interface {
	Metadata() Metadata
	Execute(ctx *Context, shell *types.ExecutionContext) error
}
