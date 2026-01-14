package command

import (
	"fmt"
	"io"
)

type BaseCommand struct{}

func PrintUsage(cmd Command, w io.Writer) {
	meta := cmd.Metadata()
	fmt.Fprintf(w, "Usage: %s\n", meta.Usage)

	if meta.Description != "" {
		fmt.Fprintf(w, "\n%s\n", meta.Description)
	}

	if len(meta.Flags) > 0 {
		fmt.Fprintln(w, "\nFlags:")
		for _, flag := range meta.Flags {
			printFlag(w, flag)
		}
	}

	printArgInfo(cmd, w)
}

func printFlag(w io.Writer, flag FlagDef) {
	var short, long string
	if flag.Short != "" {
		short = fmt.Sprintf("-%s", flag.Short)
	}
	long = fmt.Sprintf("--%s", flag.Name)

	var flagStr string
	if short != "" {
		flagStr = fmt.Sprintf("  %s, %s", short, long)
	} else {
		flagStr = fmt.Sprintf("      %s", long)
	}

	switch flag.Type {
	case Str:
		flagStr += " <string>"
	case Int:
		flagStr += " <int>"
	}

	var defaultStr string
	switch flag.Type {
	case Bool:
		if flag.Default.(bool) {
			defaultStr = " (default: true)"
		}
	case Str:
		if s := flag.Default.(string); s != "" {
			defaultStr = fmt.Sprintf(" (default: %q)", s)
		}
	case Int:
		if i := flag.Default.(int); i != 0 {
			defaultStr = fmt.Sprintf(" (default: %d)", i)
		}
	}

	fmt.Fprintf(w, "%s%s\n", flagStr, defaultStr)
}

func printArgInfo(cmd Command, w io.Writer) {
	meta := cmd.Metadata()
	if meta.MinArgs == 0 && meta.MaxArgs == 0 {
		return
	}

	fmt.Fprintln(w, "\nArguments:")

	if meta.MinArgs == meta.MaxArgs {
		if meta.MinArgs == 1 {
			fmt.Fprintln(w, "  Requires exactly 1 argument")
		} else {
			fmt.Fprintf(w, "  Requires exactly %d arguments\n", meta.MinArgs)
		}
	} else if meta.MaxArgs == -1 {
		if meta.MinArgs == 0 {
			fmt.Fprintln(w, "  Accepts any number of arguments")
		} else {
			fmt.Fprintf(w, "  Requires at least %d argument(s)\n", meta.MinArgs)
		}
	} else {
		fmt.Fprintf(w, "  Requires %d-%d arguments\n", meta.MinArgs, meta.MaxArgs)
	}
}

func UsageError(cmd Command, ctx *Context, msg string) error {
	fmt.Fprintln(ctx.Stderr(), msg)
	PrintUsage(cmd, ctx.Stderr())
	return fmt.Errorf("%s", msg)
}
