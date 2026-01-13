package core

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

	oldDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	if len(args) == 0 || args[0] == "~" {
		user, err := user.Current()
		if err != nil {
			return fmt.Errorf("failed to get current user: %w", err)
		}
		target = user.HomeDir
	} else if args[0] == "-" {
		target = context.GetEnv("OLDPWD")
		if target == "" {
			return fmt.Errorf("OLDPWD not set")
		}
	} else {
		target = args[0]
	}

	if err := os.Chdir(target); err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}

	newDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get new directory: %w", err)
	}

	context.SetEnv("OLDPWD", oldDir)
	context.SetEnv("PWD", newDir)

	return nil
}
