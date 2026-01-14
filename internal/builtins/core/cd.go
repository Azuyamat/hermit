package core

import (
	"fmt"
	"os"
	"os/user"

	"github.com/azuyamat/hermit/internal/command"
	"github.com/azuyamat/hermit/internal/types"
)

type Cd struct{}

func (c *Cd) Metadata() command.Metadata {
	return command.Metadata{
		Name:        "cd",
		Description: "Change the current working directory",
		Usage:       "cd [directory]",
		Flags:       []command.FlagDef{},
		MinArgs:     0,
		MaxArgs:     1,
	}
}

func (c *Cd) Execute(ctx *command.Context, shell *types.ExecutionContext) error {
	var target string

	oldDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	if ctx.ArgCount() == 0 || ctx.ArgOr(0, "~") == "~" {
		user, err := user.Current()
		if err != nil {
			return fmt.Errorf("failed to get current user: %w", err)
		}
		target = user.HomeDir
	} else if ctx.ArgOr(0, "") == "-" {
		target, ok := shell.GetEnv("OLDPWD")
		if !ok {
			return fmt.Errorf("OLDPWD not set")
		}
		shell.SetEnv("OLDPWD", oldDir)
		shell.SetEnv("PWD", target)
		return nil
	} else {
		target = ctx.ArgOr(0, "")
	}

	if err := os.Chdir(target); err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}

	newDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get new directory: %w", err)
	}

	shell.SetEnv("OLDPWD", oldDir)
	shell.SetEnv("PWD", newDir)

	return nil
}
