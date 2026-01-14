package command

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Parse(meta Metadata, args []string, stdout, stderr io.Writer, stdin io.Reader) (*Context, error) {
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			return nil, ErrShowHelp
		}
	}

	ctx := &Context{
		flags:  make(map[string]any),
		stdout: stdout,
		stderr: stderr,
		stdin:  stdin,
	}

	for _, flag := range meta.Flags {
		ctx.flags[flag.Name] = flag.Default
	}

	var positional []string
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if !strings.HasPrefix(arg, "-") {
			positional = append(positional, arg)
			continue
		}

		flag, value, err := parseFlag(arg, args, &i, meta.Flags)
		if err != nil {
			return nil, err
		}
		ctx.flags[flag.Name] = value
	}

	if len(positional) < meta.MinArgs {
		return nil, fmt.Errorf("expected at least %d args, got %d", meta.MinArgs, len(positional))
	}
	if meta.MaxArgs >= 0 && len(positional) > meta.MaxArgs {
		return nil, fmt.Errorf("expected at most %d args, got %d", meta.MaxArgs, len(positional))
	}

	ctx.args = positional
	return ctx, nil
}

func parseFlag(arg string, allArgs []string, idx *int, defs []FlagDef) (*FlagDef, any, error) {
	var flagName string
	var hasValue bool
	var rawValue string

	if strings.HasPrefix(arg, "--") {
		parts := strings.SplitN(arg[2:], "=", 2)
		flagName = parts[0]
		if len(parts) == 2 {
			hasValue = true
			rawValue = parts[1]
		}
	} else {
		flagName = strings.TrimPrefix(arg, "-")
	}

	// Find flag definition
	var def *FlagDef
	for i := range defs {
		if defs[i].Name == flagName || defs[i].Short == flagName {
			def = &defs[i]
			break
		}
	}
	if def == nil {
		return nil, nil, fmt.Errorf("unknown flag: %s", arg)
	}

	switch def.Type {
	case Bool:
		if hasValue {
			return def, rawValue == "true", nil
		}
		return def, true, nil

	case Str:
		if !hasValue {
			if *idx+1 >= len(allArgs) {
				return nil, nil, fmt.Errorf("flag %s requires a value", arg)
			}
			*idx++
			rawValue = allArgs[*idx]
		}
		return def, rawValue, nil

	case Int:
		if !hasValue {
			if *idx+1 >= len(allArgs) {
				return nil, nil, fmt.Errorf("flag %s requires a value", arg)
			}
			*idx++
			rawValue = allArgs[*idx]
		}
		val, err := strconv.Atoi(rawValue)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid integer for %s: %s", arg, rawValue)
		}
		return def, val, nil

	default:
		return nil, nil, fmt.Errorf("unsupported flag type")
	}
}

var ErrShowHelp = fmt.Errorf("show help")
