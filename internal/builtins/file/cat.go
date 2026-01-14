package file

import (
	"io"
	"os"

	"github.com/azuyamat/hermit/internal/command"
	"github.com/azuyamat/hermit/internal/types"
)

type Cat struct{}

func (c *Cat) Metadata() command.Metadata {
	return command.NewMetadataBuilder("cat", "Concatenate and display file contents").
		Usage("cat [file ...]").
		MinArgs(0).
		MaxArgs(-1).
		Build()
}

func (c *Cat) Execute(ctx *command.Context, shell *types.ExecutionContext) error {
	if len(ctx.Args()) == 0 {
		_, err := io.Copy(ctx.Stdout(), ctx.Stdin())
		ctx.Println()
		return err
	}
	for _, filename := range ctx.Args() {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(ctx.Stdout(), file)
		if err != nil {
			return err
		}
		ctx.Println()
	}
	return nil
}
