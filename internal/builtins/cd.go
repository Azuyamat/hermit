package builtins

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/azuyamat/hermit/internal/types"
)

type Cd struct{}

func (c *Cd) Name() string {
	return "cd"
}

func (c *Cd) Execute(args []string, context *types.ExecutionContext, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	var target string

	if len(args) == 0 {
		user, err := user.Current()
		if err != nil {
			return fmt.Errorf("failed to get current user: %w", err)
		}
		target = user.HomeDir
	} else {
		target = args[0]
	}

	if err := os.Chdir(target); err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}

	return nil
}
